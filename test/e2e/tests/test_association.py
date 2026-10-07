import logging
import time

import pytest

from acktest.k8s import resource as k8s
from acktest.resources import random_suffix_name
from e2e import CRD_GROUP, CRD_VERSION, load_ssm_resource, service_marker
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e.replacement_values import REPLACEMENT_VALUES

RESOURCE_PLURAL = "associations"
CREATE_WAIT_AFTER_SECONDS = 20
DELETE_WAIT_AFTER_SECONDS = 20
MODIFY_WAIT_AFTER_SECONDS = 20


@pytest.fixture(scope="module")
def association():
    resource_name = random_suffix_name("association", 24)

    resources = get_bootstrap_resources()
    logging.debug(resources)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["NAME"] = resource_name
    resource_data = load_ssm_resource(
        "association",
        additional_replacements=replacements,
    )

    reference = k8s.CustomResourceReference(
        CRD_GROUP,
        CRD_VERSION,
        RESOURCE_PLURAL,
        resource_name,
        namespace="default",
    )

    k8s.create_custom_resource(reference, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(reference)

    assert cr is not None
    assert k8s.get_resource_exists(reference)
    yield reference, cr, resource_name

    _, deleted = k8s.delete_custom_resource(reference)
    time.sleep(DELETE_WAIT_AFTER_SECONDS)
    assert deleted is True


@service_marker
class TestAssociation:
    def test_create_delete(self, association, ssm_client):
        reference, _, resource_name = association
        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(
            reference,
            "ACK.ResourceSynced",
            "True",
            wait_periods=10,
        )

        cr = k8s.get_resource(reference)
        assert cr is not None
        assert cr["spec"]["name"] == "AWS-RunShellScript"
        assert cr["spec"]["associationName"] == resource_name
        assert cr["status"]["associationID"]

        association_aws = ssm_client.describe_association(
            AssociationId=cr["status"]["associationID"],
        )["AssociationDescription"]
        assert association_aws["Name"] == "AWS-RunShellScript"
        assert association_aws["AssociationName"] == resource_name

    def test_update_parameters(self, association):
        reference, _, _ = association
        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(
            reference,
            "ACK.ResourceSynced",
            "True",
            wait_periods=10,
        )

        update_data = {
            "spec": {
                "parameters": {
                    "commands": ["echo updated ACK association test"],
                },
            },
        }

        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(
            reference,
            "ACK.ResourceSynced",
            "True",
            wait_periods=10,
        )

        updated_cr = k8s.get_resource(reference)
        assert updated_cr["spec"]["parameters"] == update_data["spec"]["parameters"]
