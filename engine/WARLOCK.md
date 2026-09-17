<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# engine

Holds the ledger's core subsystem, defining account limits, settlement checks, a hash function, and the ledger's fundamental types, versioning, and posting validity rules.

## Directories

- `core/` — Ledger primitives (balances, hashing, core types/versioning/posting logic); go there for VAULT_LIMIT, is_settled, hash, LEDGER_VERSION, Posting, Entry, Side, or post().
