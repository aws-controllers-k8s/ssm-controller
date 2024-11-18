import pytest
import logging
import time
from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import service_marker, load_ssm_resource, CRD_GROUP, CRD_VERSION
from e2e.replacement_values import REPLACEMENT_VALUES

RESOURCE_PLURAL = "resourcedatasyncs"
CREATE_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 10
MODIFY_WAIT_AFTER_SECONDS = 10

@pytest.fixture(scope="module")
def resourcedatasync():
    RESOURCE_NAME = random_suffix_name("resourcedatasync", 24)

    resources = get_bootstrap_resources()
    logging.debug(f"Bootstrap resources: {resources}")

    replacements = REPLACEMENT_VALUES.copy()
    replacements["NAME"] = RESOURCE_NAME 
    resource_data = load_ssm_resource("resourcedatasync", additional_replacements=replacements,)

    logging.debug(resource_data)

    reference = k8s.CustomResourceReference(
        CRD_GROUP,
        CRD_VERSION,
        RESOURCE_PLURAL,
        RESOURCE_NAME,
        namespace='default',
    )

    k8s.create_custom_resource(reference, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(reference)

    assert cr is not None
    assert k8s.get_resource_exists(reference)
    yield reference, cr

    k8s.delete_custom_resource(reference)
    time.sleep(DELETE_WAIT_AFTER_SECONDS)


@service_marker
class TestResourceDataSync:
    def test_create_delete(self, resourcedatasync, ssm_client):
        (reference, _) = resourcedatasync

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
                    "sourceRegions": ["us-east-1"]
                }
            }
        }

        k8s.patch_custom_resource(reference, update_data)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)
        assert k8s.wait_on_condition(reference, "ACK.ResourceSynced", "True", wait_periods=10)

        updated_cr = k8s.get_resource(reference)
        assert updated_cr is not None
        assert 'sourceRegions' in updated_cr["spec"]["syncSource"]
        assert updated_cr["spec"]["syncSource"]["sourceRegions"] == update_data["spec"]["syncSource"]["sourceRegions"]

        # check aws api
        paginator = ssm_client.get_paginator('list_resource_data_sync')
        sync_aws = None
        
        for page in paginator.paginate(SyncType='SyncFromSource'):
            for sync in page['ResourceDataSyncItems']:
                if sync['SyncName'] == cr['spec']['syncName']:
                    sync_aws = sync
                    break
            if sync_aws:
                break

        assert sync_aws is not None
        assert sync_aws['SyncName'] == updated_cr['spec']['syncName']
        assert sync_aws['SyncType'] == updated_cr['spec']['syncType']
        assert sync_aws['SyncSource']['SourceType'] == updated_cr['spec']['syncSource']['sourceType']
        assert sync_aws['SyncSource']['SourceRegions'] == updated_cr['spec']['syncSource']['sourceRegions']
