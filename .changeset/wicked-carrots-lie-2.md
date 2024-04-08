---
"@common-fate/sdk": patch
---

The keychain now stores credentials with the key <Issuer URL/Client ID>, rather than the config file context key. This mitigates an issue where using the SDK without the file source would cause invalid credential issues, despite the user having valid credentials for a particular OIDC provider.
