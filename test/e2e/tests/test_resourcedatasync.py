import pytest
import logging
import time
from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import service_marker, load_ssm_resource
from e2e.replacement_values import REPLACEMENT_VALUES

RESOURCE_PLURAL = "resourcedatasyncs"
CREATE_WAIT_AFTER_SECONDS = 20
DELETE_WAIT_AFTER_SECONDS = 20
MODIFY_WAIT_AFTER_SECONDS = 20

@pytest.fixture(scope="module")
def resourcedatasync():
    RESOURCE_NAME = random_suffix_name("ssm", 12)

    resources = get_bootstrap_resources()
    logging.debug(resources)

    replacements = REPLACEMENT_VALUES.copy()
    resource_data = load_ssm_resource("resourcedatasync", additional_replacements=replacements)
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
class TestResourceDataSync:
    def test_create_delete(self, resourcedatasync):
        (reference, spec) = resourcedatasync

        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        cr = k8s.get_resource(reference)
        assert cr is not None
        assert 'spec' in cr
        assert 'syncName' in cr["spec"]
        assert 'syncType' in cr["spec"]
        assert 's3Destination' in cr["spec"]
        assert 'bucketName' in cr["spec"]["s3Destination"]
        assert 'syncFormat' in cr["spec"]["s3Destination"]
        assert 'region' in cr["spec"]["s3Destination"]
        assert 'prefix' in cr["spec"]["s3Destination"]

        # Update test
        update_data = {
            "spec": {
                "s3Destination": {
                    "prefix": "updated-sync-prefix"
                }
            }
        }

        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        updated_cr = k8s.get_resource(reference)       
        assert updated_cr["spec"]["s3Destination"]["prefix"] == update_data["spec"]["s3Destination"]["prefix"]
