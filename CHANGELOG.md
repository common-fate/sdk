# @common-fate/sdk

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
