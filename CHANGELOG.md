# @common-fate/sdk

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
