<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# scripts

Collection of build/deploy tooling and inventory and report utilities with accompanying tests, spanning shell, Ruby, and Python scripts.

## Files

- `deploy.sh` (99 B) — Shell script that runs the build, test, and ship stages in sequence for region eu-west-1, exiting on error.
- `inventory.rb` (256 B) — Inventory module with Shelf class (#initialize, #count summing items) and SHELF_LIMIT=40 constant; Inventory.build(items) constructs a Shelf. · declares `Inventory`, `Shelf`, `initialize`, `count`, `self.build`
- `inventory_spec.rb` (252 B) — RSpec spec for Inventory::Shelf, testing item counting and related shelf behavior across its examples.
- `report.py` (360 B) — Defines Report class (rows, total()) plus build_report/build_report_async factories; constants REPORT_VERSION=3, DEFAULT_ROWS=100. · declares `build_report_async`, `Report`, `__init__`, `total`, `build_report`
- `test_report.py` (224 B) — Test module with test_report_totals_its_rows, verifying report row totals are computed correctly.

## Structure

- Shell script that runs the build, test, and ship stages in sequence for region eu-west-1, exiting on error.
- Inventory module with Shelf class (#initialize, #count summing items) and SHELF_LIMIT=40 constant; Inventory.build(items) constructs a Shelf.
- RSpec spec for Inventory::Shelf, testing item counting and related shelf behavior across its examples.
- Defines Report class (rows, total()) plus build_report/build_report_async factories; constants REPORT_VERSION=3, DEFAULT_ROWS=100.
- Test module with test_report_totals_its_rows, verifying report row totals are computed correctly.
