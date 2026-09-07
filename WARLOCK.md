<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# warlock-test-repo

Root of warlock-test-repo, a fixture repository whose check.sh validates that a documentation-generating pass over the codebase surfaces real symbols and excludes planted lies, spanning engine, services, web, tools, pipeline, infra, legacy, and data subsystems.

## Files

- `check.sh` (4.0 KB) — Sanity-check script grepping generated WARLOCK.md files for expected symbols and planted lies, and verifying every directory has a document.

## Directories

- `data/` — Inventory dataset and SQL schema; open for inventory record shape, sku indexing, or table columns.
- `engine/` — Ledger engine core with balance settlement, hashing, and Entry/Posting types; open for accounting or hashing logic.
- `infra/` — Terraform config provisioning the AWS S3 artifacts bucket; open for infrastructure/region/bucket questions.
- `legacy/` — C/C++ codec module for frame encoding and decoding keyed to CODEC_VERSION; open for legacy frame codec questions.
- `pipeline/` — Elixir Pipeline.Stage module summing items via run/1; open for pipeline stage summation logic and its tests.
- `services/` — Backend services: api (HTTP routing/Boot) and billing (Invoice, Payment, Refund domain types); open for server startup or billing rules.
- `tools/` — Utility scripts for deploy staging, Ruby inventory Shelf counting, and Python report totals; open for build/test/ship or report questions.
- `web/` — Web app source with fetchCart API helper, retry utility, and Cart component/CART_LIMIT; open for shopping cart UI logic.

## Structure

- check.sh invokes the warlock binary to pact (generate) WARLOCK.md files across every subdirectory before checking them
- check.sh greps each subdirectory's WARLOCK.md for required symbols and forbidden planted lies
- check.sh iterates a fixed list of directories to confirm each has its own WARLOCK.md

## Rules

- check.sh assertions are greps, never exact text comparisons
- a passing run rewrites its whole document in fresh wording every time, so no golden-file comparison is possible
- check.sh exits with status 1 if any check fails, printing pass/fail counts

## Where to look

- how to run the fixture's validation checks → `check.sh` `pact`
- what strings must never appear in generated docs → `check.sh` `absent`
- inventory data shape or sku field → `data` `sku`
- ledger balance or hashing logic → `engine` `LEDGER_VERSION`
- AWS resource provisioning → `infra` `aws_s3_bucket`
- legacy frame codec encode/decode → `legacy` `CODEC_VERSION`
- pipeline stage item summation → `pipeline` `Stage`
- API server routing or billing invoice logic → `services` `Invoice`
- shelf counting or report totals script → `tools` `Shelf`
- cart fetch, retry, or cart limit UI logic → `web` `CART_LIMIT`
