import pytest
import logging
import time
from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from acktest.aws.identity import get_region
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import service_marker, load_ssm_resource, CRD_GROUP, CRD_VERSION
from e2e.replacement_values import REPLACEMENT_VALUES

RESOURCE_PLURAL = "parameters"
CREATE_WAIT_AFTER_SECONDS = 20
DELETE_WAIT_AFTER_SECONDS = 20
MODIFY_WAIT_AFTER_SECONDS = 20


@pytest.fixture(scope="module")
def simple_parameter():
    RESOURCE_NAME = random_suffix_name("parameter", 24)
    PARAMETER_NAME = f"/ack-test/{RESOURCE_NAME}"

    resources = get_bootstrap_resources()
    logging.debug(resources)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["NAME"] = RESOURCE_NAME
    replacements["PARAMETER_NAME"] = PARAMETER_NAME
    replacements["PARAMETER_VALUE"] = "initial-value"

    resource_data = load_ssm_resource("parameter", additional_replacements=replacements)

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
    yield reference, cr, PARAMETER_NAME

    k8s.delete_custom_resource(reference)
    time.sleep(DELETE_WAIT_AFTER_SECONDS)


@pytest.fixture(scope="module")
def secure_parameter():
    RESOURCE_NAME = random_suffix_name("secure-param", 24)
    PARAMETER_NAME = f"/ack-test/secure/{RESOURCE_NAME}"

    resources = get_bootstrap_resources()
    logging.debug(resources)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["NAME"] = RESOURCE_NAME
    replacements["PARAMETER_NAME"] = PARAMETER_NAME
    replacements["PARAMETER_VALUE"] = "secret-value"

    # Create SecureString parameter
    resource_data = load_ssm_resource("parameter", additional_replacements=replacements)
    resource_data["spec"]["type"] = "SecureString"

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
    yield reference, cr, PARAMETER_NAME

    k8s.delete_custom_resource(reference)
    time.sleep(DELETE_WAIT_AFTER_SECONDS)


@service_marker
class TestParameter:
    def test_create_delete_simple(self, simple_parameter):
        (reference, _, param_name) = simple_parameter
        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        cr = k8s.get_resource(reference)
        assert cr is not None
        assert 'spec' in cr
        assert 'name' in cr["spec"]
        assert cr["spec"]["name"] == param_name
        assert 'type' in cr["spec"]
        assert cr["spec"]["type"] == "String"
        assert 'value' in cr["spec"]
        assert cr["spec"]["value"] == "initial-value"

        # Verify status fields are populated
        assert 'status' in cr
        assert 'ackResourceMetadata' in cr["status"]

    def test_update_parameter_value(self, simple_parameter):
        (reference, _, param_name) = simple_parameter
        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        # Update parameter value
        update_data = {
            "spec": {
                "value": "updated-value"
            }
        }

        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        updated_cr = k8s.get_resource(reference)
        assert updated_cr["spec"]["value"] == "updated-value"

    def test_update_parameter_description(self, simple_parameter):
        (reference, _, param_name) = simple_parameter
        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        # Update parameter description
        update_data = {
            "spec": {
                "description": "Updated description for test parameter"
            }
        }

        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        updated_cr = k8s.get_resource(reference)
        assert updated_cr["spec"]["description"] == "Updated description for test parameter"

    def test_create_delete_secure(self, secure_parameter):
        (reference, _, param_name) = secure_parameter
        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        cr = k8s.get_resource(reference)
        assert cr is not None
        assert 'spec' in cr
        assert 'name' in cr["spec"]
        assert cr["spec"]["name"] == param_name
        assert 'type' in cr["spec"]
        assert cr["spec"]["type"] == "SecureString"
        assert 'value' in cr["spec"]

        # Verify status fields are populated
        assert 'status' in cr
        assert 'ackResourceMetadata' in cr["status"]

    def test_update_secure_parameter_value(self, secure_parameter):
        (reference, _, param_name) = secure_parameter
        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        update_data = {
            "spec": {
                "value": "new-secret-value-1"
            }
        }

        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        updated_cr = k8s.get_resource(reference)
        assert updated_cr["spec"]["value"] == "new-secret-value-1"

        update_data = {
            "spec": {
                "value": "new-secret-value-2"
            }
        }

        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        updated_cr = k8s.get_resource(reference)
        assert updated_cr["spec"]["value"] == "new-secret-value-2"

