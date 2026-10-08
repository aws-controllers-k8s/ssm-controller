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

// Package tags centralizes the tag-synchronization logic shared by the
// resource packages whose SSM resources are tagged via the
// AddTagsToResource / RemoveTagsFromResource API pair (currently Document and
// PatchBaseline).
//
// Each pkg/resource/<name> package defines its own concrete resource and
// resourceManager types, so a shared helper cannot take those types or be a
// method on the resource manager. SyncTags therefore accepts only shared and
// primitive types: the SDK client, the metrics object, the ResourceType enum,
// the ResourceId string, and the desired/latest Spec.Tags slices. The ACK-tag
// conversion and diff are performed here using the importable apis/v1alpha1
// Tag type and the runtime compare package, so callers remain thin wrappers
// and this package depends on no per-package generated function.
package tags

import (
	"context"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackmetrics "github.com/aws-controllers-k8s/runtime/pkg/metrics"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/ssm"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"

	svcapitypes "github.com/aws-controllers-k8s/ssm-controller/apis/v1alpha1"
)

// SyncTags keeps the tags on an SSM resource in sync by calling
// RemoveTagsFromResource and AddTagsToResource for the tag keys that were
// removed from / added to the desired Spec.Tags relative to the latest
// observed Spec.Tags.
//
// resourceType selects the SSM tagging resource type (e.g.
// ResourceTypeForTaggingDocument). resourceID is the resource's tagging
// identifier (the Document name, the PatchBaseline ID, ...). desiredTags and
// latestTags are the Spec.Tags slices from the desired and latest resource
// state respectively.
func SyncTags(
	ctx context.Context,
	sdkapi *svcsdk.Client,
	metrics *ackmetrics.Metrics,
	resourceType svcsdktypes.ResourceTypeForTagging,
	resourceID string,
	desiredTags []*svcapitypes.Tag,
	latestTags []*svcapitypes.Tag,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.syncTags")

	defer func(err error) {
		exit(err)
	}(err)

	desired, _ := convertToOrderedACKTags(desiredTags)
	latest, _ := convertToOrderedACKTags(latestTags)

	added, _, removed := ackcompare.GetTagsDifference(latest, desired)

	toAdd := fromACKTags(added, nil)

	var toDeleteTagKeys []string
	for k := range removed {
		toDeleteTagKeys = append(toDeleteTagKeys, k)
	}

	// Remove tags
	if len(toDeleteTagKeys) > 0 {
		rlog.Debug("removing tags from resource", "tags", toDeleteTagKeys)
		_, err = sdkapi.RemoveTagsFromResource(
			ctx,
			&svcsdk.RemoveTagsFromResourceInput{
				ResourceType: resourceType,
				ResourceId:   aws.String(resourceID),
				TagKeys:      toDeleteTagKeys,
			},
		)

		metrics.RecordAPICall("UPDATE", "RemoveTagsFromResource", err)
		if err != nil {
			return err
		}
	}

	// Add tags
	if len(toAdd) > 0 {
		rlog.Debug("adding tags to resource", "tags", toAdd)
		_, err = sdkapi.AddTagsToResource(
			ctx,
			&svcsdk.AddTagsToResourceInput{
				ResourceType: resourceType,
				ResourceId:   aws.String(resourceID),
				Tags:         sdkTags(added),
			},
		)

		metrics.RecordAPICall("UPDATE", "AddTagsToResource", err)
		if err != nil {
			return err
		}
	}
	return nil
}

// sdkTags converts a map of tags to a []svcsdktypes.Tag array
func sdkTags(tags map[string]string) (sdkTags []svcsdktypes.Tag) {
	for key, value := range tags {
		sdktag := svcsdktypes.Tag{
			Key:   aws.String(key),
			Value: aws.String(value),
		}
		sdkTags = append(sdkTags, sdktag)
	}
	return sdkTags
}

// convertToOrderedACKTags converts the tags parameter into 'acktags.Tags'
// shape, returning a slice of keys maintaining the original key order. It
// mirrors the per-package generated helper of the same name so this shared
// package depends on no per-package generated function.
func convertToOrderedACKTags(tags []*svcapitypes.Tag) (acktags.Tags, []string) {
	result := acktags.NewTags()
	keyOrder := []string{}

	if len(tags) == 0 {
		return result, keyOrder
	}
	for _, t := range tags {
		if t.Key != nil {
			keyOrder = append(keyOrder, *t.Key)
			if t.Value != nil {
				result[*t.Key] = *t.Value
			} else {
				result[*t.Key] = ""
			}
		}
	}

	return result, keyOrder
}

// fromACKTags converts the tags parameter into []*svcapitypes.Tag shape,
// maintaining the order given by keyOrder. It mirrors the per-package
// generated helper of the same name.
func fromACKTags(tags acktags.Tags, keyOrder []string) []*svcapitypes.Tag {
	result := []*svcapitypes.Tag{}

	for _, k := range keyOrder {
		v, ok := tags[k]
		if ok {
			tag := svcapitypes.Tag{Key: &k, Value: &v}
			result = append(result, &tag)
			delete(tags, k)
		}
	}
	for k, v := range tags {
		tag := svcapitypes.Tag{Key: &k, Value: &v}
		result = append(result, &tag)
	}

	return result
}
