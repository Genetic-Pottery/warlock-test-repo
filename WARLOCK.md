<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# warlock-test-repo

Root of a test-fixture repository spanning data assets, financial engine primitives, infra config, a legacy codec, a monolith of near-duplicate domain wrappers, a pipeline stage, HTTP/billing services, deployment tools, and a web frontend, plus a check.sh fixture test asserting WARLOCK.md output correctness.

## Files

- `check.sh` (4.1 KB) — Bash fixture test asserting WARLOCK.md outputs via absent()/present() grep checks: lies excluded, symbols like LEDGER_VERSION/Posting/Stage surfaced, every directory documented.

## Directories

- `data/` — Inventory JSON dataset, a binary logo asset, and the SQL schema defining the inventory table and index.
- `engine/` — core/ subdirectory holds financial primitives: ledger entries, postings, settlement checks, and a hash utility.
- `infra/` — Terraform config provisioning a single AWS S3 bucket for artifacts.
- `legacy/` — Legacy C codec (Frame, CODEC_VERSION, encode()) plus a C++ Decoder, kept for backward compatibility.
- `monolith/` — Tenant-scoped near-duplicate Apply/Resolve/Compact/Validate/Project/Reconcile/Emit/Settle wrappers per domain entity across Go, Python, and TypeScript, backed by a shared schema.sql.
- `pipeline/` — Pipeline.Stage sums items via run/1, plus its ExUnit tests.
- `services/` — api/ HTTP server and routing; billing/ Invoice, Ledger, Payment, Refund domain types and tests.
- `tools/` — scripts/ holds deployment shell script plus Ruby inventory and Python report modules with specs and tests.
- `web/` — src/ holds BASE_URL/fetchCart, RETRY_LIMIT/retry, and the Cart component.

## Structure

- Bash fixture test asserting WARLOCK.md outputs via absent()/present() grep checks: lies excluded, symbols like LEDGER_VERSION/Posting/Stage surfaced, every directory documented.
