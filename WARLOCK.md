<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# warlock-test-repo

Root of warlock-test-repo, a fixture repository exercising the warlock documentation pass across multiple languages and subsystems: a ledger engine, a billing/API monolith and services split, a data-backed inventory system, legacy codec, pipeline stage, infra, web app, and tooling.

## Files

- `check.sh` (4.1 KB) — Bash test harness that runs warlock pact over the fixture and greps generated WARLOCK.md files for required symbols and absent planted lies.

## Directories

- `data/` — Inventory dataset and schema.sql defining the inventory table; go here for record fields or sku indexing questions.
- `engine/` — Ledger engine top-level module wrapping the core subsystem's entry, posting, Money, and hashing types; go here for ledger data type or hashing questions.
- `infra/` — Terraform config provisioning the S3 artifacts bucket; go here for deployment region or bucket_name questions.
- `legacy/` — Old C/C++ codec subsystem with Frame, encode, CODEC_VERSION, and Decoder; go here for legacy frame encode/decode questions.
- `monolith/` — Go/TypeScript/Python monolith implementing billing, subscription, fulfillment, and platform-ops domains over one Postgres schema; go here for ledger, invoice, quota, webhook, or reconciliation logic.
- `pipeline/` — Elixir Pipeline.Stage module reducing items into a summed value; go here for pipeline stage summation questions.
- `services/` — Api and billing service subsystems: HTTP server boot/routing and invoice/ledger/payment/refund domain logic.
- `tools/` — Scripts subdirectory of build/deploy automation plus Inventory and Report utility modules; go here for deploy stages or report totals.
- `web/` — Web app top-level with api.ts client, legacy.js retry helper, and the src/components cart subsystem.

## Structure

- check.sh runs a warlock pact over this repository then greps the generated WARLOCK.md files under engine, services, tools, pipeline, legacy, data, and root.
- check.sh checks that every directory it lists, including engine, engine/core, services, services/api, services/api/handlers, services/billing, web, web/src, web/src/components, tools, tools/scripts, pipeline, infra, legacy, and data, has its own WARLOCK.md.
- engine depends on its core subdirectory for the ledger data types and constants it exposes.
- services groups the api and billing subsystems as independent subdirectories.
- tools contains only the scripts subdirectory.
- web depends on its src subdirectory for the api client, legacy retry helper, and cart components.

## Rules

- check.sh treats every assertion as a grep against generated text, never an exact-text comparison against a golden file.
- check.sh deletes existing WARLOCK.md files and restores pipeline/WARLOCK.md from git before pacting, unless run with --no-pact.

## Where to look

- how the fixture's generated documents are validated → `check.sh` `pact`
- ledger data types, Money, or hashing constants → `engine` `VAULT_LIMIT`
- billing invoices, ledgers, payments, or refunds → `services` `Invoice`
- inventory record fields or sku indexing → `data` `inventory_sku_idx`
- legacy codec frame encoding or decoding → `legacy` `CODEC_VERSION`
- monolith billing/subscription/fulfillment domain logic → `monolith` `LedgerApplyer`
- pipeline stage item summation → `pipeline` `Pipeline.Stage`
- build/deploy scripts or report totals → `tools` `build_report_async`
- shopping cart component or api client questions → `web` `fetchCart`
- S3 artifacts bucket or deploy region → `infra` `bucket_name`
