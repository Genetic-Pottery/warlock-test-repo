<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# tools

The tools directory holds the scripts subdirectory of utility scripts used for build/deploy automation and small Inventory and Report modules with their tests.

## Directories

- `scripts/` — Deploy shell script plus Ruby Inventory and Python Report modules with tests; go here for build/deploy stages, shelf counting, or report totals.

## Structure

- The only content of this directory is the scripts subdirectory.

## Where to look

- how deployment stages run → `scripts` `deploy.sh`
- shelf item limit or counting logic → `scripts` `SHELF_LIMIT`
- building a report asynchronously → `scripts` `build_report_async`
