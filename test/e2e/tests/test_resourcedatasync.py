import pytest
import logging
import time
from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import service_marker, load_ssm_resource, CRD_GROUP, CRD_VERSION
from e2e.replacement_values import REPLACEMENT_VALUES

RESOURCE_PLURAL = "resourcedatasyncs"
CREATE_WAIT_AFTER_SECONDS = 20
DELETE_WAIT_AFTER_SECONDS = 20
MODIFY_WAIT_AFTER_SECONDS = 20

@pytest.fixture(scope="module")
def resourcedatasync():
    RESOURCE_NAME = random_suffix_name("resourcedatasync", 12)

    resources = get_bootstrap_resources()
    logging.debug(f"Bootstrap resources: {resources}")

    replacements = REPLACEMENT_VALUES.copy()
    resource_data = load_ssm_resource("resourcedatasync", additional_replacements=replacements)

    reference = k8s.CustomResourceReference(
        CRD_GROUP,
        CRD_VERSION,
        RESOURCE_PLURAL,
        RESOURCE_NAME,
        namespace='default',
    )

    k8s.create_custom_resource(reference, resource_data)
    assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=20)
    
    cr = k8s.get_resource(reference)
    assert cr is not None
    logging.info(f"ResourceDataSync CR: {cr}")

    yield reference, cr

    k8s.delete_custom_resource(reference)
    time.sleep(DELETE_WAIT_AFTER_SECONDS)


@service_marker
class TestResourceDataSync:
    def test_create_delete(self, resourcedatasync):
        (reference, _) = resourcedatasync

        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        cr = k8s.get_resource(reference)
        assert cr is not None
        assert 'spec' in cr
        assert cr["spec"]["syncType"] == "SyncFromSource"
        assert 'syncSource' in cr["spec"]
        assert 'sourceType' in cr["spec"]["syncSource"]
        assert cr["spec"]["syncSource"]["sourceType"] == "SingleAccountMultiRegions"
        assert 'sourceRegions' in cr["spec"]["syncSource"]

        update_data = {
            "spec": {
                "syncSource": {
                    "sourceRegions": ["us-west-2", "us-east-1"]
                }
            }
        }

        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        updated_cr = k8s.get_resource(reference)
        assert updated_cr is not None
        assert 'sourceRegions' in updated_cr["spec"]["syncSource"]
        assert set(updated_cr["spec"]["syncSource"]["sourceRegions"]) == set(["us-west-2", "us-east-1"])