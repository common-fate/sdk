# @common-fate/sdk

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
