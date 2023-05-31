---
code: null
id: I_kwDOHSRlrM5JrU_5
databaseId: 1236094969
number: 3
url: https://github.com/octolab/config/issues/3
title: "feature: decode from string"
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
createdAt: 2022-05-14T18:39:44Z
updatedAt: 2022-05-27T08:39:19Z
lastEditedAt: 2022-05-27T08:39:19Z
closedAt: null
---

# feature: decode from string

**Motivation:** there is no way to operate features (enable/disable them externally).

**Draft**

```go
header := "X-Features: feature-a=true, feature-b=false"

features.Unmarshal(header)
```

**PoC**

```go
vals := req.Header.Get(header.Features)

features.Apply(header)
if features.FindByName().Enabled {
  // do something
}
```

Related to octolab/protocol/issues/10.

**To do:**

- [ ] Find way to decode "multi-value" header. Is `;` the separator?
