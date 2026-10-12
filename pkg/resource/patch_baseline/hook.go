package patch_baseline

import (
	"context"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/ssm"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"

	"github.com/aws-controllers-k8s/ssm-controller/pkg/tags"
)

// syncTags keeps tags in sync by delegating to the shared tag-sync helper.
func (rm *resourceManager) syncTags(
	ctx context.Context,
	desired *resource,
	latest *resource,
) (err error) {
	return tags.SyncTags(
		ctx,
		rm.sdkapi,
		rm.metrics,
		svcsdktypes.ResourceTypeForTaggingPatchBaseline,
		*latest.ko.Status.BaselineID,
		desired.ko.Spec.Tags,
		latest.ko.Spec.Tags,
	)
}

func compareTags(
	delta *ackcompare.Delta,
	a *resource,
	b *resource,
) {
	if len(a.ko.Spec.Tags) != len(b.ko.Spec.Tags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	} else if len(a.ko.Spec.Tags) > 0 {
		desiredTags, _ := convertToOrderedACKTags(a.ko.Spec.Tags)
		latestTags, _ := convertToOrderedACKTags(b.ko.Spec.Tags)

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
