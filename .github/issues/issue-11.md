---
code: null
id: I_kwDOHSRlrM5LFL0X
databaseId: 1259650327
number: 11
url: https://github.com/octolab/config/issues/11
title: "input: File type for spf13/pflag integration"
issueType: null
labels:
  - "difficulty: easy"
  - "kind: feature"
  - "scope: code"
  - "scope: docs"
  - "scope: test"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2022-06-03T09:00:36Z
updatedAt: 2022-06-03T09:00:37Z
lastEditedAt: null
closedAt: null
---

# input: File type for spf13/pflag integration

**Motivation:** extend flag types.

**Interface**

```go
type Value interface {
	String() string
	Set(string) error
	Type() string
}
```

See [[issue-4]].
