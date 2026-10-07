package association

import (
	"context"
	"testing"

	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"

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
