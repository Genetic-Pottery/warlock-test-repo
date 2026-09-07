<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# engine

The engine directory holds the ledger engine, containing the core subsystem that implements balance accounting, hashing, and ledger posting logic.

## Directories

- `core/` — Ledger core: balance settlement checks, FNV-style hashing, and Entry/Posting types; open for accounting or hashing questions.

## Where to look

- how are accounts marked settled → `core` `is_settled`
- ledger schema/version number → `core` `LEDGER_VERSION`
- hashing byte data → `core` `hash`
