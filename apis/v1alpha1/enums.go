// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AssociationComplianceSeverity string

const (
	AssociationComplianceSeverity_CRITICAL    AssociationComplianceSeverity = "CRITICAL"
	AssociationComplianceSeverity_HIGH        AssociationComplianceSeverity = "HIGH"
	AssociationComplianceSeverity_LOW         AssociationComplianceSeverity = "LOW"
	AssociationComplianceSeverity_MEDIUM      AssociationComplianceSeverity = "MEDIUM"
	AssociationComplianceSeverity_UNSPECIFIED AssociationComplianceSeverity = "UNSPECIFIED"
)

type AssociationExecutionFilterKey string

const (
	AssociationExecutionFilterKey_CreatedTime AssociationExecutionFilterKey = "CreatedTime"
	AssociationExecutionFilterKey_ExecutionId AssociationExecutionFilterKey = "ExecutionId"
	AssociationExecutionFilterKey_Status      AssociationExecutionFilterKey = "Status"
)

type AssociationExecutionTargetsFilterKey string

const (
	AssociationExecutionTargetsFilterKey_ResourceId   AssociationExecutionTargetsFilterKey = "ResourceId"
	AssociationExecutionTargetsFilterKey_ResourceType AssociationExecutionTargetsFilterKey = "ResourceType"
	AssociationExecutionTargetsFilterKey_Status       AssociationExecutionTargetsFilterKey = "Status"
)

type AssociationFilterKey string

const (
	AssociationFilterKey_AssociationId         AssociationFilterKey = "AssociationId"
	AssociationFilterKey_AssociationName       AssociationFilterKey = "AssociationName"
	AssociationFilterKey_AssociationStatusName AssociationFilterKey = "AssociationStatusName"
	AssociationFilterKey_InstanceId            AssociationFilterKey = "InstanceId"
	AssociationFilterKey_LastExecutedAfter     AssociationFilterKey = "LastExecutedAfter"
	AssociationFilterKey_LastExecutedBefore    AssociationFilterKey = "LastExecutedBefore"
	AssociationFilterKey_Name                  AssociationFilterKey = "Name"
	AssociationFilterKey_ResourceGroupName     AssociationFilterKey = "ResourceGroupName"
)

type AssociationFilterOperatorType string

const (
	AssociationFilterOperatorType_EQUAL        AssociationFilterOperatorType = "EQUAL"
	AssociationFilterOperatorType_GREATER_THAN AssociationFilterOperatorType = "GREATER_THAN"
	AssociationFilterOperatorType_LESS_THAN    AssociationFilterOperatorType = "LESS_THAN"
)

type AssociationStatusName string

const (
	AssociationStatusName_Failed  AssociationStatusName = "Failed"
	AssociationStatusName_Pending AssociationStatusName = "Pending"
	AssociationStatusName_Success AssociationStatusName = "Success"
)

type AssociationSyncCompliance string

const (
	AssociationSyncCompliance_AUTO   AssociationSyncCompliance = "AUTO"
	AssociationSyncCompliance_MANUAL AssociationSyncCompliance = "MANUAL"
)

type AttachmentHashType string

const (
	AttachmentHashType_Sha256 AttachmentHashType = "Sha256"
)

type AttachmentsSourceKey string

const (
	AttachmentsSourceKey_AttachmentReference AttachmentsSourceKey = "AttachmentReference"
	AttachmentsSourceKey_S3FileUrl           AttachmentsSourceKey = "S3FileUrl"
	AttachmentsSourceKey_SourceUrl           AttachmentsSourceKey = "SourceUrl"
)

type AutomationExecutionFilterKey string

const (
	AutomationExecutionFilterKey_AutomationSubtype   AutomationExecutionFilterKey = "AutomationSubtype"
	AutomationExecutionFilterKey_AutomationType      AutomationExecutionFilterKey = "AutomationType"
	AutomationExecutionFilterKey_CurrentAction       AutomationExecutionFilterKey = "CurrentAction"
	AutomationExecutionFilterKey_DocumentNamePrefix  AutomationExecutionFilterKey = "DocumentNamePrefix"
	AutomationExecutionFilterKey_ExecutionId         AutomationExecutionFilterKey = "ExecutionId"
	AutomationExecutionFilterKey_ExecutionStatus     AutomationExecutionFilterKey = "ExecutionStatus"
	AutomationExecutionFilterKey_OpsItemId           AutomationExecutionFilterKey = "OpsItemId"
	AutomationExecutionFilterKey_ParentExecutionId   AutomationExecutionFilterKey = "ParentExecutionId"
	AutomationExecutionFilterKey_StartTimeAfter      AutomationExecutionFilterKey = "StartTimeAfter"
	AutomationExecutionFilterKey_StartTimeBefore     AutomationExecutionFilterKey = "StartTimeBefore"
	AutomationExecutionFilterKey_TagKey              AutomationExecutionFilterKey = "TagKey"
	AutomationExecutionFilterKey_TargetResourceGroup AutomationExecutionFilterKey = "TargetResourceGroup"
)

type AutomationExecutionStatus string

