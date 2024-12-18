# @common-fate/sdk

## 1.71.0

### Minor Changes

- b0a703d: Add Jira ticket attachment support

### Patch Changes

- b0a703d: Add url and summary to justification attachment type

## 1.70.2

### Patch Changes

- bb9bb2c: add user details to session log details

## 1.70.1

### Patch Changes

- dac031a: add filters for query audit logs and query session logs

## 1.70.0

### Minor Changes

- c329c27: Add Snowflake integration.

## 1.69.1

### Patch Changes

- 3937b8e: add role and target to session detail

## 1.69.0

### Minor Changes

- 2742cb0: Added default_local_port to AWSRDS output

## 1.68.0

### Minor Changes

- 7ea0a0e: Added default_local_port option for AWS RDS Users

## 1.67.0

### Minor Changes

- f7c133d: Add Kubernetes protocol support to the Common Fate proxy.

## 1.66.0

### Minor Changes

- 34fe6a7: Adds new request review api for requesting an extra review from one or more principals

## 1.65.2

### Patch Changes

- 1aacec1: Remove requirement to specify JQL on the JIRA integration.

## 1.65.1

### Patch Changes

- 69dba8f: Rename ATTESTATION_TYPES to ATTACHMENT_TYPES.

## 1.65.0

### Minor Changes

- e4fa9a6: Add attachments to Justification to allow admins to query access requests matching certain JIRA tickets.
- d70e318: Add attachment service which provides contextual information which can be added to access requests

## 1.64.1

### Patch Changes

- 82914f8: add cloud_id and filter to jira integration

## 1.64.0

### Minor Changes

- fd56dc4: Add Integration type for Jira

### Patch Changes

- 629ab5a: added field for slack alerts to disable alerts for auto approved requests

## 1.63.0

### Minor Changes

- 831c704: Adds approval steps to access workflows.

## 1.62.0

### Minor Changes

- 7c399c1: add expiry options for closing requests onto workflow

### Patch Changes

- 82f9601: Update AWSProxy integration to store the AWS IDC integration as a reference.

## 1.61.0

### Minor Changes

- 945a9a9: Adds EKS access integration.

## 1.60.0

### Minor Changes

- 37f6906: Add reason pattern matching to validation.
- b49280d: Use standard filters for authz logs and my access requests.

### Patch Changes

- b8a17c9: add closing reason to close access request
- e5a8aa3: add optional rds endpoint to rds user

## 1.59.2

### Patch Changes

- aac7e5c: Access Request API now uses standardised filters.

## 1.59.1

### Patch Changes

- f5b6096: Added a field for breakglass reason to batchEnsure and activate requests and enforce that it is provided when breakglass is used.

## 1.59.0

### Minor Changes

- f65aa28: Added RPC for deleting a user.

## 1.58.0

### Minor Changes

- ddae8b0: Added a background task api for the integration service

## 1.57.1

### Patch Changes

- c0d32eb: Marks ListIntegrations RPC as read only.

## 1.57.0

### Minor Changes

- 10dce79: Adds List Selector API and Test Selector API

### Patch Changes

- 9a63227: Create QueryMyAccessRequests to find access requests where the current user is the requestor.

## 1.56.0

### Minor Changes

- b918ce7: Adds boolean conditions to the existing EntityFilter type for principal resource and actions.

  Adds new EntityTypeFilter with boolean condition

## 1.55.0

### Minor Changes

- 1f335f7: add apis for registering proxy and rds resource
- 15f432f: Adds API for managing secrets
- bc044db: Add filters for requested_by, closed_by, approved_by and request_status on QueryAccessRequest.

## 1.54.1

### Patch Changes

- 98e4102: Add access request id to ProvisionRequest.

## 1.54.0

### Minor Changes

- 9cc9766: Add configuration for slack user notifications, allowing opt in and opt out for specific DMs.

## 1.53.0

### Minor Changes

