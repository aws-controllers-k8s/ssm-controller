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

package document

import (
	"context"

	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"

	"github.com/aws-controllers-k8s/ssm-controller/pkg/tags"
)

// syncTags keeps tags in sync by delegating to the shared tag-sync helper.
func (rm *resourceManager) syncTags(
	ctx context.Context,
	desired *resource,
	latest *resource,
) (err error) {
	// Name is a required, immutable field, so it is always present on the
	// desired resource and serves as the tag ResourceId for Documents.
	return tags.SyncTags(
		ctx,
		rm.sdkapi,
		rm.metrics,
		svcsdktypes.ResourceTypeForTaggingDocument,
		*desired.ko.Spec.Name,
		desired.ko.Spec.Tags,
		latest.ko.Spec.Tags,
	)
}
