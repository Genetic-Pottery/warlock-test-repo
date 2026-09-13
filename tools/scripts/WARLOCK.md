<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# scripts

Utility scripts directory holding a deploy shell script and small Ruby and Python modules (Inventory, Report) with their accompanying tests, used for build/deploy and inventory or reporting logic.

## Files

- `deploy.sh` (99 B) — Shell script that sets REGION to eu-west-1 and loops through build, test, ship stages, echoing the running stage.
- `inventory.rb` (256 B) — Ruby module Inventory with SHELF_LIMIT constant, Shelf class holding items and a count method, and self.build factory method. · declares `Inventory`, `Shelf`, `initialize`, `count`, `self.build`
- `inventory_spec.rb` (252 B) — Spec file testing Inventory::Shelf, including a test that counts the items on it.
- `report.py` (360 B) — Python module defining REPORT_VERSION and DEFAULT_ROWS constants, Report class with a total method, and build_report and build_report_async functions. · declares `build_report_async`, `Report`, `__init__`, `total`, `build_report`
- `test_report.py` (224 B) — Test file with test_report_totals_its_rows verifying Report totals behavior.

## Structure

- inventory_spec.rb tests the Shelf class and count method defined in inventory.rb.
- test_report.py tests the Report class and build_report function defined in report.py.
- deploy.sh runs its stages independently of the Ruby and Python files in this directory.

## Rules

- deploy.sh sets REGION to eu-west-1 for its run.
- deploy.sh uses set -eu to fail on errors and unset variables.
- Inventory module defines SHELF_LIMIT as 40.
- report.py pins REPORT_VERSION to 3 and DEFAULT_ROWS to 100.

## Where to look

- how deployment stages run → `deploy.sh` `REGION`
- shelf item limit or counting logic → `inventory.rb` `SHELF_LIMIT`
- tests for shelf counting behavior → `inventory_spec.rb` `Shelf`
- building a report asynchronously → `report.py` `build_report_async`
- tests verifying report totals → `test_report.py` `test_report_totals_its_rows`