- f30b35a: Adds GetDeploymentSecret API and exposes additional deployment configuration
- 94bc67f: Add SDK methods to get and set Terraform Output for deployments

### Patch Changes

- 56551d7: Add extension_used, extendable_after to Extension.
- 4ef6933: adds healthcheck client for factory service

## 1.52.0

### Minor Changes

- 690de7d: Adds new deployment config service

## 1.51.1

### Patch Changes

- f48f8a4: Add opsgenie and datastax integration to directory

## 1.51.0

### Minor Changes

- 78b781e: Adds CLI Access Instructions to Grants.

## 1.50.0

### Minor Changes

- f6a98f1: Adds GetGRantOutput API which can be used to return the strongly typed grant output from provisioning if it exists for the grant.

## 1.49.0

### Minor Changes

- 5024170: Adds AWS Resource Scanner configuration service

## 1.48.0

### Minor Changes

- ef06a0d: Factory service: allow connect.ClientOptions to be passed to client constructor

### Patch Changes

- 6de9832: Update the OIDC provider library that we're using to the latest version to fix CVE-2024-28180

## 1.47.0

### Minor Changes

- 1816d6c: Add Department attribute to User
- 1816d6c: Directory: adds Okta integration type

## 1.46.0

### Minor Changes

- 743ebd1: Add update feature API

## 1.45.2

### Patch Changes

- a0cbe20: Package name and constructor fixes

## 1.45.1

### Patch Changes

- e2bac27: Detect a possible bad environment configuration when an invalid_scope error is returned when loading the access token.

## 1.45.0

### Minor Changes

- 6ad5170: Add force_close to forcibly close an Access Request, ignoring deprovisioning errors.

## 1.44.0

### Minor Changes

- 7da63c4: Add Directory API for user and group read access.

## 1.43.0

### Minor Changes

- b1e889d: Add UsageService to report deployment usage.
- c72f721: Add RDS integration resources
- c72f721: Adds GetGrant Method to the Grants service

## 1.42.2

### Patch Changes

- 2851c82: Fix Display() method that was removed

## 1.42.1

### Patch Changes

- 0bcb38a: Fix an issue where Display methods had been removed for accessv1alpha1 types.

## 1.42.0

### Minor Changes

- 704784c: Add support for extend access configuration with max extensions and extension duration in access workflows.
- 4ba41f4: Adds validation to rpc messages

### Patch Changes

- 82a26d0: add duration configuration to batch ensure request

## 1.41.1

### Patch Changes

- 1b9f2dc: Move the insight package to be inside the 'control' package

## 1.41.0

### Minor Changes

- 12d766f: Add Insights service

## 1.40.1

### Patch Changes

- 215cba4: Add support for suggesting a highest priority default role for availability specs

## 1.40.0

### Minor Changes

- f3854b7: Add notify_expiry_in_seconds to slack notification so that users can be notified at a preset time before their access expires.
- 5bfbd2f: Add support for breakglass access on BatchEnsure and Activate
- 7d93769: Add filtering for webhooks on specific actions

## 1.39.1

### Patch Changes

- 7b91b42: adds pagination fields onto entitlements tree api

## 1.39.0

### Minor Changes

- 4156cf4: adds api for querying entitlement tree

## 1.38.0

### Minor Changes

- 0a542c6: Add activate_allowed to list of AccessRequestActions.

## 1.37.0

### Minor Changes

- 6d4f1cd: Adds additional fields to the Integration resources types

## 1.36.1

### Patch Changes

- 61cc1f2: Support disabling all webhook handlers for Slack integration.

## 1.36.0

### Minor Changes

- 3857224: Add 'sso_access_portal_url' to allow the AWS SSO Start URL to be customised in the AWS IAM Identity Center integration.

## 1.35.1

### Patch Changes

- 1a1eca4: Adds additional methods to the policyset client

## 1.35.0

### Minor Changes

