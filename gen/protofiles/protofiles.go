// Code generated by go generate; DO NOT EDIT.
// This file contains the list of public variables with the specified prefix.

package protofiles

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	gen_commonfate_access_provisioner_v1alpha1_provisioner_pb_go "github.com/common-fate/sdk/gen/commonfate/access/provisioner/v1alpha1"
	gen_commonfate_access_v1alpha1_access_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_access_request_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_audit_logs_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_diagnostic_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_extension_conditions_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_grant_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_grants_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_identity_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_integration_audit_logs_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_named_eid_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_proxy_session_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_resource_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_access_v1alpha1_user_pb_go "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	gen_commonfate_authz_v1alpha1_action_pb_go "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	gen_commonfate_authz_v1alpha1_authz_pb_go "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	gen_commonfate_authz_v1alpha1_evaluation_pb_go "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	gen_commonfate_authz_v1alpha1_health_pb_go "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	gen_commonfate_authz_v1alpha1_policy_pb_go "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	gen_commonfate_authz_v1alpha1_read_only_pb_go "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	gen_commonfate_authz_v1alpha1_request_pb_go "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	gen_commonfate_authz_v1alpha1_schema_pb_go "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	gen_commonfate_authz_v1alpha1_validation_pb_go "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	gen_commonfate_control_attest_v1alpha1_attest_pb_go "github.com/common-fate/sdk/gen/commonfate/control/attest/v1alpha1"
	gen_commonfate_control_attest_v1alpha1_attestation_pb_go "github.com/common-fate/sdk/gen/commonfate/control/attest/v1alpha1"
	gen_commonfate_control_config_v1alpha1_access_workflow_pb_go "github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1"
	gen_commonfate_control_config_v1alpha1_availability_spec_pb_go "github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1"
	gen_commonfate_control_config_v1alpha1_aws_resource_scanner_pb_go "github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1"
	gen_commonfate_control_config_v1alpha1_deployment_pb_go "github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1"
	gen_commonfate_control_config_v1alpha1_gcp_role_group_pb_go "github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1"
	gen_commonfate_control_config_v1alpha1_selector_pb_go "github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1"
	gen_commonfate_control_config_v1alpha1_slack_alert_pb_go "github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1"
	gen_commonfate_control_config_v1alpha1_webhook_provisioner_pb_go "github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1"
	gen_commonfate_control_diagnostic_v1alpha1_diagnostic_pb_go "github.com/common-fate/sdk/gen/commonfate/control/diagnostic/v1alpha1"
	gen_commonfate_control_directory_v1alpha1_directory_pb_go "github.com/common-fate/sdk/gen/commonfate/control/directory/v1alpha1"
	gen_commonfate_control_directory_v1alpha1_group_pb_go "github.com/common-fate/sdk/gen/commonfate/control/directory/v1alpha1"
	gen_commonfate_control_directory_v1alpha1_integration_pb_go "github.com/common-fate/sdk/gen/commonfate/control/directory/v1alpha1"
	gen_commonfate_control_directory_v1alpha1_user_pb_go "github.com/common-fate/sdk/gen/commonfate/control/directory/v1alpha1"
	gen_commonfate_control_feature_v1alpha1_feature_pb_go "github.com/common-fate/sdk/gen/commonfate/control/feature/v1alpha1"
	gen_commonfate_control_filters_v1alpha1_filters_pb_go "github.com/common-fate/sdk/gen/commonfate/control/filters/v1alpha1"
	gen_commonfate_control_insight_v1alpha1_insight_pb_go "github.com/common-fate/sdk/gen/commonfate/control/insight/v1alpha1"
	gen_commonfate_control_integration_reset_v1alpha1_reset_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/reset/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_auth0_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_aws_idc_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_aws_proxy_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_datastax_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_entra_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_gcp_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_healthcheck_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_integration_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_okta_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_opsgenie_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_pagerduty_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_proxy_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_proxy_resource_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_s3_log_destination_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_secret_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_slack_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_integration_v1alpha1_webhook_pb_go "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
	gen_commonfate_control_log_v1alpha1_authz_eval_pb_go "github.com/common-fate/sdk/gen/commonfate/control/log/v1alpha1"
	gen_commonfate_control_notification_v1alpha1_notification_pb_go "github.com/common-fate/sdk/gen/commonfate/control/notification/v1alpha1"
	gen_commonfate_control_oauth_v1alpha1_oauth_pb_go "github.com/common-fate/sdk/gen/commonfate/control/oauth/v1alpha1"
	gen_commonfate_control_resource_v1alpha1_resource_pb_go "github.com/common-fate/sdk/gen/commonfate/control/resource/v1alpha1"
	gen_commonfate_control_support_v1alpha1_support_pb_go "github.com/common-fate/sdk/gen/commonfate/control/support/v1alpha1"
	gen_commonfate_control_user_v1alpha1_user_pb_go "github.com/common-fate/sdk/gen/commonfate/control/user/v1alpha1"
	gen_commonfate_entity_v1alpha1_child_relation_pb_go "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	gen_commonfate_entity_v1alpha1_eid_pb_go "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	gen_commonfate_entity_v1alpha1_entity_pb_go "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	gen_commonfate_entity_v1alpha1_entity_service_pb_go "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	gen_commonfate_filters_v1alpha1_filters_pb_go "github.com/common-fate/sdk/gen/commonfate/filters/v1alpha1"
	gen_commonfate_leastprivilege_v1alpha1_entitlement_usage_pb_go "github.com/common-fate/sdk/gen/commonfate/leastprivilege/v1alpha1"
	gen_commonfate_leastprivilege_v1alpha1_leastprivilege_pb_go "github.com/common-fate/sdk/gen/commonfate/leastprivilege/v1alpha1"
	gen_commonfate_leastprivilege_v1alpha1_report_pb_go "github.com/common-fate/sdk/gen/commonfate/leastprivilege/v1alpha1"
)

