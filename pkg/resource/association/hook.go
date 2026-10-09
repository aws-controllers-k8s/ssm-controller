// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package association

import (
	"context"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/ssm"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

func (rm *resourceManager) syncTags(
	ctx context.Context,
	desired *resource,
	latest *resource,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.syncTags")
	defer func() {
		exit(err)
	}()

	desiredTags, _ := convertToOrderedACKTags(desired.ko.Spec.Tags)
	latestTags, _ := convertToOrderedACKTags(latest.ko.Spec.Tags)
	added, _, removed := ackcompare.GetTagsDifference(latestTags, desiredTags)

	if len(removed) > 0 {
		tagKeys := make([]string, 0, len(removed))
		for key := range removed {
			tagKeys = append(tagKeys, key)
		}
		rlog.Debug("removing tags from resource", "tags", tagKeys)
		_, err = rm.sdkapi.RemoveTagsFromResource(ctx, &svcsdk.RemoveTagsFromResourceInput{
			ResourceId:   aws.String(*latest.ko.Status.AssociationID),
			ResourceType: svcsdktypes.ResourceTypeForTaggingAssociation,
			TagKeys:      tagKeys,
		})
		rm.metrics.RecordAPICall("UPDATE", "RemoveTagsFromResource", err)
		if err != nil {
			return err
		}
	}

	if len(added) > 0 {
		rlog.Debug("adding tags to resource", "tags", added)
		_, err = rm.sdkapi.AddTagsToResource(ctx, &svcsdk.AddTagsToResourceInput{
			ResourceId:   aws.String(*latest.ko.Status.AssociationID),
			ResourceType: svcsdktypes.ResourceTypeForTaggingAssociation,
			Tags:         rm.sdkTags(added),
		})
		rm.metrics.RecordAPICall("UPDATE", "AddTagsToResource", err)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rm *resourceManager) sdkTags(tags map[string]string) []svcsdktypes.Tag {
	sdkTags := make([]svcsdktypes.Tag, 0, len(tags))
	for key, value := range tags {
		sdkTags = append(sdkTags, svcsdktypes.Tag{
			Key:   aws.String(key),
			Value: aws.String(value),
		})
	}
	return sdkTags
}

func (rm *resourceManager) fetchCurrentTags(
	ctx context.Context,
	resourceID *string,
) (map[string]string, error) {
	output, err := rm.sdkapi.ListTagsForResource(ctx, &svcsdk.ListTagsForResourceInput{
		ResourceId:   resourceID,
		ResourceType: svcsdktypes.ResourceTypeForTaggingAssociation,
	})
	if err != nil {
		return nil, err
	}

	tags := make(map[string]string, len(output.TagList))
	for _, tag := range output.TagList {
		tags[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
	}
	return tags, nil
}
