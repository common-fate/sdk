---
"@common-fate/sdk": major
---

The api definitions have been updated to use a new authorization options system.
The option:

```
(commonfate.authz.v1alpha1.read_only) = true;
```

has been replaced with:

```
 option (commonfate.authz.v1alpha1.action_options) = {
      read_only: true
    };
```

Cedar Actions are now generated automatically based on the action_options configuration.
