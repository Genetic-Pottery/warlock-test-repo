<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# scripts

Collection of standalone operational and utility scripts covering deployment staging, inventory shelf modeling, and reporting, each paired with its own test file.

## Files

- `deploy.sh` (99 B) — POSIX shell script setting REGION=eu-west-1 and looping stages build, test, ship, echoing "running $stage" for each.
- `inventory.rb` (256 B) — Defines Inventory module with SHELF_LIMIT constant, Shelf class (count method) and self.build factory for creating shelves from items. · declares `Inventory`, `Shelf`, `initialize`, `count`, `self.build`
- `inventory_spec.rb` (252 B) — RSpec spec for Inventory::Shelf, testing item counting behavior.
- `report.py` (360 B) — Reporting helpers: REPORT_VERSION, DEFAULT_ROWS, Report class with total(), build_report(rows), build_report_async(rows). · declares `build_report_async`, `Report`, `__init__`, `total`, `build_report`
- `test_report.py` (224 B) — Test module verifying report row-totaling logic via test_report_totals_its_rows.

## Structure

- POSIX shell script setting REGION=eu-west-1 and looping stages build, test, ship, echoing "running $stage" for each.
- Defines Inventory module with SHELF_LIMIT constant, Shelf class (count method) and self.build factory for creating shelves from items.
- RSpec spec for Inventory::Shelf, testing item counting behavior.
- Reporting helpers: REPORT_VERSION, DEFAULT_ROWS, Report class with total(), build_report(rows), build_report_async(rows).
- Test module verifying report row-totaling logic via test_report_totals_its_rows.