- 796f5f3: Add support for configuring deployment DNS nameservers for the default deployment domain
- 910dc31: Simplifies the ProvisionResponse type to return a single optional output entity

## 1.34.0

### Minor Changes

- ec19fbd: Add S3 Log Destination integration.

### Patch Changes

- ac6e044: Create GetAccessRequestActions API to return the list of actions that the current user is allowed to perform on the access request.

## 1.33.2

### Patch Changes

- 711849f: Adjust the suffix of the auth token on Windows.

## 1.33.0

### Minor Changes

- 6d98951: Add justification requirements configuration

### Patch Changes

- f0c0a65: Revert "add support for validation with protovalidate" - reverted for now because it is causing build issues.

## 1.32.1

### Patch Changes

- bb7b25b: Add support for validation with protovalidate

## 1.32.0

### Minor Changes

- ee2d9ee: Adds OpenTelemetry instrumentation to entity client writes.

### Patch Changes

- 2b39260: Add default duration to grant.

## 1.31.0

### Minor Changes

- 15a264e: Add default duration to Access Workflow.

## 1.30.0

### Minor Changes

- 83875f0: Adds Auth0 integration configuration.

## 1.29.0

### Minor Changes

- 513a673: Add GCP Role Group Configuration API

## 1.28.0

### Minor Changes

- f417cc7: add workflow id to request detail
- a41d9b8: add variable to slack notification resource to optionally send direct messages

## 1.27.0

### Minor Changes

- e96e5d3: Adds ResourceService and UserService which together replace the functionality of the GraphQL API which was implemented by the Authz service.
- 21585f0: Adds Managed Monitoring APIs.

## 1.26.2

### Patch Changes

- 8d985d9: Fix an issue where users were prompted continuously to enter a password when accessing the fallback file keychain in Linux.

## 1.26.1

### Patch Changes

- 464a458: Fix an issue where an entity marshalled with an eid.EID field could not be unmarshalled to a regular struct.

## 1.26.0

### Minor Changes

- 07124e7: Adds the Granted Profile Registry API specification.

## 1.25.0

### Minor Changes

- cc74bbb: Adds APIs for Common Fate in-app product support.
- e0fa93a: adds activation expiry to access workflows

### Patch Changes

- 4841873: Add new RPC for DebugEntitlementAccess

## 1.24.0

### Minor Changes

- df7756e: Add 'allowed' field to Evaluation to simplify audit log events
- df7756e: Add Webhook integration.

### Patch Changes

- df7756e: Fixes a breaking change made in entity ID JSON marshalling behaviour.

## 1.23.1

### Patch Changes

- 381d165: Updates the generated Go SDK to match the protobuf spec.

## 1.23.0

### Minor Changes

- c88b7f2: Adds 'target_path' attribute to entitlement.

## 1.22.1

### Patch Changes

- 4053841: Fix EID parsing for LinkedIdentity type

## 1.22.0

### Minor Changes

- 20a28f1: adds API endpoint for background job summaries.
- 0616f0b: Make tokenstore.Keyring public so that it can be used for listing tokens in the keychain (used for logging out of Common Fate)

## 1.21.0

### Minor Changes

- 025eb53: Add support for remote configuration of Common Fate SDK by providing a configuration URL.

### Patch Changes

- 4b404d6: Makes the EID parser configurable with the option to enforce the ID component to be quoted.
- 025eb53: The keychain now stores credentials with the key <Issuer URL/Client ID>, rather than the config file context key. This mitigates an issue where using the SDK without the file source would cause invalid credential issues, despite the user having valid credentials for a particular OIDC provider.
- 025eb53: Fix an issue where the config OAuth2.0 token was not updated after the login flow is completed.

## 1.20.0

### Minor Changes

- 3f34656: Add additional configuration to slack alerts to make the Approval action link to the web console to require SSO authentication.

## 1.19.1

### Patch Changes

