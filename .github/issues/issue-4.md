---
code: null
id: I_kwDOHSRlrM5JrVQ0
databaseId: 1236096052
number: 4
url: https://github.com/octolab/config/issues/4
title: "input: Regexp type for kelseyhightower/envconfig integration"
issueType: null
labels:
  - "difficulty: easy"
  - "kind: feature"
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
createdAt: 2022-05-14T18:45:24Z
updatedAt: 2022-05-27T08:56:35Z
lastEditedAt: 2022-05-27T08:36:37Z
closedAt: null
---

# input: Regexp type for kelseyhightower/envconfig integration

**Motivation:** extend flag types.

**Draft**

```go
type Regexp regexp.Regexp

func (reg *Regexp) Decode(value string) error {
	raw, err := regexp.Compile(fmt.Sprintf(`^%s$`, value))
	if err != nil {
		return err
	}

	*reg = Regexp(*raw)
	return nil
}

func (reg *Regexp) Matcher() *regexp.Regexp {
	return (*regexp.Regexp)(reg)
}
```

See https://github.com/kelseyhightower/envconfig#custom-decoders.
