<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# warlock-test-repo

Root of the warlock-test-repo fixture, a multi-language sample codebase covering an accounting engine, API/billing services, a web cart app, a data schema, an elixir pipeline stage, legacy C/C++ codec, tools scripts, and terraform infra, used to validate warlock's documentation pass.

## Files

- `check.sh` (4.1 KB) — Sanity-check script that pacts the fixture with warlock, then greps generated WARLOCK.md files for planted lies (absent) and required symbols (present), e.g. LEDGER_VERSION, Posting, HASH_SEED, NewRouter, Boot, CART_LIMIT, build_report, REPORT_VERSION, Shelf, Invoice, Stage.

## Directories

- `data/` — Inventory dataset and SQL schema; go there for inventory table columns, sku indexing, or the logo asset.
- `engine/` — Ledger and accounting subsystem; go there for entry/posting creation, balance settlement, or hashing.
- `infra/` — Terraform config provisioning the artifacts S3 bucket; go there for AWS region or bucket name questions.
- `legacy/` — Legacy C/C++ codec subsystem; go there for frame encoding, codec version, or Decoder logic.
- `pipeline/` — Elixir Pipeline.Stage module and its tests; go there for how items are summed and tested.
- `services/` — API entry-point and billing domain services; go there for server boot/listen or invoice/payment/refund logic.
- `tools/` — Utility scripts subsystem with deploy, Ruby Shelf, and Python report modules; go there for deployment or report-building logic.
- `web/` — Web app with shared API utilities, retry helper, and Cart component subsystem; go there for cart or fetch/retry logic.

## Structure

- check.sh invokes the warlock binary to pact this repo, then verifies the resulting WARLOCK.md files across every directory listed.

## Rules

- check.sh treats a pass as one that rewrites its whole document in fresh wording every run, so checks compare via grep, never against a golden file.
- WARLOCK env var overrides the warlock binary path, defaulting to ../warlock/target/release/warlock.

## Where to look

- how to run the fixture's self-check → `check.sh` `ok`
- inventory table columns or sku index → `data` `schema.sql`
- ledger entries, postings, or balance checks → `engine` `Entry`
- AWS region or artifacts bucket name → `infra` `aws_s3_bucket`
- legacy frame encoding or decoding → `legacy` `Decoder`
- how pipeline items are summed → `pipeline` `run`
- API server boot or billing invoice totals → `services` `Boot`
- deployment stages or report building scripts → `tools` `build_report`
- cart line items or CART_LIMIT → `web` `Cart`
