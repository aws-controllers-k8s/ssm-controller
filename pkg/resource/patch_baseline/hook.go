package patch_baseline

import (
	"context"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/ssm"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

// syncTags used to keep tags in sync by calling Create and Delete API's
func (rm *resourceManager) syncTags(
	ctx context.Context,
	desired *resource,
	latest *resource,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.syncTags")

	defer func(err error) {
		exit(err)
	}(err)

	resourceID := latest.ko.Status.BaselineID

	desiredTags := ToACKTags(desired.ko.Spec.Tags)
	latestTags := ToACKTags(latest.ko.Spec.Tags)

	added, _, removed := ackcompare.GetTagsDifference(latestTags, desiredTags)

	toAdd := FromACKTags(added)

	var toDeleteTagKeys []string
	for k := range removed {
		toDeleteTagKeys = append(toDeleteTagKeys, k)
	}

	// Remove tags
	if len(toDeleteTagKeys) > 0 {
		rlog.Debug("removing tags from resource", "tags", toDeleteTagKeys)
		_, err = rm.sdkapi.RemoveTagsFromResource(
			ctx,
			&svcsdk.RemoveTagsFromResourceInput{
				ResourceType: svcsdktypes.ResourceTypeForTaggingPatchBaseline,
				ResourceId:   aws.String(*resourceID),
				TagKeys:      toDeleteTagKeys,
			},
		)

		rm.metrics.RecordAPICall("UPDATE", "RemoveTagsFromResource", err)
		if err != nil {
			return err
		}
	}

	// Add tags
	if len(toAdd) > 0 {
		rlog.Debug("adding tags to resource", "tags", toAdd)
		_, err = rm.sdkapi.AddTagsToResource(
			ctx,
			&svcsdk.AddTagsToResourceInput{
				ResourceType: svcsdktypes.ResourceTypeForTaggingPatchBaseline,
				ResourceId:   aws.String(*resourceID),
				Tags:         rm.sdkTags(added),
			},
		)

		rm.metrics.RecordAPICall("UPDATE", "AddTagsToResource", err)
		if err != nil {
			return err
		}
	}
	return nil
}

// sdkTags converts *svcapitypes.Tag array to a *svcsdk.Tag array
func (rm *resourceManager) sdkTags(tags map[string]string) (sdkTags []svcsdktypes.Tag) {

	for key, value := range tags {
		sdktag := svcsdktypes.Tag{
			Key:   aws.String(key),
			Value: aws.String(value),
		}
		sdkTags = append(sdkTags, sdktag)
	}
	return sdkTags
}

func compareTags(
	delta *ackcompare.Delta,
	a *resource,
	b *resource,
) {
	if len(a.ko.Spec.Tags) != len(b.ko.Spec.Tags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	} else if len(a.ko.Spec.Tags) > 0 {
		desiredTags := ToACKTags(a.ko.Spec.Tags)
		latestTags := ToACKTags(b.ko.Spec.Tags)

		added, _, removed := ackcompare.GetTagsDifference(latestTags, desiredTags)

		if len(added) != 0 || len(removed) != 0 {
			delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
		}
	}
}

func (rm *resourceManager) fetchCurrentTags(
	ctx context.Context,
	resourceID *string,
) (map[string]string, error) {
	output, err := rm.sdkapi.ListTagsForResource(
		ctx,
		&svcsdk.ListTagsForResourceInput{
			ResourceId:   resourceID,
			ResourceType: svcsdktypes.ResourceTypeForTaggingPatchBaseline,
		},
	)

	if err != nil {
		return nil, err
	}

	tags := make(map[string]string)

	for _, tag := range output.TagList {
		tags[*tag.Key] = *tag.Value
	}
	return tags, nil
}
