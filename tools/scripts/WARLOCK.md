<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# scripts

Utility scripts directory holding a deployment shell script, a Ruby inventory module with its spec, and a Python reporting module with its test, unrelated small standalone examples.

## Files

- `deploy.sh` (99 B) — Shell script running deploy stages build, test, ship in sequence for REGION eu-west-1.
- `inventory.rb` (256 B) — Ruby module Inventory with SHELF_LIMIT and Shelf class holding items, counting totals; Inventory.build creates a Shelf. · declares `Inventory`, `Shelf`, `initialize`, `count`, `self.build`
- `inventory_spec.rb` (252 B) — Spec for Inventory::Shelf, testing that it counts the items on it.
- `report.py` (360 B) — Python module defining REPORT_VERSION, DEFAULT_ROWS, Report class with total(), and build_report/build_report_async factory functions. · declares `build_report_async`, `Report`, `__init__`, `total`, `build_report`
- `test_report.py` (224 B) — Test file verifying Report totals its rows via test_report_totals_its_rows.

## Structure

- inventory_spec.rb tests Inventory::Shelf defined in inventory.rb
- test_report.py tests Report and build_report from report.py
- deploy.sh runs independently of the Ruby and Python files

## Rules

- deploy.sh sets -eu to fail fast on errors and unset variables
- Inventory::Shelf enforces SHELF_LIMIT as a constant on the module
- report.py pins REPORT_VERSION to 3 and DEFAULT_ROWS to 100

## Where to look

- how deployment stages run → `deploy.sh` `REGION`
- counting items in a shelf → `inventory.rb` `Shelf`
- shelf item count test → `inventory_spec.rb` `Inventory::Shelf`
- building a report from rows → `report.py` `build_report`
- async report building → `report.py` `build_report_async`
- report totals test → `test_report.py` `test_report_totals_its_rows`