const (
	AutomationExecutionStatus_Approved                       AutomationExecutionStatus = "Approved"
	AutomationExecutionStatus_Cancelled                      AutomationExecutionStatus = "Cancelled"
	AutomationExecutionStatus_Cancelling                     AutomationExecutionStatus = "Cancelling"
	AutomationExecutionStatus_ChangeCalendarOverrideApproved AutomationExecutionStatus = "ChangeCalendarOverrideApproved"
	AutomationExecutionStatus_ChangeCalendarOverrideRejected AutomationExecutionStatus = "ChangeCalendarOverrideRejected"
	AutomationExecutionStatus_CompletedWithFailure           AutomationExecutionStatus = "CompletedWithFailure"
	AutomationExecutionStatus_CompletedWithSuccess           AutomationExecutionStatus = "CompletedWithSuccess"
	AutomationExecutionStatus_Exited                         AutomationExecutionStatus = "Exited"
	AutomationExecutionStatus_Failed                         AutomationExecutionStatus = "Failed"
	AutomationExecutionStatus_InProgress                     AutomationExecutionStatus = "InProgress"
	AutomationExecutionStatus_Pending                        AutomationExecutionStatus = "Pending"
	AutomationExecutionStatus_PendingApproval                AutomationExecutionStatus = "PendingApproval"
	AutomationExecutionStatus_PendingChangeCalendarOverride  AutomationExecutionStatus = "PendingChangeCalendarOverride"
	AutomationExecutionStatus_Rejected                       AutomationExecutionStatus = "Rejected"
	AutomationExecutionStatus_RunbookInProgress              AutomationExecutionStatus = "RunbookInProgress"
	AutomationExecutionStatus_Scheduled                      AutomationExecutionStatus = "Scheduled"
	AutomationExecutionStatus_Success                        AutomationExecutionStatus = "Success"
	AutomationExecutionStatus_TimedOut                       AutomationExecutionStatus = "TimedOut"
	AutomationExecutionStatus_Waiting                        AutomationExecutionStatus = "Waiting"
)

type AutomationSubtype string

const (
	AutomationSubtype_ChangeRequest AutomationSubtype = "ChangeRequest"
)

type AutomationType string

const (
	AutomationType_CrossAccount AutomationType = "CrossAccount"
	AutomationType_Local        AutomationType = "Local"
)

type CalendarState string

const (
	CalendarState_CLOSED CalendarState = "CLOSED"
	CalendarState_OPEN   CalendarState = "OPEN"
)

type CommandFilterKey string

const (
	CommandFilterKey_DocumentName   CommandFilterKey = "DocumentName"
	CommandFilterKey_ExecutionStage CommandFilterKey = "ExecutionStage"
	CommandFilterKey_InvokedAfter   CommandFilterKey = "InvokedAfter"
	CommandFilterKey_InvokedBefore  CommandFilterKey = "InvokedBefore"
	CommandFilterKey_Status         CommandFilterKey = "Status"
)

type CommandInvocationStatus string

const (
	CommandInvocationStatus_Cancelled  CommandInvocationStatus = "Cancelled"
	CommandInvocationStatus_Cancelling CommandInvocationStatus = "Cancelling"
	CommandInvocationStatus_Delayed    CommandInvocationStatus = "Delayed"
	CommandInvocationStatus_Failed     CommandInvocationStatus = "Failed"
	CommandInvocationStatus_InProgress CommandInvocationStatus = "InProgress"
	CommandInvocationStatus_Pending    CommandInvocationStatus = "Pending"
	CommandInvocationStatus_Success    CommandInvocationStatus = "Success"
	CommandInvocationStatus_TimedOut   CommandInvocationStatus = "TimedOut"
)

type CommandPluginStatus string

const (
	CommandPluginStatus_Cancelled  CommandPluginStatus = "Cancelled"
	CommandPluginStatus_Failed     CommandPluginStatus = "Failed"
	CommandPluginStatus_InProgress CommandPluginStatus = "InProgress"
	CommandPluginStatus_Pending    CommandPluginStatus = "Pending"
	CommandPluginStatus_Success    CommandPluginStatus = "Success"
	CommandPluginStatus_TimedOut   CommandPluginStatus = "TimedOut"
)

type CommandStatus string

const (
	CommandStatus_Cancelled  CommandStatus = "Cancelled"
	CommandStatus_Cancelling CommandStatus = "Cancelling"
	CommandStatus_Failed     CommandStatus = "Failed"
	CommandStatus_InProgress CommandStatus = "InProgress"
	CommandStatus_Pending    CommandStatus = "Pending"
	CommandStatus_Success    CommandStatus = "Success"
	CommandStatus_TimedOut   CommandStatus = "TimedOut"
)

type ComplianceQueryOperatorType string

const (
	ComplianceQueryOperatorType_BEGIN_WITH   ComplianceQueryOperatorType = "BEGIN_WITH"
	ComplianceQueryOperatorType_EQUAL        ComplianceQueryOperatorType = "EQUAL"
	ComplianceQueryOperatorType_GREATER_THAN ComplianceQueryOperatorType = "GREATER_THAN"
	ComplianceQueryOperatorType_LESS_THAN    ComplianceQueryOperatorType = "LESS_THAN"
	ComplianceQueryOperatorType_NOT_EQUAL    ComplianceQueryOperatorType = "NOT_EQUAL"
)

type ComplianceSeverity string

const (
	ComplianceSeverity_CRITICAL      ComplianceSeverity = "CRITICAL"
	ComplianceSeverity_HIGH          ComplianceSeverity = "HIGH"
	ComplianceSeverity_INFORMATIONAL ComplianceSeverity = "INFORMATIONAL"
	ComplianceSeverity_LOW           ComplianceSeverity = "LOW"
	ComplianceSeverity_MEDIUM        ComplianceSeverity = "MEDIUM"
	ComplianceSeverity_UNSPECIFIED   ComplianceSeverity = "UNSPECIFIED"
)

type ComplianceStatus string

const (
	ComplianceStatus_COMPLIANT     ComplianceStatus = "COMPLIANT"
	ComplianceStatus_NON_COMPLIANT ComplianceStatus = "NON_COMPLIANT"
)

type ComplianceUploadType string

const (
	ComplianceUploadType_COMPLETE ComplianceUploadType = "COMPLETE"
	ComplianceUploadType_PARTIAL  ComplianceUploadType = "PARTIAL"
)

type ConnectionStatus string

