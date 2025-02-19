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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PatchBaselineSpec defines the desired state of PatchBaseline.
type PatchBaselineSpec struct {

	// A set of rules used to include patches in the baseline.

	ApprovalRules *PatchRuleGroup `json:"approvalRules,omitempty"`
	// A list of explicitly approved patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and
	// rejected patches, see Package name formats for approved and rejected patch
	// lists (https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html)
	// in the Amazon Web Services Systems Manager User Guide.

	ApprovedPatches []*string `json:"approvedPatches,omitempty"`
	// Defines the compliance level for approved patches. When an approved patch
	// is reported as missing, this value describes the severity of the compliance
	// violation. The default value is UNSPECIFIED.

	ApprovedPatchesComplianceLevel *string `json:"approvedPatchesComplianceLevel,omitempty"`
	// Indicates whether the list of approved patches includes non-security updates
	// that should be applied to the managed nodes. The default value is false.
	// Applies to Linux managed nodes only.

	ApprovedPatchesEnableNonSecurity *bool `json:"approvedPatchesEnableNonSecurity,omitempty"`
	// User-provided idempotency token.

	ClientToken *string `json:"clientToken,omitempty"`
	// A description of the patch baseline.

	Description *string `json:"description,omitempty"`
	// A set of global filters used to include patches in the baseline.
	//
	// The GlobalFilters parameter can be configured only by using the CLI or an
	// Amazon Web Services SDK. It can't be configured from the Patch Manager console,
	// and its value isn't displayed in the console.

	GlobalFilters *PatchFilterGroup `json:"globalFilters,omitempty"`
	// The name of the patch baseline.

	// +kubebuilder:validation:Required

	Name *string `json:"name"`
	// Defines the operating system the patch baseline applies to. The default value
	// is WINDOWS.

	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// A list of explicitly rejected patches for the baseline.
	//
	// For information about accepted formats for lists of approved patches and
	// rejected patches, see Package name formats for approved and rejected patch
	// lists (https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html)
	// in the Amazon Web Services Systems Manager User Guide.

	RejectedPatches []*string `json:"rejectedPatches,omitempty"`
	// The action for Patch Manager to take on patches included in the RejectedPackages
	// list.
	//
	// ALLOW_AS_DEPENDENCY
	//
	// Linux and macOS: A package in the rejected patches list is installed only
	// if it is a dependency of another package. It is considered compliant with
	// the patch baseline, and its status is reported as INSTALLED_OTHER. This is
	// the default action if no option is specified.
	//
	// Windows Server: Windows Server doesn't support the concept of package dependencies.
	// If a package in the rejected patches list and already installed on the node,
	// its status is reported as INSTALLED_OTHER. Any package not already installed
	// on the node is skipped. This is the default action if no option is specified.
	//
	// BLOCK
	//
	// All OSs: Packages in the rejected patches list, and packages that include
	// them as dependencies, aren't installed by Patch Manager under any circumstances.
	// If a package was installed before it was added to the rejected patches list,
	// or is installed outside of Patch Manager afterward, it's considered noncompliant
	// with the patch baseline and its status is reported as INSTALLED_REJECTED.

	RejectedPatchesAction *string `json:"rejectedPatchesAction,omitempty"`
	// Information about the patches to use to update the managed nodes, including
	// target operating systems and source repositories. Applies to Linux managed
	// nodes only.

	Sources []*PatchSource `json:"sources,omitempty"`
	// Optional metadata that you assign to a resource. Tags enable you to categorize
	// a resource in different ways, such as by purpose, owner, or environment.
	// For example, you might want to tag a patch baseline to identify the severity
	// level of patches it specifies and the operating system family it applies
	// to. In this case, you could specify the following key-value pairs:
	//
	//    * Key=PatchSeverity,Value=Critical
	//
	//    * Key=OS,Value=Windows
	//
	// To add tags to an existing patch baseline, use the AddTagsToResource operation.

	Tags []*Tag `json:"tags,omitempty"`
}

// PatchBaselineStatus defines the observed state of PatchBaseline
type PatchBaselineStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The ID of the created patch baseline.
	// +kubebuilder:validation:Optional
	BaselineID *string `json:"baselineID,omitempty"`
}

// PatchBaseline is the Schema for the PatchBaselines API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type PatchBaseline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PatchBaselineSpec   `json:"spec,omitempty"`
	Status            PatchBaselineStatus `json:"status,omitempty"`
}

// PatchBaselineList contains a list of PatchBaseline
// +kubebuilder:object:root=true
type PatchBaselineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PatchBaseline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PatchBaseline{}, &PatchBaselineList{})
}
