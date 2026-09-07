<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# tools

Utility tools directory containing a scripts subdirectory of deploy, inventory, and report examples used for build/test/ship staging and inventory/report logic.

## Directories

- `scripts/` — Deploy shell script plus Ruby inventory and Python report modules with tests; look here for staging, shelf counting, or report totals.

## Where to look

- how deployment stages are sequenced → `scripts` `deploy.sh`
- counting items on a shelf → `scripts` `Shelf`
- building a report or totaling rows → `scripts` `Report`