const (
	ConnectionStatus_connected    ConnectionStatus = "connected"
	ConnectionStatus_notconnected ConnectionStatus = "notconnected"
)

type DescribeActivationsFilterKeys string

const (
	DescribeActivationsFilterKeys_ActivationIds       DescribeActivationsFilterKeys = "ActivationIds"
	DescribeActivationsFilterKeys_DefaultInstanceName DescribeActivationsFilterKeys = "DefaultInstanceName"
	DescribeActivationsFilterKeys_IamRole             DescribeActivationsFilterKeys = "IamRole"
)

type DocumentFilterKey string

const (
	DocumentFilterKey_DocumentType  DocumentFilterKey = "DocumentType"
	DocumentFilterKey_Name          DocumentFilterKey = "Name"
	DocumentFilterKey_Owner         DocumentFilterKey = "Owner"
	DocumentFilterKey_PlatformTypes DocumentFilterKey = "PlatformTypes"
)

type DocumentFormat string

const (
	DocumentFormat_JSON DocumentFormat = "JSON"
	DocumentFormat_TEXT DocumentFormat = "TEXT"
	DocumentFormat_YAML DocumentFormat = "YAML"
)

type DocumentHashType string

const (
	DocumentHashType_Sha1   DocumentHashType = "Sha1"
	DocumentHashType_Sha256 DocumentHashType = "Sha256"
)

type DocumentMetadataEnum string

const (
	DocumentMetadataEnum_DocumentReviews DocumentMetadataEnum = "DocumentReviews"
)

type DocumentParameterType string

const (
	DocumentParameterType_String     DocumentParameterType = "String"
	DocumentParameterType_StringList DocumentParameterType = "StringList"
)

type DocumentPermissionType string

const (
	DocumentPermissionType_Share DocumentPermissionType = "Share"
)

type DocumentReviewAction string

const (
	DocumentReviewAction_Approve       DocumentReviewAction = "Approve"
	DocumentReviewAction_Reject        DocumentReviewAction = "Reject"
	DocumentReviewAction_SendForReview DocumentReviewAction = "SendForReview"
	DocumentReviewAction_UpdateReview  DocumentReviewAction = "UpdateReview"
)

type DocumentReviewCommentType string

const (
	DocumentReviewCommentType_Comment DocumentReviewCommentType = "Comment"
)

type DocumentStatus_SDK string

const (
	DocumentStatus_SDK_Active   DocumentStatus_SDK = "Active"
	DocumentStatus_SDK_Creating DocumentStatus_SDK = "Creating"
	DocumentStatus_SDK_Deleting DocumentStatus_SDK = "Deleting"
	DocumentStatus_SDK_Failed   DocumentStatus_SDK = "Failed"
	DocumentStatus_SDK_Updating DocumentStatus_SDK = "Updating"
)

type DocumentType string

const (
	DocumentType_ApplicationConfiguration       DocumentType = "ApplicationConfiguration"
	DocumentType_ApplicationConfigurationSchema DocumentType = "ApplicationConfigurationSchema"
	DocumentType_Automation                     DocumentType = "Automation"
	DocumentType_Automation_ChangeTemplate      DocumentType = "Automation.ChangeTemplate"
	DocumentType_ChangeCalendar                 DocumentType = "ChangeCalendar"
	DocumentType_CloudFormation                 DocumentType = "CloudFormation"
	DocumentType_Command                        DocumentType = "Command"
	DocumentType_ConformancePackTemplate        DocumentType = "ConformancePackTemplate"
	DocumentType_DeploymentStrategy             DocumentType = "DeploymentStrategy"
	DocumentType_Package                        DocumentType = "Package"
	DocumentType_Policy                         DocumentType = "Policy"
	DocumentType_ProblemAnalysis                DocumentType = "ProblemAnalysis"
	DocumentType_ProblemAnalysisTemplate        DocumentType = "ProblemAnalysisTemplate"
	DocumentType_QuickSetup                     DocumentType = "QuickSetup"
	DocumentType_Session                        DocumentType = "Session"
)

type ExecutionMode string

const (
	ExecutionMode_Auto        ExecutionMode = "Auto"
	ExecutionMode_Interactive ExecutionMode = "Interactive"
)

type ExecutionPreviewStatus string

const (
	ExecutionPreviewStatus_Failed     ExecutionPreviewStatus = "Failed"
	ExecutionPreviewStatus_InProgress ExecutionPreviewStatus = "InProgress"
	ExecutionPreviewStatus_Pending    ExecutionPreviewStatus = "Pending"
	ExecutionPreviewStatus_Success    ExecutionPreviewStatus = "Success"
)

type ExternalAlarmState string

const (
	ExternalAlarmState_ALARM   ExternalAlarmState = "ALARM"
	ExternalAlarmState_UNKNOWN ExternalAlarmState = "UNKNOWN"
)

type Fault string

const (
	Fault_Client  Fault = "Client"
	Fault_Server  Fault = "Server"
	Fault_Unknown Fault = "Unknown"
)

type ImpactType string

const (
	ImpactType_Mutating     ImpactType = "Mutating"
	ImpactType_NonMutating  ImpactType = "NonMutating"
	ImpactType_Undetermined ImpactType = "Undetermined"
)

type InstanceInformationFilterKey string

const (
	InstanceInformationFilterKey_ActivationIds     InstanceInformationFilterKey = "ActivationIds"
	InstanceInformationFilterKey_AgentVersion      InstanceInformationFilterKey = "AgentVersion"
	InstanceInformationFilterKey_AssociationStatus InstanceInformationFilterKey = "AssociationStatus"
	InstanceInformationFilterKey_IamRole           InstanceInformationFilterKey = "IamRole"
	InstanceInformationFilterKey_InstanceIds       InstanceInformationFilterKey = "InstanceIds"
	InstanceInformationFilterKey_PingStatus        InstanceInformationFilterKey = "PingStatus"
	InstanceInformationFilterKey_PlatformTypes     InstanceInformationFilterKey = "PlatformTypes"
	InstanceInformationFilterKey_ResourceType      InstanceInformationFilterKey = "ResourceType"
)

