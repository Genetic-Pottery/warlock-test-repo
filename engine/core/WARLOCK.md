<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# core

Core ledger primitives: account balance limits/settlement checks, a byte-hashing utility, and the ledger entry/posting model with currency and version constants.

## Files

- `balance.rs` (391 B) — Defines VAULT_LIMIT (512) and is_settled(open), which despite its name returns whether open account count is zero, not a vault total. · declares `is_settled`, `VAULT_LIMIT`
- `hash.zig` (469 B) — FNV-1a-style hash: HASH_SEED constant and hash(bytes) computing a u64 by XOR-folding each byte into the seed. · declares `HASH_SEED`, `hash`, `std`
- `ledger.rs` (1.0 KB) — Ledger core: LEDGER_VERSION, DEFAULT_CURRENCY, Posting trait, Money type, Entry struct with new(), Side enum, and post() checking amount > 0. · declares `LEDGER_VERSION`, `Posting`, `Money`, `Entry`, `Side`, `new`, `post`, `amount`, `macro_rules`

## Structure

- Defines VAULT_LIMIT (512) and is_settled(open), which despite its name returns whether open account count is zero, not a vault total.
- FNV-1a-style hash: HASH_SEED constant and hash(bytes) computing a u64 by XOR-folding each byte into the seed.
- Ledger core: LEDGER_VERSION, DEFAULT_CURRENCY, Posting trait, Money type, Entry struct with new(), Side enum, and post() checking amount > 0.
