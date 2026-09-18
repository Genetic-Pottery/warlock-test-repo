<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# warlock-test-repo

Root of the warlock-test-repo, holding bash test fixtures that exercise the warlock tool's check, incremental refresh, and scope/sigil behaviors against the rest of the repo's generated pacts and configs.

## Files

- `check.sh` (9.9 KB) — Bash fixture test that pacts the repo and greps generated WARLOCK.md files for attribution, undeclared names, symbol coverage, and completeness.
- `incremental.sh` (6.1 KB) — Bash script asserting warlock refresh reuses unchanged file lines (byte-identical) on engine/core when files are changed, added, or deleted.
- `scopes.sh` (6.4 KB) — scopes.sh: bash test asserting scope/sigil boundary behavior (status/--json fields) against the pacts.toml fixture via `warlock check/scope/pact/unpact/config`.

## Directories

- `data/` — Inventory database schema plus generated JSON data and a binary logo asset; go there for raw inventory records or the SQL table definition.
- `engine/` — Houses core/ with account balance/settlement logic, the FNV-1a hash utility, and the ledger Entry/Posting/Money model.
- `infra/` — Terraform config declaring the artifacts S3 bucket, its region variable, and the bucket_name output.
- `legacy/` — Legacy codec (CODEC_VERSION, Frame, encode) and Decoder class, superseded but retained for compatibility.
- `monolith/` — Tenant-scoped Go stores, Python event-fold projections/repos, and TypeScript gateways repeating Apply/Resolve/Compact/Validate/Project/Reconcile/Emit/Settle over a shared schema.sql.
- `pipeline/` — Elixir Pipeline.Stage summing items via run/1, with its ExUnit test.
- `services/` — Independent api/ (HTTP Server, ListenAddr) and billing/ (Invoice, Ledger, Payment, Refund) service implementations.
- `tools/` — scripts/ holding deploy.sh stages, Inventory::Shelf modeling, and Report totals, each with a paired test.
- `web/` — src/ holding a cart API stub, legacy retry helper, and Cart UI/state logic with tests.

## Structure

- Bash fixture test that pacts the repo and greps generated WARLOCK.md files for attribution, undeclared names, symbol coverage, and completeness.
- Bash script asserting warlock refresh reuses unchanged file lines (byte-identical) on engine/core when files are changed, added, or deleted.
- scopes.sh: bash test asserting scope/sigil boundary behavior (status/--json fields) against the pacts.toml fixture via warlock check/scope/pact/unpact/config.