type InstancePatchStateOperatorType string

const (
	InstancePatchStateOperatorType_Equal       InstancePatchStateOperatorType = "Equal"
	InstancePatchStateOperatorType_GreaterThan InstancePatchStateOperatorType = "GreaterThan"
	InstancePatchStateOperatorType_LessThan    InstancePatchStateOperatorType = "LessThan"
	InstancePatchStateOperatorType_NotEqual    InstancePatchStateOperatorType = "NotEqual"
)

type InstancePropertyFilterKey string

const (
	InstancePropertyFilterKey_ActivationIds     InstancePropertyFilterKey = "ActivationIds"
	InstancePropertyFilterKey_AgentVersion      InstancePropertyFilterKey = "AgentVersion"
	InstancePropertyFilterKey_AssociationStatus InstancePropertyFilterKey = "AssociationStatus"
	InstancePropertyFilterKey_DocumentName      InstancePropertyFilterKey = "DocumentName"
	InstancePropertyFilterKey_IamRole           InstancePropertyFilterKey = "IamRole"
	InstancePropertyFilterKey_InstanceIds       InstancePropertyFilterKey = "InstanceIds"
	InstancePropertyFilterKey_PingStatus        InstancePropertyFilterKey = "PingStatus"
	InstancePropertyFilterKey_PlatformTypes     InstancePropertyFilterKey = "PlatformTypes"
	InstancePropertyFilterKey_ResourceType      InstancePropertyFilterKey = "ResourceType"
)

type InstancePropertyFilterOperator string

const (
	InstancePropertyFilterOperator_BeginWith   InstancePropertyFilterOperator = "BeginWith"
	InstancePropertyFilterOperator_Equal       InstancePropertyFilterOperator = "Equal"
	InstancePropertyFilterOperator_GreaterThan InstancePropertyFilterOperator = "GreaterThan"
	InstancePropertyFilterOperator_LessThan    InstancePropertyFilterOperator = "LessThan"
	InstancePropertyFilterOperator_NotEqual    InstancePropertyFilterOperator = "NotEqual"
)

type InventoryAttributeDataType string

const (
	InventoryAttributeDataType_number InventoryAttributeDataType = "number"
	InventoryAttributeDataType_string InventoryAttributeDataType = "string"
)

type InventoryDeletionStatus string

const (
	InventoryDeletionStatus_Complete   InventoryDeletionStatus = "Complete"
	InventoryDeletionStatus_InProgress InventoryDeletionStatus = "InProgress"
)

type InventoryQueryOperatorType string

const (
	InventoryQueryOperatorType_BeginWith   InventoryQueryOperatorType = "BeginWith"
	InventoryQueryOperatorType_Equal       InventoryQueryOperatorType = "Equal"
	InventoryQueryOperatorType_Exists      InventoryQueryOperatorType = "Exists"
	InventoryQueryOperatorType_GreaterThan InventoryQueryOperatorType = "GreaterThan"
	InventoryQueryOperatorType_LessThan    InventoryQueryOperatorType = "LessThan"
	InventoryQueryOperatorType_NotEqual    InventoryQueryOperatorType = "NotEqual"
)

type InventorySchemaDeleteOption string

const (
	InventorySchemaDeleteOption_DeleteSchema  InventorySchemaDeleteOption = "DeleteSchema"
	InventorySchemaDeleteOption_DisableSchema InventorySchemaDeleteOption = "DisableSchema"
)

type LastResourceDataSyncStatus string

const (
	LastResourceDataSyncStatus_Failed     LastResourceDataSyncStatus = "Failed"
	LastResourceDataSyncStatus_InProgress LastResourceDataSyncStatus = "InProgress"
	LastResourceDataSyncStatus_Successful LastResourceDataSyncStatus = "Successful"
)

type MaintenanceWindowExecutionStatus string

const (
	MaintenanceWindowExecutionStatus_CANCELLED           MaintenanceWindowExecutionStatus = "CANCELLED"
	MaintenanceWindowExecutionStatus_CANCELLING          MaintenanceWindowExecutionStatus = "CANCELLING"
	MaintenanceWindowExecutionStatus_FAILED              MaintenanceWindowExecutionStatus = "FAILED"
	MaintenanceWindowExecutionStatus_IN_PROGRESS         MaintenanceWindowExecutionStatus = "IN_PROGRESS"
	MaintenanceWindowExecutionStatus_PENDING             MaintenanceWindowExecutionStatus = "PENDING"
	MaintenanceWindowExecutionStatus_SKIPPED_OVERLAPPING MaintenanceWindowExecutionStatus = "SKIPPED_OVERLAPPING"
	MaintenanceWindowExecutionStatus_SUCCESS             MaintenanceWindowExecutionStatus = "SUCCESS"
	MaintenanceWindowExecutionStatus_TIMED_OUT           MaintenanceWindowExecutionStatus = "TIMED_OUT"
)

type MaintenanceWindowResourceType string

const (
	MaintenanceWindowResourceType_INSTANCE       MaintenanceWindowResourceType = "INSTANCE"
	MaintenanceWindowResourceType_RESOURCE_GROUP MaintenanceWindowResourceType = "RESOURCE_GROUP"
)

type MaintenanceWindowTaskCutoffBehavior string

const (
	MaintenanceWindowTaskCutoffBehavior_CANCEL_TASK   MaintenanceWindowTaskCutoffBehavior = "CANCEL_TASK"
	MaintenanceWindowTaskCutoffBehavior_CONTINUE_TASK MaintenanceWindowTaskCutoffBehavior = "CONTINUE_TASK"
)

