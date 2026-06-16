# cofiswarm-infer-sglang

Cofiswarm component: `infer-sglang`.

- Layout: [REPO-STANDARD-LAYOUT](https://github.com/keepdevops/cofiswarmdev/blob/main/docs/REPO-STANDARD-LAYOUT.md)
- Migration: [MIGRATION-SPRINTS](https://github.com/keepdevops/cofiswarmdev/blob/main/docs/MIGRATION-SPRINTS.md)

## FHS paths

| Path | Purpose |
|------|---------|
| `/etc/cofiswarm/infer-sglang/` | config |
| `/var/lib/cofiswarm/infer-sglang/` | state |
| `/var/log/cofiswarm/infer-sglang/` | logs |

## Test

```bash
./test/scripts/assert-layout.sh infer-sglang
```
