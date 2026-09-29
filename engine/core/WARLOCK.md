<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# core

Core defines the ledger's fundamental data types, balance checks, and a byte hashing utility used across the engine.

## Files

- `balance.rs` (391 B) — Tiny helper: is_settled(open: usize) -> bool checks open == 0; VAULT_LIMIT const usize = 512. · declares `is_settled`, `VAULT_LIMIT`
- `hash.zig` (469 B) — Defines HASH_SEED (FNV-offset-like u64 1469598103934665603) and hash(bytes), a simple XOR-fold byte hash over a slice. · declares `HASH_SEED`, `hash`, `std`
- `ledger.rs` (978 B) — Core ledger types: Entry, Side enum, Posting trait, Money alias, post() checks amount>0; LEDGER_VERSION=7, DEFAULT_CURRENCY="GBP". · declares `LEDGER_VERSION`, `Posting`, `Money`, `Entry`, `Side`, `new`, `post`, `amount`, `macro_rules`

## Structure

- balance.rs — Tiny helper: is_settled(open: usize) -> bool checks open == 0; VAULT_LIMIT const usize = 512.
- hash.zig — Defines HASH_SEED (FNV-offset-like u64 1469598103934665603) and hash(bytes), a simple XOR-fold byte hash over a slice.
- ledger.rs — Core ledger types: Entry, Side enum, Posting trait, Money alias, post() checks amount>0; LEDGER_VERSION=7, DEFAULT_CURRENCY="GBP".
