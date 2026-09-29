<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# core

Core module providing foundational ledger types, hashing, and balance utilities used across the engine.

## Files

- `balance.rs` (391 B) — Tiny helper: is_settled(open: usize) -> bool checks open == 0; VAULT_LIMIT const usize = 512. · declares `is_settled`, `VAULT_LIMIT`
- `hash.zig` (469 B) — Defines HASH_SEED (FNV-offset-like u64 1469598103934665603) and hash(bytes), a simple XOR-fold byte hash over a slice. · declares `HASH_SEED`, `hash`, `std`
- `ledger.rs` (1.8 KB) — Defines core ledger types Entry, Side, Money, Posting trait, post(), LEDGER_VERSION, DEFAULT_CURRENCY, with tests for posting and reversal. · declares `LEDGER_VERSION`, `Posting`, `Money`, `Entry`, `Side`, `new`, `reverse`, `post`, `amount`, `macro_rules`

## Structure

- Tiny helper: is_settled(open: usize) -> bool checks open == 0; VAULT_LIMIT const usize = 512.
- Defines HASH_SEED (FNV-offset-like u64 1469598103934665603) and hash(bytes), a simple XOR-fold byte hash over a slice.
- Defines core ledger types Entry, Side, Money, Posting trait, post(), LEDGER_VERSION, DEFAULT_CURRENCY, with tests for posting and reversal.
