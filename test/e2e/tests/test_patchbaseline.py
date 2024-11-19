import pytest
import logging
import time
from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import service_marker, load_ssm_resource, CRD_GROUP, CRD_VERSION
from e2e.replacement_values import REPLACEMENT_VALUES

RESOURCE_PLURAL = "patchbaselines"
CREATE_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 10
MODIFY_WAIT_AFTER_SECONDS = 10

@pytest.fixture(scope="module")
def patchbaseline():
    RESOURCE_NAME = random_suffix_name("patchbaseline", 24)

    resources = get_bootstrap_resources()
    logging.debug(resources)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["NAME"] = RESOURCE_NAME
    resource_data = load_ssm_resource("patchbaseline", additional_replacements=replacements,)

    logging.debug(resource_data)

    reference = k8s.CustomResourceReference(
        CRD_GROUP,
        CRD_VERSION,
        RESOURCE_PLURAL,
        RESOURCE_NAME,
        namespace='default',
    )

    k8s.create_custom_resource(reference, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(reference)

    assert cr is not None
    assert k8s.get_resource_exists(reference)
    yield reference, cr

    k8s.delete_custom_resource(reference)
    time.sleep(DELETE_WAIT_AFTER_SECONDS)

@service_marker
class TestPatchBaseline:
    def test_create_delete(self, patchbaseline, ssm_client):
        (reference, _) = patchbaseline
        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        cr = k8s.get_resource(reference)
        assert cr is not None
        assert 'spec' in cr
        assert 'name' in cr["spec"]
        assert 'operatingSystem' in cr["spec"]
        assert 'approvedPatches' in cr["spec"]
        assert 'approvalRules' in cr["spec"]

        update_data = {
            "spec": {
                "approvedPatches": ["KB5029928"]
            }
        }


        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        updated_cr = k8s.get_resource(reference)       
        assert updated_cr["spec"]["approvedPatches"] == update_data["spec"]["approvedPatches"]
        assert "approvalRules" in updated_cr["spec"]
        
        assert "patchRules" in updated_cr["spec"]["approvalRules"]
        assert len(updated_cr["spec"]["approvalRules"]["patchRules"]) == len(cr["spec"]["approvalRules"]["patchRules"])

        baselineID = updated_cr["status"]["baselineID"]

        # Check AWS API
        patch_aws = ssm_client.get_patch_baseline(BaselineId=baselineID)

        assert patch_aws is not None
        assert patch_aws["BaselineId"] == baselineID
        assert patch_aws["Name"] == updated_cr["spec"]["name"]
        assert patch_aws["OperatingSystem"] == updated_cr["spec"]["operatingSystem"]
        assert set(patch_aws["ApprovedPatches"]) == set(updated_cr["spec"]["approvedPatches"])

        # Check approval rules
        assert "ApprovalRules" in patch_aws
        assert "PatchRules" in patch_aws["ApprovalRules"]

        for aws_rule, cr_rule in zip(patch_aws["ApprovalRules"]["PatchRules"], updated_cr["spec"]["approvalRules"]["patchRules"]):
            assert aws_rule["ApproveAfterDays"] == cr_rule["approveAfterDays"]
            assert aws_rule["ComplianceLevel"] == cr_rule["complianceLevel"]
            
            aws_filters = aws_rule["PatchFilterGroup"]["PatchFilters"]
            cr_filters = cr_rule["patchFilterGroup"]["patchFilters"]
            
            for aws_filter, cr_filter in zip(aws_filters, cr_filters):
                assert aws_filter["Key"] == cr_filter["key"]
                assert set(aws_filter["Values"]) == set(cr_filter["values"])