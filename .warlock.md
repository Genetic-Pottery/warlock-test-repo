<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# warlock-test-repo

Root of the warlock-test-repo test suite: shell scripts driving the warlock CLI (pact, refresh, scope commands) against the sample directories to verify WARLOCK.md generation and updates.

## Files

- `check.sh` (4.6 KB) — Bash test harness: runs warlock pact then asserts each directory's WARLOCK.md surfaces expected symbols (LEDGER_VERSION, Posting, HASH_SEED, NewRouter, Boot, CART_LIMIT, build_report, REPORT_VERSION, Shelf, Invoice, Stage) via present(), tallying pass/fail.
- `incremental.sh` (6.1 KB) — Bash test suite exercising `warlock refresh` on engine/core: checks changed/new/deleted files each update only their own WARLOCK.md line.
- `scopes.sh` (6.4 KB) — Bash test suite exercising scope add/remove/check/pact/unpact and sigil membership semantics against a sandboxed .warlock/pacts.toml manifest via the warlock CLI.

## Directories

- `data/` — Inventory JSON dataset, its SQL schema, and a binary logo asset; go there for inventory data or schema questions.
- `docs/` — Empty directory with no files or subdirectories.
- `engine/` — Houses core submodule with ledger types, hashing, and balance helpers; go there for Entry, Side, Money, Posting questions.
- `infra/` — Terraform config provisioning the artifacts S3 bucket; go there for infrastructure/bucket questions.
- `legacy/` — C frame codec (Frame, encode(), CODEC_VERSION) plus a stub C++ decoder; go there for legacy codec questions.
- `monolith/` — Flat Go/Python/TypeScript domain modules with duplicated Apply/Resolve/Compact/etc. worker types and a shared schema.sql; go there for ledger, invoice, tenant, quota, and similar domain logic.
- `pipeline/` — Elixir Pipeline.Stage module summing items plus its test suite; go there for stage run/1 questions.
- `services/` — Groups the API server and billing domain types; go there for routing, server boot, or billing logic.
- `tools/` — Build/deploy tooling and inventory/report utilities with tests; go there for deploy, inventory, or report questions.
- `web/` — Web project root with a stubbed cart API client, retry utility, and Cart component; go there for cart or retry questions.

## Structure

- Bash test harness: runs warlock pact then asserts each directory's WARLOCK.md surfaces expected symbols (LEDGER_VERSION, Posting, HASH_SEED, NewRouter, Boot, CART_LIMIT, build_report, REPORT_VERSION, Shelf, Invoice, Stage) via present(), tallying pass/fail.
- Bash test suite exercising `warlock refresh` on engine/core: checks changed/new/deleted files each update only their own WARLOCK.md line.
- Bash test suite exercising scope add/remove/check/pact/unpact and sigil membership semantics against a sandboxed .warlock/pacts.toml manifest via the warlock CLI.