var ProtoFiles = []protoreflect.FileDescriptor{
	gen_commonfate_access_provisioner_v1alpha1_provisioner_pb_go.File_commonfate_access_provisioner_v1alpha1_provisioner_proto,
	gen_commonfate_access_v1alpha1_access_pb_go.File_commonfate_access_v1alpha1_access_proto,
	gen_commonfate_access_v1alpha1_access_request_pb_go.File_commonfate_access_v1alpha1_access_request_proto,
	gen_commonfate_access_v1alpha1_audit_logs_pb_go.File_commonfate_access_v1alpha1_audit_logs_proto,
	gen_commonfate_access_v1alpha1_diagnostic_pb_go.File_commonfate_access_v1alpha1_diagnostic_proto,
	gen_commonfate_access_v1alpha1_extension_conditions_pb_go.File_commonfate_access_v1alpha1_extension_conditions_proto,
	gen_commonfate_access_v1alpha1_grant_pb_go.File_commonfate_access_v1alpha1_grant_proto,
	gen_commonfate_access_v1alpha1_grants_pb_go.File_commonfate_access_v1alpha1_grants_proto,
	gen_commonfate_access_v1alpha1_identity_pb_go.File_commonfate_access_v1alpha1_identity_proto,
	gen_commonfate_access_v1alpha1_integration_audit_logs_pb_go.File_commonfate_access_v1alpha1_integration_audit_logs_proto,
	gen_commonfate_access_v1alpha1_named_eid_pb_go.File_commonfate_access_v1alpha1_named_eid_proto,
	gen_commonfate_access_v1alpha1_proxy_session_pb_go.File_commonfate_access_v1alpha1_proxy_session_proto,
	gen_commonfate_access_v1alpha1_resource_pb_go.File_commonfate_access_v1alpha1_resource_proto,
	gen_commonfate_access_v1alpha1_user_pb_go.File_commonfate_access_v1alpha1_user_proto,
	gen_commonfate_authz_v1alpha1_action_pb_go.File_commonfate_authz_v1alpha1_action_proto,
	gen_commonfate_authz_v1alpha1_authz_pb_go.File_commonfate_authz_v1alpha1_authz_proto,
	gen_commonfate_authz_v1alpha1_evaluation_pb_go.File_commonfate_authz_v1alpha1_evaluation_proto,
	gen_commonfate_authz_v1alpha1_health_pb_go.File_commonfate_authz_v1alpha1_health_proto,
	gen_commonfate_authz_v1alpha1_policy_pb_go.File_commonfate_authz_v1alpha1_policy_proto,
	gen_commonfate_authz_v1alpha1_read_only_pb_go.File_commonfate_authz_v1alpha1_read_only_proto,
	gen_commonfate_authz_v1alpha1_request_pb_go.File_commonfate_authz_v1alpha1_request_proto,
	gen_commonfate_authz_v1alpha1_schema_pb_go.File_commonfate_authz_v1alpha1_schema_proto,
	gen_commonfate_authz_v1alpha1_validation_pb_go.File_commonfate_authz_v1alpha1_validation_proto,
	gen_commonfate_control_attest_v1alpha1_attest_pb_go.File_commonfate_control_attest_v1alpha1_attest_proto,
	gen_commonfate_control_attest_v1alpha1_attestation_pb_go.File_commonfate_control_attest_v1alpha1_attestation_proto,
	gen_commonfate_control_config_v1alpha1_access_workflow_pb_go.File_commonfate_control_config_v1alpha1_access_workflow_proto,
	gen_commonfate_control_config_v1alpha1_availability_spec_pb_go.File_commonfate_control_config_v1alpha1_availability_spec_proto,
	gen_commonfate_control_config_v1alpha1_aws_resource_scanner_pb_go.File_commonfate_control_config_v1alpha1_aws_resource_scanner_proto,
	gen_commonfate_control_config_v1alpha1_deployment_pb_go.File_commonfate_control_config_v1alpha1_deployment_proto,
	gen_commonfate_control_config_v1alpha1_gcp_role_group_pb_go.File_commonfate_control_config_v1alpha1_gcp_role_group_proto,
	gen_commonfate_control_config_v1alpha1_selector_pb_go.File_commonfate_control_config_v1alpha1_selector_proto,
	gen_commonfate_control_config_v1alpha1_slack_alert_pb_go.File_commonfate_control_config_v1alpha1_slack_alert_proto,
	gen_commonfate_control_config_v1alpha1_webhook_provisioner_pb_go.File_commonfate_control_config_v1alpha1_webhook_provisioner_proto,
	gen_commonfate_control_diagnostic_v1alpha1_diagnostic_pb_go.File_commonfate_control_diagnostic_v1alpha1_diagnostic_proto,
	gen_commonfate_control_directory_v1alpha1_directory_pb_go.File_commonfate_control_directory_v1alpha1_directory_proto,
	gen_commonfate_control_directory_v1alpha1_group_pb_go.File_commonfate_control_directory_v1alpha1_group_proto,
	gen_commonfate_control_directory_v1alpha1_integration_pb_go.File_commonfate_control_directory_v1alpha1_integration_proto,
	gen_commonfate_control_directory_v1alpha1_user_pb_go.File_commonfate_control_directory_v1alpha1_user_proto,
	gen_commonfate_control_feature_v1alpha1_feature_pb_go.File_commonfate_control_feature_v1alpha1_feature_proto,
	gen_commonfate_control_filters_v1alpha1_filters_pb_go.File_commonfate_control_filters_v1alpha1_filters_proto,
	gen_commonfate_control_insight_v1alpha1_insight_pb_go.File_commonfate_control_insight_v1alpha1_insight_proto,
	gen_commonfate_control_integration_reset_v1alpha1_reset_pb_go.File_commonfate_control_integration_reset_v1alpha1_reset_proto,
	gen_commonfate_control_integration_v1alpha1_auth0_pb_go.File_commonfate_control_integration_v1alpha1_auth0_proto,
	gen_commonfate_control_integration_v1alpha1_aws_idc_pb_go.File_commonfate_control_integration_v1alpha1_aws_idc_proto,
	gen_commonfate_control_integration_v1alpha1_aws_proxy_pb_go.File_commonfate_control_integration_v1alpha1_aws_proxy_proto,
	gen_commonfate_control_integration_v1alpha1_datastax_pb_go.File_commonfate_control_integration_v1alpha1_datastax_proto,
	gen_commonfate_control_integration_v1alpha1_entra_pb_go.File_commonfate_control_integration_v1alpha1_entra_proto,
	gen_commonfate_control_integration_v1alpha1_gcp_pb_go.File_commonfate_control_integration_v1alpha1_gcp_proto,
	gen_commonfate_control_integration_v1alpha1_healthcheck_pb_go.File_commonfate_control_integration_v1alpha1_healthcheck_proto,
	gen_commonfate_control_integration_v1alpha1_integration_pb_go.File_commonfate_control_integration_v1alpha1_integration_proto,
	gen_commonfate_control_integration_v1alpha1_okta_pb_go.File_commonfate_control_integration_v1alpha1_okta_proto,
	gen_commonfate_control_integration_v1alpha1_opsgenie_pb_go.File_commonfate_control_integration_v1alpha1_opsgenie_proto,
	gen_commonfate_control_integration_v1alpha1_pagerduty_pb_go.File_commonfate_control_integration_v1alpha1_pagerduty_proto,
	gen_commonfate_control_integration_v1alpha1_proxy_pb_go.File_commonfate_control_integration_v1alpha1_proxy_proto,
	gen_commonfate_control_integration_v1alpha1_proxy_resource_pb_go.File_commonfate_control_integration_v1alpha1_proxy_resource_proto,
	gen_commonfate_control_integration_v1alpha1_s3_log_destination_pb_go.File_commonfate_control_integration_v1alpha1_s3_log_destination_proto,
	gen_commonfate_control_integration_v1alpha1_secret_pb_go.File_commonfate_control_integration_v1alpha1_secret_proto,
	gen_commonfate_control_integration_v1alpha1_slack_pb_go.File_commonfate_control_integration_v1alpha1_slack_proto,
	gen_commonfate_control_integration_v1alpha1_webhook_pb_go.File_commonfate_control_integration_v1alpha1_webhook_proto,
	gen_commonfate_control_log_v1alpha1_authz_eval_pb_go.File_commonfate_control_log_v1alpha1_authz_eval_proto,
	gen_commonfate_control_notification_v1alpha1_notification_pb_go.File_commonfate_control_notification_v1alpha1_notification_proto,
	gen_commonfate_control_oauth_v1alpha1_oauth_pb_go.File_commonfate_control_oauth_v1alpha1_oauth_proto,
	gen_commonfate_control_resource_v1alpha1_resource_pb_go.File_commonfate_control_resource_v1alpha1_resource_proto,
	gen_commonfate_control_support_v1alpha1_support_pb_go.File_commonfate_control_support_v1alpha1_support_proto,
	gen_commonfate_control_user_v1alpha1_user_pb_go.File_commonfate_control_user_v1alpha1_user_proto,
	gen_commonfate_entity_v1alpha1_child_relation_pb_go.File_commonfate_entity_v1alpha1_child_relation_proto,
	gen_commonfate_entity_v1alpha1_eid_pb_go.File_commonfate_entity_v1alpha1_eid_proto,
	gen_commonfate_entity_v1alpha1_entity_pb_go.File_commonfate_entity_v1alpha1_entity_proto,
	gen_commonfate_entity_v1alpha1_entity_service_pb_go.File_commonfate_entity_v1alpha1_entity_service_proto,
	gen_commonfate_filters_v1alpha1_filters_pb_go.File_commonfate_filters_v1alpha1_filters_proto,
	gen_commonfate_leastprivilege_v1alpha1_entitlement_usage_pb_go.File_commonfate_leastprivilege_v1alpha1_entitlement_usage_proto,
	gen_commonfate_leastprivilege_v1alpha1_leastprivilege_pb_go.File_commonfate_leastprivilege_v1alpha1_leastprivilege_proto,
	gen_commonfate_leastprivilege_v1alpha1_report_pb_go.File_commonfate_leastprivilege_v1alpha1_report_proto,
}
