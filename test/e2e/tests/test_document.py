import pytest
import logging
import json
import time
from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import service_marker, load_ssm_resource, CRD_GROUP, CRD_VERSION
from e2e.replacement_values import REPLACEMENT_VALUES
from acktest.tags import to_dict, clean, assert_ack_system_tags, assert_equal_without_ack_tags

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
                "content": json.dumps({
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
                })
            }
        }

        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)
        
        updated_cr = k8s.get_resource(reference)       
        assert updated_cr["spec"]["content"] == update_data["spec"]["content"]

    def test_update_document_with_tags(self, document, ssm_client):
        """Regression test: UpdateDocumentInput has no Tags member, so a
        tags-only delta used to drive a no-op UpdateDocument every reconcile
        and the tag change was never applied. The sdk_update_pre_build_request
        hook now calls syncTags (Add/RemoveTagsFromResource), mirroring the
        PatchBaseline pattern. Verify the full add -> modify -> remove cycle
        against the AWS API (the CR spec round-trips through sdkFind and would
        mask a no-op tag sync, so we assert on list_tags_for_resource).
        """
        (reference, _) = document
        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        cr = k8s.get_resource(reference)
        assert cr is not None
        document_name = cr["spec"]["name"]

        # Add tags via spec.tags
        add_tags = {
            "spec": {
                "tags": [
                    {"key": "environment", "value": "dev"},
                    {"key": "team", "value": "ack"},
                ]
            }
        }
        k8s.patch_custom_resource(reference, add_tags)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        tags_aws = ssm_client.list_tags_for_resource(
            ResourceType="Document",
            ResourceId=document_name,
        )["TagList"]
        expected_tags = [
            {"Key": tag["key"], "Value": tag["value"]} for tag in add_tags["spec"]["tags"]
        ]
        clean_tags_aws = clean(tags=tags_aws, key_member_name='Key')
        clean_expected_tags = clean(tags=expected_tags, key_member_name='Key')
        assert_equal_without_ack_tags(
            expected=clean_expected_tags,
            actual=clean_tags_aws,
            key_member_name='Key',
            value_member_name='Value'
        )

        # Modify to a DIFFERENT tag set: changes one value, drops "team",
        # adds "owner". This exercises BOTH AddTagsToResource and
        # RemoveTagsFromResource.
        modify_tags = {
            "spec": {
                "tags": [
                    {"key": "environment", "value": "prod"},
                    {"key": "owner", "value": "platform"},
                ]
            }
        }
        k8s.patch_custom_resource(reference, modify_tags)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        tags_aws = ssm_client.list_tags_for_resource(
            ResourceType="Document",
            ResourceId=document_name,
        )["TagList"]
        expected_tags = [
            {"Key": tag["key"], "Value": tag["value"]} for tag in modify_tags["spec"]["tags"]
        ]
        clean_tags_aws = clean(tags=tags_aws, key_member_name='Key')
        clean_expected_tags = clean(tags=expected_tags, key_member_name='Key')
        # Equality against {environment:prod, owner:platform} inherently proves
        # the "team" tag was removed.
        assert_equal_without_ack_tags(
            expected=clean_expected_tags,
            actual=clean_tags_aws,
            key_member_name='Key',
            value_member_name='Value'
        )

        # Clear all tags via spec.tags: []. Mirrors PatchBaseline's remove step:
        # only ACK system tags should remain afterwards.
        remove_tags = {
            "spec": {
                "tags": []
            }
        }
        k8s.patch_custom_resource(reference, remove_tags)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        tags_aws = ssm_client.list_tags_for_resource(
            ResourceType="Document",
            ResourceId=document_name,
        )["TagList"]
        assert_ack_system_tags(tags=tags_aws, key_member_name='Key', value_member_name='Value')
