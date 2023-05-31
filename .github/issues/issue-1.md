---
code: null
id: I_kwDOHSRlrM5JkCZR
databaseId: 1234183761
number: 1
url: https://github.com/octolab/config/issues/1
title: "input: improve config experience with Viper and Cobra"
issueType: null
labels:
  - "difficulty: medium"
  - "kind: improvement"
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
createdAt: 2022-05-12T15:50:38Z
updatedAt: 2022-05-27T09:08:38Z
lastEditedAt: 2022-05-27T09:08:38Z
closedAt: null
---

# input: improve config experience with Viper and Cobra

**Motivation:** there is no generic way to work with "prefixed" embedded structures. It leads to code duplication and makes it dirty.

**Scope:**

- improve Load functionality, make it generic
https://github.com/octomation/maintainer/blob/b2543194067c4e0b4daebd0b7def287762e5f4b7/internal/config/tool.go#L25-L50

- remove workaround with prefixes

- solve tech debt with keys warming
https://github.com/octomation/maintainer/blob/b2543194067c4e0b4daebd0b7def287762e5f4b7/internal/config/tool_test.go#L42-L50

- find a way how to make `squash with prefix` work
https://github.com/octomation/maintainer/blob/b2543194067c4e0b4daebd0b7def287762e5f4b7/internal/config/tool_test.go#L82-L120

<!-- 2022-05-27T09:06Z https://github.com/octolab/config/issues/1#issuecomment-1139428458
Related to mitchellh/mapstructure/issues/288.
-->
