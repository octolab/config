---
code: null
id: I_kwDOHSRlrM5KiTWV
databaseId: 1250506133
number: 8
url: https://github.com/octolab/config/issues/8
title: "input: Secret type support"
issueType: null
labels:
  - "difficulty: easy"
  - "kind: improvement"
  - "scope: code"
  - "scope: test"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2022-05-27T08:45:30Z
updatedAt: 2022-05-27T08:45:31Z
lastEditedAt: null
closedAt: null
---

# input: Secret type support

**Motivation:** there is no way to define this type of value.

Example

- https://github.com/octomation/maintainer/blob/af46d826f7c070347b26fa46efee6dbdee4f8be3/internal/command/github/root.go#L22
- https://github.com/octomation/maintainer/blob/af46d826f7c070347b26fa46efee6dbdee4f8be3/internal/config/tool.go#L18

**PoC**

```go
set := command.PersistentFlags()
set.Secret("token", "", "personal access token")
```
