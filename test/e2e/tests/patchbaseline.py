import pytest
from acktest.k8s import resource as k8s
from acktest.resources import load_resource_file
import time

RESOURCE_PLURAL = "patchbaselines"
RESOURCE_NAME = "test-patchbaseline"


@pytest.fixture(scope="module")
def patchbaseline():
    resource_data = load_resource_file("ssm", "patchbaseline")
    reference, _ = k8s.create_custom_resource(
        resource_plural=RESOURCE_PLURAL,
        custom_resource_name=RESOURCE_NAME,
        spec=resource_data,
    )
    yield reference
    k8s.delete_custom_resource(reference)


def test_patchbaseline_create(patchbaseline):
    reference = patchbaseline
    assert k8s.wait_on_condition(
        reference, "ACK.ResourceSynced", "True", wait_periods=10)

    resource = k8s.get_resource(reference)
    assert resource is not None
    assert resource["spec"]["name"] == RESOURCE_NAME


def test_patchbaseline_update(patchbaseline):
    reference = patchbaseline
    update_data = load_resource_file("ssm", "patchbaseline_update")
    k8s.patch_custom_resource(reference, update_data)
    assert k8s.wait_on_condition(
        reference, "ACK.ResourceSynced", "True", wait_periods=10)

    updated_resource = k8s.get_resource(reference)
    assert updated_resource["spec"]["approvedPatches"] == update_data["spec"]["approvedPatches"]


def test_patchbaseline_delete():
    k8s.delete_custom_resource_by_name("patchbaselines", RESOURCE_NAME)
    time.sleep(5)
    assert k8s.get_resource_by_name("patchbaselines", RESOURCE_NAME) is None