import pytest
import time
from acktest.k8s import resource as k8s
from acktest.resources import load_resource_file

RESOURCE_PLURAL = "documents"
RESOURCE_NAME = "test-document"

@pytest.fixture(scope="module")
def document():
    resource_data = load_resource_file("ssm", "document")
    reference, spec = k8s.create_custom_resource(
        resource_plural=RESOURCE_PLURAL,
        custom_resource_name=RESOURCE_NAME,
        spec=resource_data,
    )
    yield reference, spec
    k8s.delete_custom_resource(reference)
    time.sleep(5)

def test_document_create(document):
    reference, spec = document
    assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)
    created_doc = k8s.get_resource(reference)
    assert created_doc is not None
    assert created_doc["spec"]["name"] == RESOURCE_NAME

def test_document_update(document):
    reference, spec = document
    update_data = load_resource_file("ssm", "document_update")
    k8s.patch_custom_resource(reference, update_data)
    assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)
    updated_doc = k8s.get_resource(reference)
    assert updated_doc["spec"]["content"] == update_data["spec"]["content"]

def test_document_delete():
    k8s.delete_custom_resource_by_name(RESOURCE_PLURAL, RESOURCE_NAME)
    time.sleep(5)
    result = k8s.get_resource_by_name(RESOURCE_PLURAL, RESOURCE_NAME)
    assert result is None