type MaintenanceWindowTaskType string

const (
	MaintenanceWindowTaskType_AUTOMATION     MaintenanceWindowTaskType = "AUTOMATION"
	MaintenanceWindowTaskType_LAMBDA         MaintenanceWindowTaskType = "LAMBDA"
	MaintenanceWindowTaskType_RUN_COMMAND    MaintenanceWindowTaskType = "RUN_COMMAND"
	MaintenanceWindowTaskType_STEP_FUNCTIONS MaintenanceWindowTaskType = "STEP_FUNCTIONS"
)

type ManagedStatus string

const (
	ManagedStatus_All       ManagedStatus = "All"
	ManagedStatus_Managed   ManagedStatus = "Managed"
	ManagedStatus_Unmanaged ManagedStatus = "Unmanaged"
)

type NodeAggregatorType string

const (
	NodeAggregatorType_Count NodeAggregatorType = "Count"
)

type NodeAttributeName string

const (
	NodeAttributeName_AgentVersion    NodeAttributeName = "AgentVersion"
	NodeAttributeName_PlatformName    NodeAttributeName = "PlatformName"
	NodeAttributeName_PlatformType    NodeAttributeName = "PlatformType"
	NodeAttributeName_PlatformVersion NodeAttributeName = "PlatformVersion"
	NodeAttributeName_Region          NodeAttributeName = "Region"
	NodeAttributeName_ResourceType    NodeAttributeName = "ResourceType"
)

type NodeFilterKey string

const (
	NodeFilterKey_AccountId              NodeFilterKey = "AccountId"
	NodeFilterKey_AgentType              NodeFilterKey = "AgentType"
	NodeFilterKey_AgentVersion           NodeFilterKey = "AgentVersion"
	NodeFilterKey_ComputerName           NodeFilterKey = "ComputerName"
	NodeFilterKey_InstanceId             NodeFilterKey = "InstanceId"
	NodeFilterKey_InstanceStatus         NodeFilterKey = "InstanceStatus"
	NodeFilterKey_IpAddress              NodeFilterKey = "IpAddress"
	NodeFilterKey_ManagedStatus          NodeFilterKey = "ManagedStatus"
	NodeFilterKey_OrganizationalUnitId   NodeFilterKey = "OrganizationalUnitId"
	NodeFilterKey_OrganizationalUnitPath NodeFilterKey = "OrganizationalUnitPath"
	NodeFilterKey_PlatformName           NodeFilterKey = "PlatformName"
	NodeFilterKey_PlatformType           NodeFilterKey = "PlatformType"
	NodeFilterKey_PlatformVersion        NodeFilterKey = "PlatformVersion"
	NodeFilterKey_Region                 NodeFilterKey = "Region"
	NodeFilterKey_ResourceType           NodeFilterKey = "ResourceType"
)

type NodeFilterOperatorType string

const (
	NodeFilterOperatorType_BeginWith NodeFilterOperatorType = "BeginWith"
	NodeFilterOperatorType_Equal     NodeFilterOperatorType = "Equal"
	NodeFilterOperatorType_NotEqual  NodeFilterOperatorType = "NotEqual"
)

type NodeTypeName string

const (
	NodeTypeName_Instance NodeTypeName = "Instance"
)

type NotificationEvent string

const (
	NotificationEvent_All        NotificationEvent = "All"
	NotificationEvent_Cancelled  NotificationEvent = "Cancelled"
	NotificationEvent_Failed     NotificationEvent = "Failed"
	NotificationEvent_InProgress NotificationEvent = "InProgress"
	NotificationEvent_Success    NotificationEvent = "Success"
	NotificationEvent_TimedOut   NotificationEvent = "TimedOut"
)

type NotificationType string

const (
	NotificationType_Command    NotificationType = "Command"
	NotificationType_Invocation NotificationType = "Invocation"
)

type OperatingSystem string

const (
	OperatingSystem_ALMA_LINUX              OperatingSystem = "ALMA_LINUX"
	OperatingSystem_AMAZON_LINUX            OperatingSystem = "AMAZON_LINUX"
	OperatingSystem_AMAZON_LINUX_2          OperatingSystem = "AMAZON_LINUX_2"
	OperatingSystem_AMAZON_LINUX_2022       OperatingSystem = "AMAZON_LINUX_2022"
	OperatingSystem_AMAZON_LINUX_2023       OperatingSystem = "AMAZON_LINUX_2023"
	OperatingSystem_CENTOS                  OperatingSystem = "CENTOS"
	OperatingSystem_DEBIAN                  OperatingSystem = "DEBIAN"
	OperatingSystem_MACOS                   OperatingSystem = "MACOS"
	OperatingSystem_ORACLE_LINUX            OperatingSystem = "ORACLE_LINUX"
	OperatingSystem_RASPBIAN                OperatingSystem = "RASPBIAN"
	OperatingSystem_REDHAT_ENTERPRISE_LINUX OperatingSystem = "REDHAT_ENTERPRISE_LINUX"
	OperatingSystem_ROCKY_LINUX             OperatingSystem = "ROCKY_LINUX"
	OperatingSystem_SUSE                    OperatingSystem = "SUSE"
	OperatingSystem_UBUNTU                  OperatingSystem = "UBUNTU"
	OperatingSystem_WINDOWS                 OperatingSystem = "WINDOWS"
)

type OpsFilterOperatorType string

