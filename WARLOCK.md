<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# warlock-test-repo

Root of the warlock-test-repo fixture repository, holding shell test scripts that exercise warlock's pact/unpact, refresh, and scope/sigil boundary behavior against generated WARLOCK.md files, plus the source subdirectories those scripts target.

## Files

- `check.sh` (4.8 KB) — Bash fixture test that pacts the repo and greps generated WARLOCK.md files for absent lies and present symbols like LEDGER_VERSION, Posting, NewRouter.
- `incremental.sh` (6.1 KB) — Shell script testing warlock refresh's incremental reuse against a real model: changed/new/deleted files under engine/core keep sibling lines byte-identical.
- `scopes.sh` (6.4 KB) — scopes.sh: shell fixture asserting the scope/sigil boundary against warlock's check/scope/pact/unpact exit statuses and --json fields, with HOME sandboxed.

## Directories

- `data/` — Inventory JSON dataset, binary logo asset, and inventory table SQL schema; go here for inventory schema or asset questions.
- `engine/` — Ledger core subsystem (via core/) with VAULT_LIMIT, is_settled, hash, LEDGER_VERSION, Posting, Entry, Side, post(); go here for ledger primitive questions.
- `infra/` — Terraform config for a single AWS S3 artifacts bucket; go here for infrastructure provisioning questions.
- `legacy/` — Legacy C codec (Frame, CODEC_VERSION, encode()) plus a C++ Decoder for reading frames back; go here for codec/decoder compatibility questions.
- `monolith/` — Tenant-scoped near-duplicate Apply/Resolve/Compact/Validate/Project/Reconcile/Emit/Settle wrappers per domain entity across Go, Python, and TypeScript, backed by a shared SQL schema; go here for domain-entity CRUD-wrapper questions.
- `pipeline/` — Pipeline.Stage Elixir module summing items via run/1, plus its ExUnit tests; go here for pipeline stage logic questions.
- `services/` — HTTP API server/routing (api/) and billing domain types across Java, Kotlin, C#, Swift (billing/); go here for server, routing, or billing logic questions.
- `tools/` — Deployment and inventory/reporting scripts under scripts/; go here for build/deploy or inventory/report logic questions.
- `web/` — Web src/ holding BASE_URL/fetchCart, RETRY_LIMIT/retry, and the Cart component; go here for API, retry, or cart questions.

## Structure

- Bash fixture test that pacts the repo and greps generated WARLOCK.md files for absent lies and present symbols like LEDGER_VERSION, Posting, NewRouter.
- Shell script testing warlock refresh's incremental reuse against a real model: changed/new/deleted files under engine/core keep sibling lines byte-identical.
- scopes.sh: shell fixture asserting the scope/sigil boundary against warlock's check/scope/pact/unpact exit statuses and --json fields, with HOME sandboxed.
