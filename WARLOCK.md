<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# warlock-test-repo

Root of warlock-test-repo, a test fixture repository for the warlock tool, containing shell test suites (check.sh, incremental.sh, scopes.sh) plus subdirectories of sample source across many languages used to validate WARLOCK.md generation, refresh, and scope commands.

## Files

- `check.sh` (4.6 KB) — Bash test harness: runs warlock pact then asserts each directory's WARLOCK.md surfaces expected symbols (LEDGER_VERSION, Posting, HASH_SEED, NewRouter, Boot, CART_LIMIT, build_report, REPORT_VERSION, Shelf, Invoice, Stage) via present(), tallying pass/fail.
- `incremental.sh` (6.1 KB) — Bash test suite exercising `warlock refresh` on engine/core: checks changed/new/deleted files each update only their own WARLOCK.md line.
- `scopes.sh` (6.4 KB) — Bash test suite exercising scope add/remove/check/pact/unpact and sigil membership semantics against a sandboxed .warlock/pacts.toml manifest via the warlock CLI.

## Directories

- `data/` — Inventory JSON dataset, its SQL schema, and an unrelated logo.bin asset.
- `docs/` — Empty directory with no files or subdirectories.
- `engine/` — Holds core/ with ledger types (Entry, Side, Posting, Money), balance checks, and hashing.
- `infra/` — Terraform config provisioning the artifacts S3 bucket (aws_s3_bucket.artifacts).
- `legacy/` — C frame codec (Frame, CODEC_VERSION, encode) plus a stub C++ decoder class.
- `monolith/` — Flat Go/Python/TypeScript domain modules with duplicate Apply/Resolve/Compact/Validate/Project/Reconcile/Emit/Settle workers, a shared schema.sql, and clock.ts.
- `pipeline/` — Elixir Pipeline.Stage module (run/1 via Enum.reduce) with its test suite.
- `services/` — Groups api/ (server boot on 0.0.0.0:8080) and billing/ (Invoice, Ledger, Payment, Refund across Java, Kotlin, C#, Swift).
- `tools/` — Holds scripts/ with build/deploy, inventory (Shelf, SHELF_LIMIT), and report (Report, build_report) utilities plus tests.
- `web/` — Holds src/ with fetchCart client, legacy retry(fn) utility, and the Cart component.

## Structure

- Bash test harness: runs warlock pact then asserts each directory's WARLOCK.md surfaces expected symbols (LEDGER_VERSION, Posting, HASH_SEED, NewRouter, Boot, CART_LIMIT, build_report, REPORT_VERSION, Shelf, Invoice, Stage) via present(), tallying pass/fail.
- Bash test suite exercising `warlock refresh` on engine/core: checks changed/new/deleted files each update only their own WARLOCK.md line.
- Bash test suite exercising scope add/remove/check/pact/unpact and sigil membership semantics against a sandboxed .warlock/pacts.toml manifest via the warlock CLI.
