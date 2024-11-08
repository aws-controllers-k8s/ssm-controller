import pytest
from acktest.k8s import resource as k8s
from acktest.resources import load_resource_file
import time

RESOURCE_PLURAL = "resourcedatasyncs"
RESOURCE_NAME = "test-resourcedatasync"

@pytest.fixture(scope="module")
def resourcedatasync():
    resource_data = load_resource_file("ssm", "resourcedatasync")
    reference, _ = k8s.create_custom_resource(
        resource_plural=RESOURCE_PLURAL,
        custom_resource_name=RESOURCE_NAME,
        spec=resource_data,
    )
    yield reference
    k8s.delete_custom_resource(reference)

def test_resourcedatasync_create(resourcedatasync):
    reference = resourcedatasync
    assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

    resource = k8s.get_resource(reference)
    assert resource is not None
    assert resource["spec"]["syncName"] == RESOURCE_NAME

def test_resourcedatasync_update(resourcedatasync):
    reference = resourcedatasync
    update_data = load_resource_file("ssm", "resourcedatasync_update")
    k8s.patch_custom_resource(reference, update_data)
    assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

    updated_resource = k8s.get_resource(reference)
    assert updated_resource["spec"]["s3Destination"]["prefix"] == update_data["spec"]["s3Destination"]["prefix"]

def test_resourcedatasync_delete():
    k8s.delete_custom_resource_by_name("resourcedatasyncs", RESOURCE_NAME)
    time.sleep(5)
    assert k8s.get_resource_by_name("resourcedatasyncs", RESOURCE_NAME) is None