const (
	OpsFilterOperatorType_BeginWith   OpsFilterOperatorType = "BeginWith"
	OpsFilterOperatorType_Equal       OpsFilterOperatorType = "Equal"
	OpsFilterOperatorType_Exists      OpsFilterOperatorType = "Exists"
	OpsFilterOperatorType_GreaterThan OpsFilterOperatorType = "GreaterThan"
	OpsFilterOperatorType_LessThan    OpsFilterOperatorType = "LessThan"
	OpsFilterOperatorType_NotEqual    OpsFilterOperatorType = "NotEqual"
)

type OpsItemDataType string

const (
	OpsItemDataType_SearchableString OpsItemDataType = "SearchableString"
	OpsItemDataType_String           OpsItemDataType = "String"
)

type OpsItemEventFilterKey string

const (
	OpsItemEventFilterKey_OpsItemId OpsItemEventFilterKey = "OpsItemId"
)

type OpsItemEventFilterOperator string

const (
	OpsItemEventFilterOperator_Equal OpsItemEventFilterOperator = "Equal"
)

type OpsItemFilterKey string

const (
	OpsItemFilterKey_AccountId                           OpsItemFilterKey = "AccountId"
	OpsItemFilterKey_ActualEndTime                       OpsItemFilterKey = "ActualEndTime"
	OpsItemFilterKey_ActualStartTime                     OpsItemFilterKey = "ActualStartTime"
	OpsItemFilterKey_AutomationId                        OpsItemFilterKey = "AutomationId"
	OpsItemFilterKey_Category                            OpsItemFilterKey = "Category"
	OpsItemFilterKey_ChangeRequestByApproverArn          OpsItemFilterKey = "ChangeRequestByApproverArn"
	OpsItemFilterKey_ChangeRequestByApproverName         OpsItemFilterKey = "ChangeRequestByApproverName"
	OpsItemFilterKey_ChangeRequestByRequesterArn         OpsItemFilterKey = "ChangeRequestByRequesterArn"
	OpsItemFilterKey_ChangeRequestByRequesterName        OpsItemFilterKey = "ChangeRequestByRequesterName"
	OpsItemFilterKey_ChangeRequestByTargetsResourceGroup OpsItemFilterKey = "ChangeRequestByTargetsResourceGroup"
	OpsItemFilterKey_ChangeRequestByTemplate             OpsItemFilterKey = "ChangeRequestByTemplate"
	OpsItemFilterKey_CreatedBy                           OpsItemFilterKey = "CreatedBy"
	OpsItemFilterKey_CreatedTime                         OpsItemFilterKey = "CreatedTime"
	OpsItemFilterKey_InsightByType                       OpsItemFilterKey = "InsightByType"
	OpsItemFilterKey_LastModifiedTime                    OpsItemFilterKey = "LastModifiedTime"
	OpsItemFilterKey_OperationalData                     OpsItemFilterKey = "OperationalData"
	OpsItemFilterKey_OperationalDataKey                  OpsItemFilterKey = "OperationalDataKey"
	OpsItemFilterKey_OperationalDataValue                OpsItemFilterKey = "OperationalDataValue"
	OpsItemFilterKey_OpsItemId                           OpsItemFilterKey = "OpsItemId"
	OpsItemFilterKey_OpsItemType                         OpsItemFilterKey = "OpsItemType"
	OpsItemFilterKey_PlannedEndTime                      OpsItemFilterKey = "PlannedEndTime"
	OpsItemFilterKey_PlannedStartTime                    OpsItemFilterKey = "PlannedStartTime"
	OpsItemFilterKey_Priority                            OpsItemFilterKey = "Priority"
	OpsItemFilterKey_ResourceId                          OpsItemFilterKey = "ResourceId"
	OpsItemFilterKey_Severity                            OpsItemFilterKey = "Severity"
	OpsItemFilterKey_Source                              OpsItemFilterKey = "Source"
	OpsItemFilterKey_Status                              OpsItemFilterKey = "Status"
	OpsItemFilterKey_Title                               OpsItemFilterKey = "Title"
)

type OpsItemFilterOperator string

const (
	OpsItemFilterOperator_Contains    OpsItemFilterOperator = "Contains"
	OpsItemFilterOperator_Equal       OpsItemFilterOperator = "Equal"
	OpsItemFilterOperator_GreaterThan OpsItemFilterOperator = "GreaterThan"
	OpsItemFilterOperator_LessThan    OpsItemFilterOperator = "LessThan"
)

type OpsItemRelatedItemsFilterKey string

const (
	OpsItemRelatedItemsFilterKey_AssociationId OpsItemRelatedItemsFilterKey = "AssociationId"
	OpsItemRelatedItemsFilterKey_ResourceType  OpsItemRelatedItemsFilterKey = "ResourceType"
	OpsItemRelatedItemsFilterKey_ResourceUri   OpsItemRelatedItemsFilterKey = "ResourceUri"
)

type OpsItemRelatedItemsFilterOperator string

const (
	OpsItemRelatedItemsFilterOperator_Equal OpsItemRelatedItemsFilterOperator = "Equal"
)

type OpsItemStatus string

const (
	OpsItemStatus_Approved                       OpsItemStatus = "Approved"
	OpsItemStatus_Cancelled                      OpsItemStatus = "Cancelled"
	OpsItemStatus_Cancelling                     OpsItemStatus = "Cancelling"
	OpsItemStatus_ChangeCalendarOverrideApproved OpsItemStatus = "ChangeCalendarOverrideApproved"
	OpsItemStatus_ChangeCalendarOverrideRejected OpsItemStatus = "ChangeCalendarOverrideRejected"
	OpsItemStatus_Closed                         OpsItemStatus = "Closed"
	OpsItemStatus_CompletedWithFailure           OpsItemStatus = "CompletedWithFailure"
	OpsItemStatus_CompletedWithSuccess           OpsItemStatus = "CompletedWithSuccess"
	OpsItemStatus_Failed                         OpsItemStatus = "Failed"
	OpsItemStatus_InProgress                     OpsItemStatus = "InProgress"
	OpsItemStatus_Open                           OpsItemStatus = "Open"
	OpsItemStatus_Pending                        OpsItemStatus = "Pending"
	OpsItemStatus_PendingApproval                OpsItemStatus = "PendingApproval"
	OpsItemStatus_PendingChangeCalendarOverride  OpsItemStatus = "PendingChangeCalendarOverride"
	OpsItemStatus_Rejected                       OpsItemStatus = "Rejected"
	OpsItemStatus_Resolved                       OpsItemStatus = "Resolved"
	OpsItemStatus_RunbookInProgress              OpsItemStatus = "RunbookInProgress"
	OpsItemStatus_Scheduled                      OpsItemStatus = "Scheduled"
	OpsItemStatus_TimedOut                       OpsItemStatus = "TimedOut"
)

