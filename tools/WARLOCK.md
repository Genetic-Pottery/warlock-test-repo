<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# tools

Utility scripts directory of the tools subsystem, holding a deployment shell script and a nested scripts directory with example Ruby and Python inventory/report modules and their tests.

## Directories

- `scripts/` — Holds deploy.sh, a Ruby Inventory/Shelf module with spec, and a Python Report module with test; go there for deployment stages, shelf item counting, or report building.

## Where to look

- how deployment stages run → `scripts` `REGION`
- counting items in a shelf → `scripts` `Shelf`
- building a report from rows → `scripts` `build_report`
