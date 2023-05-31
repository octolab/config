---
code: null
id: I_kwDOHSRlrM5KiVuQ
databaseId: 1250515856
number: 9
url: https://github.com/octolab/config/issues/9
title: "feature: find way to build flexible nested feature set"
issueType: null
labels:
  - "difficulty: medium"
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
createdAt: 2022-05-27T08:55:45Z
updatedAt: 2022-05-27T08:55:45Z
lastEditedAt: null
closedAt: null
---

# feature: find way to build flexible nested feature set

**Motivation:** there is no way to work with features granularly, based on nested context, e.g. global > workspace > project > user.

**Inspiration:**

```php
final class Feature
{
    public function __construct(
        private string $name,
        private ?int $workspace = null,
        private ?int $user = null,
    ) {
    }

    public function keys(): array
    {
        $keys = [sprintf('feature:%s', $this->name)];
        if ($this->workspace) {
            $keys[] = sprintf('feature:%s:workspace:%d', $this->name, $this->workspace);
        }
        if ($this->user) {
            $keys[] = sprintf('feature:%s:user:%d', $this->name, $this->user);
        }
        if (count($keys) > 1) {
            // move up from specific to generic
            $keys = array_reverse($keys);
        }
        return $keys;
    }
}

final class Service
{
    public function isEnabled(Feature $feature, bool $default = false): bool
    {
        foreach ($feature->keys() as $key) {
            $toggle = Redis::get($key);
            if ($toggle === null || $toggle === false) {
                continue;
            }
            return filter_var($toggle, FILTER_VALIDATE_BOOL);
        }

        return $default;
    }

    public function toggle(Feature $feature, ?bool $value): void
    {
        $userFeatureKey = $feature->userKey();
        Redis::set($userFeatureKey, $value ?? !$this->isEnabled($feature));
    }
}
```

Need abstraction to adapt it to different storages: MySQL, Redis, PostgreSQL, Bolt, etc.

To do:

- [ ] Test abstraction on Bolt: https://github.com/boltdb/bolt.
