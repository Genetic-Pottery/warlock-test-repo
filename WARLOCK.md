<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# warlock-test-repo

Root of the warlock-test-repo, holding shell-based test harnesses that exercise the warlock CLI's pact, refresh, and scope commands against the repo's own generated WARLOCK.md files.

## Files

- `check.sh` (4.6 KB) — Bash test harness: runs warlock pact then asserts each directory's WARLOCK.md surfaces expected symbols (LEDGER_VERSION, Posting, HASH_SEED, NewRouter, Boot, CART_LIMIT, build_report, REPORT_VERSION, Shelf, Invoice, Stage) via present(), tallying pass/fail.
- `incremental.sh` (6.1 KB) — Bash test suite exercising `warlock refresh` on engine/core: checks changed/new/deleted files each update only their own WARLOCK.md line.
- `scopes.sh` (6.4 KB) — Bash test suite exercising scope add/remove/check/pact/unpact and sigil membership semantics against a sandboxed .warlock/pacts.toml manifest via the warlock CLI.

## Directories

- `data/` — Inventory JSON dataset, its SQL schema, and a binary logo asset; go here for inventory data or schema questions.
- `docs/` — Empty directory with no files or subdirectories.
- `engine/` — Holds core/ with ledger types (Entry, Side, Posting, Money), balance checks, and hashing; go here for ledger data model questions.
- `infra/` — Terraform config provisioning the artifacts S3 bucket; go here for infrastructure or bucket questions.
- `legacy/` — C frame codec (Frame, CODEC_VERSION, encode) and a stub C++ decoder; go here for legacy codec questions.
- `monolith/` — Flat Go/Python/TypeScript domain modules with duplicated Apply/Resolve/Compact/Validate/Project/Reconcile/Emit/Settle workers plus schema.sql and clock.ts; go here for domain worker or store duplication questions.
- `pipeline/` — Elixir Pipeline.Stage module summing items via run/1, with its test suite; go here for stage/summing logic questions.
- `services/` — Groups api/ (Server, ListenAddr, Boot) and billing/ (Invoice, Ledger, Payment, Refund); go here for service startup or billing questions.
- `tools/` — Holds scripts/ with build/deploy scripts and Shelf/Report inventory and reporting utilities; go here for deploy, inventory, or report questions.
- `web/` — Holds src/ with fetchCart API client, legacy retry(fn) utility, and the Cart component; go here for cart or retry questions.

## Structure

- Bash test harness: runs warlock pact then asserts each directory's WARLOCK.md surfaces expected symbols (LEDGER_VERSION, Posting, HASH_SEED, NewRouter, Boot, CART_LIMIT, build_report, REPORT_VERSION, Shelf, Invoice, Stage) via present(), tallying pass/fail.
- Bash test suite exercising `warlock refresh` on engine/core: checks changed/new/deleted files each update only their own WARLOCK.md line.
- Bash test suite exercising scope add/remove/check/pact/unpact and sigil membership semantics against a sandboxed .warlock/pacts.toml manifest via the warlock CLI.
