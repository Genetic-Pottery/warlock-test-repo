<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# scripts

Deployment and inventory/reporting utility scripts, with accompanying specs and tests for the inventory and report logic.

## Files

- `deploy.sh` (99 B) — Shell script running build, test, and ship stages in sequence for the eu-west-1 region.
- `inventory.rb` (256 B) — Defines Inventory module with SHELF_LIMIT constant, Shelf class (count) tracking item totals, and self.build factory method. · declares `Inventory`, `Shelf`, `initialize`, `count`, `self.build`
- `inventory_spec.rb` (252 B) — RSpec spec for Inventory::Shelf, testing item counting behavior on a shelf.
- `report.py` (360 B) — Reporting helpers: Report class with total(), build_report()/build_report_async() constructors, REPORT_VERSION and DEFAULT_ROWS constants. · declares `build_report_async`, `Report`, `__init__`, `total`, `build_report`
- `test_report.py` (224 B) — Test module with test_report_totals_its_rows, verifying report row-totaling logic.

## Structure

- Shell script running build, test, and ship stages in sequence for the eu-west-1 region.
- Defines Inventory module with SHELF_LIMIT constant, Shelf class (count) tracking item totals, and self.build factory method.
- RSpec spec for Inventory::Shelf, testing item counting behavior on a shelf.
- Reporting helpers: Report class with total(), build_report()/build_report_async() constructors, REPORT_VERSION and DEFAULT_ROWS constants.
- Test module with test_report_totals_its_rows, verifying report row-totaling logic.
