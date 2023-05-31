---
code: null
id: I_kwDOHSRlrM5JrVXD
databaseId: 1236096451
number: 5
url: https://github.com/octolab/config/issues/5
title: "input: Regexp type for spf13/pflag integration"
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
createdAt: 2022-05-14T18:47:12Z
updatedAt: 2022-06-03T09:00:23Z
lastEditedAt: 2022-05-27T08:40:38Z
closedAt: null
---

# input: Regexp type for spf13/pflag integration

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
