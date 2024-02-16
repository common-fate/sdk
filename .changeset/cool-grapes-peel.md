---
"@common-fate/sdk": patch
---

Add a retry step to the loginflow which will detect errors validating the nonce and retry login 1 time. Retring usually fixes the issue.