- e7eae02: Fixes a regression in the Keychain when using file source for configuration, the context name was not set correctly when configuring the keychain access.

## 1.19.0

### Minor Changes

- 4becbd9: Adds support for specifying config sources via the CF_CONFIG_SOURCES environment variable.

## 1.18.0

### Minor Changes

- a8f9ffb: Allow configuring the SDK using environment variables like CF_API_URL, CF_OIDC_CLIENT_ID, and CF_OIDC_CLIENT_SECRET

## 1.17.0

### Minor Changes

- b668ea6: Rename PreviewEntitlements RPC to PreviewUserAccess

## 1.16.0

### Minor Changes

- fb6e665: Add new APIs for listing approvers and simulating access

## 1.15.0

### Minor Changes

- 4c10a20: Add RelinkEntraUsers RPC

### Patch Changes

- 06d4160: Adds Go clients for the ValidationService and SchemaService.
- 06d4160: Fixes the namespace for the SchemaService to be 'commonfate.authz.v1alpha1', instead of 'commonfate.control.authz.v1alpha1'.
- 0a2c0eb: Add opentelemetry spans to query methods
- 60fcc88: Add SDK method inputs to opentelemetry span

## 1.14.0

### Minor Changes

- f9d63c8: Add Cedar policy validation APIs

### Patch Changes

- 3d71a71: add simple entitlement type and api for listing target role pairs

## 1.13.1

### Patch Changes

- 48b6916: adds api for retrying river background tasks

## 1.13.0

### Minor Changes

- 0891480: Adds feature service. Clients can query the feature service to determine which Common Fate features are available to a particular user.
- 0891480: Moves the authorization evaluation query APIs to be under the 'commonfate.control' namespace, to better represent the fact that they are served by the Control Plane

## 1.12.3

### Patch Changes

- 0b1f5f8: adds duration field for adjusting duration for grant on batch ensure requests

## 1.12.2

### Patch Changes

- 67cb8c8: Add All method to entity client

## 1.12.1

### Patch Changes

- 76fe697: Add optional integration ID field to slack notification resource

## 1.12.0

### Minor Changes

- 7d1f434: Replaces the 'MetadataAttribute' with a simpler 'Tags' message used for authorization requests.
- 7d1f434: Adds endpoint to query for authorization evaluation metadata.

## 1.11.0

### Minor Changes

- 6ae9a07: Add CancelBackgroundJob and ListBackgroundJobs RPCs for managing background jobs.

## 1.10.0

### Minor Changes

- 65c8e59: Add evaluation duration on authorization Evaluation message.
- 32a6184: Adds a duration field to the Grant message.

## 1.9.0

### Minor Changes

- 19f18a5: Added DataStax integration

## 1.8.0

### Minor Changes

- 9a22117: Add provisioning attempt for grants

## 1.7.0

### Minor Changes

- 4dfc30d: Adds additional timing information to Grants and adds the Principal to Access Requests

## 1.6.1

### Patch Changes

- 2ae046d: Revert a change which incorrectly added the 'request_id' field to the BatchEnsure response

## 1.6.0

### Minor Changes

- 9fb6511: Add Okta integration type

## 1.5.0

### Minor Changes

- 68abc45: Revert the addition of provisioning status, add request ID to the batch ensure response

## 1.4.1

### Patch Changes

- f784061: Add a retry step to the loginflow which will detect errors validating the nonce and retry login 1 time. Retring usually fixes the issue.

## 1.4.0

### Minor Changes

- 68f719a: Adds additional fields to the AWS IDC integration resource

## 1.3.0

### Minor Changes

- f09745e: Add justifications to requests returned by the API

### Patch Changes

- 5bef771: Add order field to list entities rpc

## 1.2.0

### Minor Changes

- 0eda615: Add services for lease privilage analysis

## 1.1.0

### Minor Changes

- d22a3fe: add "provisioning" grant status
