<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# scripts

Utility scripts directory holding a deploy shell script and small example Ruby and Python modules with their tests, used for build/test/ship staging and inventory/report logic.

## Files

- `deploy.sh` (99 B) — Shell script setting REGION=eu-west-1 and looping over stages build, test, ship, echoing each stage as it runs.
- `inventory.rb` (256 B) — Defines Inventory module with SHELF_LIMIT=40, Shelf class (initialize, count) and self.build(items) factory method. · declares `Inventory`, `Shelf`, `initialize`, `count`, `self.build`
- `inventory_spec.rb` (252 B) — Spec tests for Inventory::Shelf, verifying it counts the items placed on it.
- `report.py` (360 B) — Reporting helpers: REPORT_VERSION=3, DEFAULT_ROWS=100, Report class with total(), build_report(rows), and async build_report_async(rows). · declares `build_report_async`, `Report`, `__init__`, `total`, `build_report`
- `test_report.py` (224 B) — Tests for report.py, including test_report_totals_its_rows verifying Report totals its rows.

## Structure

- inventory_spec.rb tests Inventory.build/Shelf defined in inventory.rb
- test_report.py tests build_report/Report defined in report.py
- deploy.sh runs independently of the Ruby and Python modules, sequencing build, test, ship stages

## Rules

- deploy.sh uses set -eu, so it exits on error or unset variable
- Shelf is bounded by SHELF_LIMIT = 40
- REPORT_VERSION is pinned to 3 and DEFAULT_ROWS to 100

## Where to look

- how deployment stages are run → `deploy.sh` `REGION`
- counting items in a shelf → `inventory.rb` `count`
- tests for shelf counting → `inventory_spec.rb` `Inventory::Shelf`
- building a report asynchronously → `report.py` `build_report_async`
- tests for report totals → `test_report.py` `test_report_totals_its_rows`
