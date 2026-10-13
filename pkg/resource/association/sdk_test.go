package association

import (
	"context"
	"testing"

	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	"github.com/aws/aws-sdk-go-v2/aws"

	svcapitypes "github.com/aws-controllers-k8s/ssm-controller/apis/v1alpha1"
)

func TestSDKFindRequiresAssociationID(t *testing.T) {
	rm := &resourceManager{}
	r := &resource{ko: &svcapitypes.Association{}}

	_, err := rm.sdkFind(context.Background(), r)
	if err != ackerr.NotFound {
		t.Fatalf("expected NotFound for an association without an ID, got %v", err)
	}
}

func TestAssociationTagDelta(t *testing.T) {
	tag := func(key, value string) *svcapitypes.Tag {
		return &svcapitypes.Tag{Key: aws.String(key), Value: aws.String(value)}
	}
	desired := &resource{ko: &svcapitypes.Association{
		Spec: svcapitypes.AssociationSpec{
			Tags: []*svcapitypes.Tag{tag("environment", "production")},
		},
	}}
	latest := &resource{ko: &svcapitypes.Association{
		Spec: svcapitypes.AssociationSpec{
			Tags: []*svcapitypes.Tag{tag("environment", "development")},
		},
	}}

	delta := newResourceDelta(desired, latest)
	if !delta.DifferentAt("Spec.Tags") {
		t.Fatal("expected a tag value change to be detected")
	}
}
