<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# warlock-test-repo

Root of the warlock-test-repo test fixture: shell test harnesses (check.sh, incremental.sh, scopes.sh) exercising the warlock CLI's pact, refresh, and scope commands, alongside subdirectories of sample code spanning many languages for those tests to target.

## Files

- `check.sh` (4.6 KB) — Bash test harness: runs warlock pact then asserts each directory's WARLOCK.md surfaces expected symbols (LEDGER_VERSION, Posting, HASH_SEED, NewRouter, Boot, CART_LIMIT, build_report, REPORT_VERSION, Shelf, Invoice, Stage) via present(), tallying pass/fail.
- `incremental.sh` (6.1 KB) — Bash test suite exercising `warlock refresh` on engine/core: checks changed/new/deleted files each update only their own WARLOCK.md line.
- `scopes.sh` (6.4 KB) — Bash test suite exercising scope add/remove/check/pact/unpact and sigil membership semantics against a sandboxed .warlock/pacts.toml manifest via the warlock CLI.

## Directories

- `data/` — Inventory JSON dataset, its SQL schema, and a binary logo asset; go here for inventory data or schema questions.
- `docs/` — Empty directory with no files or subdirectories.
- `engine/` — Holds core/ with ledger types (Entry, Side, Posting, Money) and hashing; go here for data model or hashing questions.
- `infra/` — Terraform config provisioning the artifacts S3 bucket; go here for infrastructure or bucket questions.
- `legacy/` — C frame codec (Frame, CODEC_VERSION, encode) and a stub C++ decoder; go here for legacy codec questions.
- `monolith/` — Flat collection of duplicated Go/Python/TypeScript domain modules (ledger, invoice, tenant, etc.) with repeated Apply/Resolve/Compact worker types.
- `pipeline/` — Elixir Pipeline.Stage module summing items, plus its test suite; go here for stage/run questions.
- `services/` — Holds api/ (Server, Boot) and billing/ (Invoice, Ledger, Payment, Refund) subdirectories; go here for service entry or billing questions.
- `tools/` — Holds scripts/ with build/deploy, inventory (Shelf), and report (build_report) utilities and specs.
- `web/` — Holds src/ with fetchCart, retry(fn), and the Cart component; go here for cart or retry questions.

## Structure

- Bash test harness: runs warlock pact then asserts each directory's WARLOCK.md surfaces expected symbols (LEDGER_VERSION, Posting, HASH_SEED, NewRouter, Boot, CART_LIMIT, build_report, REPORT_VERSION, Shelf, Invoice, Stage) via present(), tallying pass/fail.
- Bash test suite exercising `warlock refresh` on engine/core: checks changed/new/deleted files each update only their own WARLOCK.md line.
- Bash test suite exercising scope add/remove/check/pact/unpact and sigil membership semantics against a sandboxed .warlock/pacts.toml manifest via the warlock CLI.
