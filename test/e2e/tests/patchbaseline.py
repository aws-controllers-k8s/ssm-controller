import pytest
import logging
import time
from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import service_marker, load_ssm_resource
from e2e.replacement_values import REPLACEMENT_VALUES

RESOURCE_PLURAL = "patchbaselines"
CREATE_WAIT_AFTER_SECONDS = 20
DELETE_WAIT_AFTER_SECONDS = 20
MODIFY_WAIT_AFTER_SECONDS = 20

@pytest.fixture(scope="module")
def patchbaseline():
    RESOURCE_NAME = random_suffix_name("patchbaseline", 12)

    resources = get_bootstrap_resources()
    logging.debug(resources)

    replacements = REPLACEMENT_VALUES.copy()
    resource_data = load_ssm_resource("patchbaseline", additional_replacements=replacements)
    reference, spec = k8s.create_custom_resource(
        resource_plural=RESOURCE_PLURAL,
        custom_resource_name=RESOURCE_NAME,
        spec=resource_data,
    )
    assert reference is not None

    time.sleep(CREATE_WAIT_AFTER_SECONDS)
    yield reference, spec

    k8s.delete_custom_resource(reference)
    time.sleep(DELETE_WAIT_AFTER_SECONDS)

@service_marker
class TestPatchBaseline:
    def test_create_delete(self, patchbaseline):
        (reference, spec) = patchbaseline

        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        cr = k8s.get_resource(reference)
        assert cr is not None
        assert 'spec' in cr
        assert 'name' in cr["spec"]
        assert 'operatingSystem' in cr["spec"]
        assert 'approvedPatches' in cr["spec"]
        assert 'approvalRules' in cr["spec"]

        # Update test
        update_data = load_ssm_resource("patchbaseline_update", additional_replacements={})
        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)
        
        updated_cr = k8s.get_resource(reference)       
        assert updated_cr["spec"]["approvedPatches"] == update_data["spec"]["approvedPatches"]