type ParameterTier string

const (
	ParameterTier_Advanced            ParameterTier = "Advanced"
	ParameterTier_Intelligent_Tiering ParameterTier = "Intelligent-Tiering"
	ParameterTier_Standard            ParameterTier = "Standard"
)

type ParameterType string

const (
	ParameterType_SecureString ParameterType = "SecureString"
	ParameterType_String       ParameterType = "String"
	ParameterType_StringList   ParameterType = "StringList"
)

type ParametersFilterKey string

const (
	ParametersFilterKey_KeyId ParametersFilterKey = "KeyId"
	ParametersFilterKey_Name  ParametersFilterKey = "Name"
	ParametersFilterKey_Type  ParametersFilterKey = "Type"
)

type PatchAction string

const (
	PatchAction_ALLOW_AS_DEPENDENCY PatchAction = "ALLOW_AS_DEPENDENCY"
	PatchAction_BLOCK               PatchAction = "BLOCK"
)

type PatchComplianceDataState string

const (
	PatchComplianceDataState_FAILED                   PatchComplianceDataState = "FAILED"
	PatchComplianceDataState_INSTALLED                PatchComplianceDataState = "INSTALLED"
	PatchComplianceDataState_INSTALLED_OTHER          PatchComplianceDataState = "INSTALLED_OTHER"
	PatchComplianceDataState_INSTALLED_PENDING_REBOOT PatchComplianceDataState = "INSTALLED_PENDING_REBOOT"
	PatchComplianceDataState_INSTALLED_REJECTED       PatchComplianceDataState = "INSTALLED_REJECTED"
	PatchComplianceDataState_MISSING                  PatchComplianceDataState = "MISSING"
	PatchComplianceDataState_NOT_APPLICABLE           PatchComplianceDataState = "NOT_APPLICABLE"
)

type PatchComplianceLevel string

const (
	PatchComplianceLevel_CRITICAL      PatchComplianceLevel = "CRITICAL"
	PatchComplianceLevel_HIGH          PatchComplianceLevel = "HIGH"
	PatchComplianceLevel_INFORMATIONAL PatchComplianceLevel = "INFORMATIONAL"
	PatchComplianceLevel_LOW           PatchComplianceLevel = "LOW"
	PatchComplianceLevel_MEDIUM        PatchComplianceLevel = "MEDIUM"
	PatchComplianceLevel_UNSPECIFIED   PatchComplianceLevel = "UNSPECIFIED"
)

type PatchDeploymentStatus string

const (
	PatchDeploymentStatus_APPROVED          PatchDeploymentStatus = "APPROVED"
	PatchDeploymentStatus_EXPLICIT_APPROVED PatchDeploymentStatus = "EXPLICIT_APPROVED"
	PatchDeploymentStatus_EXPLICIT_REJECTED PatchDeploymentStatus = "EXPLICIT_REJECTED"
	PatchDeploymentStatus_PENDING_APPROVAL  PatchDeploymentStatus = "PENDING_APPROVAL"
)

type PatchFilterKey string

const (
	PatchFilterKey_ADVISORY_ID    PatchFilterKey = "ADVISORY_ID"
	PatchFilterKey_ARCH           PatchFilterKey = "ARCH"
	PatchFilterKey_BUGZILLA_ID    PatchFilterKey = "BUGZILLA_ID"
	PatchFilterKey_CLASSIFICATION PatchFilterKey = "CLASSIFICATION"
	PatchFilterKey_CVE_ID         PatchFilterKey = "CVE_ID"
	PatchFilterKey_EPOCH          PatchFilterKey = "EPOCH"
	PatchFilterKey_MSRC_SEVERITY  PatchFilterKey = "MSRC_SEVERITY"
	PatchFilterKey_NAME           PatchFilterKey = "NAME"
	PatchFilterKey_PATCH_ID       PatchFilterKey = "PATCH_ID"
	PatchFilterKey_PATCH_SET      PatchFilterKey = "PATCH_SET"
	PatchFilterKey_PRIORITY       PatchFilterKey = "PRIORITY"
	PatchFilterKey_PRODUCT        PatchFilterKey = "PRODUCT"
	PatchFilterKey_PRODUCT_FAMILY PatchFilterKey = "PRODUCT_FAMILY"
	PatchFilterKey_RELEASE        PatchFilterKey = "RELEASE"
	PatchFilterKey_REPOSITORY     PatchFilterKey = "REPOSITORY"
	PatchFilterKey_SECTION        PatchFilterKey = "SECTION"
	PatchFilterKey_SECURITY       PatchFilterKey = "SECURITY"
	PatchFilterKey_SEVERITY       PatchFilterKey = "SEVERITY"
	PatchFilterKey_VERSION        PatchFilterKey = "VERSION"
)

type PatchOperationType string

const (
	PatchOperationType_Install PatchOperationType = "Install"
	PatchOperationType_Scan    PatchOperationType = "Scan"
)

type PatchProperty string

