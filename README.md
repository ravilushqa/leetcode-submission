# LeetCode submissions

This repository syncs accepted LeetCode submissions using [`joshcai/leetcode-sync`](https://github.com/marketplace/actions/leetcode-sync).

Synced solutions are written to [`submissions/`](./submissions).

## Sync

The workflow runs weekly and can also be triggered manually from GitHub Actions:

`Actions` → `Sync LeetCode` → `Run workflow`.

Required repository secrets:

- `LEETCODE_CSRF_TOKEN`
- `LEETCODE_SESSION`
