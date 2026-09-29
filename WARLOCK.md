<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# warlock-test-repo

Test repository and harness for the warlock CLI, with shell test suites covering symbol surfacing (check.sh), incremental refresh (incremental.sh), and scope/pact management (scopes.sh), alongside sample source directories across many languages that warlock operates on.

## Files

- `check.sh` (4.6 KB) — Bash test harness: runs warlock pact then asserts each directory's WARLOCK.md surfaces expected symbols (LEDGER_VERSION, Posting, HASH_SEED, NewRouter, Boot, CART_LIMIT, build_report, REPORT_VERSION, Shelf, Invoice, Stage) via present(), tallying pass/fail.
- `incremental.sh` (6.1 KB) — Bash test suite exercising `warlock refresh` on engine/core: checks changed/new/deleted files each update only their own WARLOCK.md line.
- `scopes.sh` (6.4 KB) — Bash test suite exercising scope add/remove/check/pact/unpact and sigil membership semantics against a sandboxed .warlock/pacts.toml manifest via the warlock CLI.

## Directories

- `data/` — Inventory JSON dataset, its SQL schema, and a binary logo asset; go here for inventory data or schema questions.
- `docs/` — Empty directory with no files or subdirectories.
- `engine/` — Holds core/, with ledger types (Entry, Side, Posting, Money), balance checks, and hashing; go here for data model or hashing questions.
- `infra/` — Terraform config provisioning the artifacts S3 bucket (aws_s3_bucket.artifacts, region variable, bucket_name output); go here for infra questions.
- `legacy/` — C frame codec (Frame, CODEC_VERSION, encode) and a stub C++ decoder class; go here for legacy codec questions.
- `monolith/` — Flat Go/Python/TypeScript domain modules with duplicate Apply/Resolve/Compact/Validate/Project/Reconcile/Emit/Settle workers, schema.sql, and clock.ts; go here for domain worker or store questions.
- `pipeline/` — Elixir Pipeline.Stage module (run/1, normalise) and its test suite; go here for stage-summing questions.
- `services/` — api/ (Server, Boot) and billing/ (invoice, ledger, payment, refund types); go here for service boot, routing, or billing questions.
- `tools/` — scripts/ with build/deploy scripts and inventory (Shelf, SHELF_LIMIT) and report (Report, build_report) utilities; go here for deploy, inventory, or report questions.
- `web/` — src/ with fetchCart client, legacy retry(fn) utility, and Cart component; go here for cart or retry questions.

## Structure

- Bash test harness: runs warlock pact then asserts each directory's WARLOCK.md surfaces expected symbols (LEDGER_VERSION, Posting, HASH_SEED, NewRouter, Boot, CART_LIMIT, build_report, REPORT_VERSION, Shelf, Invoice, Stage) via present(), tallying pass/fail.
- Bash test suite exercising `warlock refresh` on engine/core: checks changed/new/deleted files each update only their own WARLOCK.md line.
- Bash test suite exercising scope add/remove/check/pact/unpact and sigil membership semantics against a sandboxed .warlock/pacts.toml manifest via the warlock CLI.