const (
	PatchProperty_CLASSIFICATION PatchProperty = "CLASSIFICATION"
	PatchProperty_MSRC_SEVERITY  PatchProperty = "MSRC_SEVERITY"
	PatchProperty_PRIORITY       PatchProperty = "PRIORITY"
	PatchProperty_PRODUCT        PatchProperty = "PRODUCT"
	PatchProperty_PRODUCT_FAMILY PatchProperty = "PRODUCT_FAMILY"
	PatchProperty_SEVERITY       PatchProperty = "SEVERITY"
)

type PatchSet string

const (
	PatchSet_APPLICATION PatchSet = "APPLICATION"
	PatchSet_OS          PatchSet = "OS"
)

type PingStatus string

const (
	PingStatus_ConnectionLost PingStatus = "ConnectionLost"
	PingStatus_Inactive       PingStatus = "Inactive"
	PingStatus_Online         PingStatus = "Online"
)

type PlatformType string

const (
	PlatformType_Linux   PlatformType = "Linux"
	PlatformType_MacOS   PlatformType = "MacOS"
	PlatformType_Windows PlatformType = "Windows"
)

type RebootOption string

const (
	RebootOption_NoReboot       RebootOption = "NoReboot"
	RebootOption_RebootIfNeeded RebootOption = "RebootIfNeeded"
)

type ResourceDataSyncS3Format string

const (
	ResourceDataSyncS3Format_JsonSerDe ResourceDataSyncS3Format = "JsonSerDe"
)

type ResourceType string

const (
	ResourceType_EC2Instance     ResourceType = "EC2Instance"
	ResourceType_ManagedInstance ResourceType = "ManagedInstance"
)

type ResourceTypeForTagging string

const (
	ResourceTypeForTagging_Association       ResourceTypeForTagging = "Association"
	ResourceTypeForTagging_Automation        ResourceTypeForTagging = "Automation"
	ResourceTypeForTagging_Document          ResourceTypeForTagging = "Document"
	ResourceTypeForTagging_MaintenanceWindow ResourceTypeForTagging = "MaintenanceWindow"
	ResourceTypeForTagging_ManagedInstance   ResourceTypeForTagging = "ManagedInstance"
	ResourceTypeForTagging_OpsItem           ResourceTypeForTagging = "OpsItem"
	ResourceTypeForTagging_OpsMetadata       ResourceTypeForTagging = "OpsMetadata"
	ResourceTypeForTagging_Parameter         ResourceTypeForTagging = "Parameter"
	ResourceTypeForTagging_PatchBaseline     ResourceTypeForTagging = "PatchBaseline"
)

type ReviewStatus string

const (
	ReviewStatus_APPROVED     ReviewStatus = "APPROVED"
	ReviewStatus_NOT_REVIEWED ReviewStatus = "NOT_REVIEWED"
	ReviewStatus_PENDING      ReviewStatus = "PENDING"
	ReviewStatus_REJECTED     ReviewStatus = "REJECTED"
)

type SessionFilterKey string

const (
	SessionFilterKey_InvokedAfter  SessionFilterKey = "InvokedAfter"
	SessionFilterKey_InvokedBefore SessionFilterKey = "InvokedBefore"
	SessionFilterKey_Owner         SessionFilterKey = "Owner"
	SessionFilterKey_SessionId     SessionFilterKey = "SessionId"
	SessionFilterKey_Status        SessionFilterKey = "Status"
	SessionFilterKey_Target        SessionFilterKey = "Target"
)

type SessionState string

const (
	SessionState_Active  SessionState = "Active"
	SessionState_History SessionState = "History"
)

type SessionStatus string

const (
	SessionStatus_Connected    SessionStatus = "Connected"
	SessionStatus_Connecting   SessionStatus = "Connecting"
	SessionStatus_Disconnected SessionStatus = "Disconnected"
	SessionStatus_Failed       SessionStatus = "Failed"
	SessionStatus_Terminated   SessionStatus = "Terminated"
	SessionStatus_Terminating  SessionStatus = "Terminating"
)

type SignalType string

const (
	SignalType_Approve   SignalType = "Approve"
	SignalType_Reject    SignalType = "Reject"
	SignalType_Resume    SignalType = "Resume"
	SignalType_StartStep SignalType = "StartStep"
	SignalType_StopStep  SignalType = "StopStep"
)

type SourceType string

const (
	SourceType_AWS__EC2__Instance        SourceType = "AWS::EC2::Instance"
	SourceType_AWS__IoT__Thing           SourceType = "AWS::IoT::Thing"
	SourceType_AWS__SSM__ManagedInstance SourceType = "AWS::SSM::ManagedInstance"
)

type StepExecutionFilterKey string

const (
	StepExecutionFilterKey_Action                  StepExecutionFilterKey = "Action"
	StepExecutionFilterKey_ParentStepExecutionId   StepExecutionFilterKey = "ParentStepExecutionId"
	StepExecutionFilterKey_ParentStepIteration     StepExecutionFilterKey = "ParentStepIteration"
	StepExecutionFilterKey_ParentStepIteratorValue StepExecutionFilterKey = "ParentStepIteratorValue"
	StepExecutionFilterKey_StartTimeAfter          StepExecutionFilterKey = "StartTimeAfter"
	StepExecutionFilterKey_StartTimeBefore         StepExecutionFilterKey = "StartTimeBefore"
	StepExecutionFilterKey_StepExecutionId         StepExecutionFilterKey = "StepExecutionId"
	StepExecutionFilterKey_StepExecutionStatus     StepExecutionFilterKey = "StepExecutionStatus"
	StepExecutionFilterKey_StepName                StepExecutionFilterKey = "StepName"
)

type StopType string

const (
	StopType_Cancel   StopType = "Cancel"
	StopType_Complete StopType = "Complete"
)
