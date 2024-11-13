import pytest
import logging
import time
from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import service_marker, load_ssm_resource, CRD_GROUP, CRD_VERSION
from e2e.replacement_values import REPLACEMENT_VALUES

RESOURCE_PLURAL = "documents"
CREATE_WAIT_AFTER_SECONDS = 20
DELETE_WAIT_AFTER_SECONDS = 20
MODIFY_WAIT_AFTER_SECONDS = 20


@pytest.fixture(scope="module")
def document():
    RESOURCE_NAME = random_suffix_name("document", 24)

    resources = get_bootstrap_resources()
    logging.debug(resources)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["NAME"] = RESOURCE_NAME
    resource_data = load_ssm_resource("document", additional_replacements=replacements,)

    reference = k8s.CustomResourceReference(
        CRD_GROUP,
        CRD_VERSION,
        RESOURCE_PLURAL,
        RESOURCE_NAME,
        namespace='default',
    )

    k8s.create_custom_resource(reference, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(reference)

    logging.debug(cr)

    assert cr is not None
    assert k8s.get_resource_exists(reference)
    yield reference, cr

    k8s.delete_custom_resource(reference)
    time.sleep(DELETE_WAIT_AFTER_SECONDS)


@service_marker
class TestDocument:
    def test_create_delete(self, document):
        (reference, _) = document
        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        cr = k8s.get_resource(reference)
        assert cr is not None
        assert 'spec' in cr
        assert 'name' in cr["spec"]
        assert 'content' in cr["spec"]

        # Update test
        update_data = {
                "spec": {
                    "content": 
                        {
                        "schemaVersion": "2.2",
                        "description": "Sample SSM Document",
                        "mainSteps": [
                            {
                            "action": "aws:runShellScript",
                            "name": "example",
                            "inputs": {
                                "runCommand": ["echo Update Hello, World!"]
                            }
                            }
                        ]
                    }
            }
        }
        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)
        
        updated_cr = k8s.get_resource(reference)       
        assert updated_cr["spec"]["content"] == update_data["spec"]["content"]
