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
    replacements["BUCKET_NAME"] = resources.ResourceSyncBucket.name
    resource_data = load_ssm_resource("resourcedatasync", additional_replacements=replacements)

    reference = k8s.CustomResourceReference(
        CRD_GROUP,
        CRD_VERSION,
        RESOURCE_PLURAL,
        RESOURCE_NAME,
        namespace='default',
    )

    try:
        k8s.create_custom_resource(reference, resource_data)
    except Exception as e:
        logging.error(f"Failed to create ResourceDataSync: {e}")
        raise

    max_attempts = 10
    for attempt in range(max_attempts):
        time.sleep(5)  # Wait for 5 seconds between checks
        if k8s.get_resource_exists(reference):
            logging.info(f"ResourceDataSync created successfully after {attempt + 1} attempts")
            break
        logging.warning(f"ResourceDataSync not found, attempt {attempt + 1}/{max_attempts}")
    else:
        raise TimeoutError("ResourceDataSync was not created within the expected time")

    assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=20)
    
    cr = k8s.get_resource(reference)
    assert cr is not None
    logging.info(f"ResourceDataSync CR: {cr}")

    yield reference, cr

    # Cleanup
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
