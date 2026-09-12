<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# engine

The engine directory is the top-level container for the engine's ledger and accounting subsystem, holding the core module that implements entries, postings, balance checks, and hashing.

## Directories

- `core/` — Core ledger primitives: entries, postings, balance checks, and hashing; go there for accounting logic, versioning, or hash utility questions.

## Where to look

- how ledger entries are created and posted → `core` `Entry`
- checking if all accounts are settled → `core` `is_settled`
- hashing byte data → `core` `hash`
