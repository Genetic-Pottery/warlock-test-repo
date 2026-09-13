<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# engine

The engine directory is the ledger engine's top-level module, housing the core subsystem that defines the fundamental ledger data types, account balance utilities, and hashing primitives.

## Directories

- `core/` — Ledger core types and utilities: entries, postings, Money, hashing, and account/balance constants like VAULT_LIMIT and is_settled.

## Where to look

- how many accounts the vault can hold → `core` `VAULT_LIMIT`
- checking if the ledger is fully settled → `core` `is_settled`
- creating or posting a ledger entry → `core` `Entry`
- computing a hash of raw bytes → `core` `hash`
- what currency the ledger defaults to → `core` `DEFAULT_CURRENCY`
