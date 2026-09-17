<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# core

Core ledger primitives: account balance limits and settlement checks, a lightweight FNV-1-style hash function, and the ledger's core types, versioning, and posting validity logic.

## Files

- `balance.rs` (391 B) — Defines VAULT_LIMIT (512) account cap and is_settled(open) checking whether open account count is zero; doc comments reference Decoder::decode() and LEDGER_VERSION defined elsewhere. · declares `is_settled`, `VAULT_LIMIT`
- `hash.zig` (469 B) — Defines HASH_SEED constant and hash(bytes) FNV-1-style XOR hashing function returning a u64. · declares `HASH_SEED`, `hash`, `std`
- `ledger.rs` (1.0 KB) — Core ledger: LEDGER_VERSION, DEFAULT_CURRENCY, Posting trait, Money alias, Entry (id, amount, ::new), Side enum, post() validity check. · declares `LEDGER_VERSION`, `Posting`, `Money`, `Entry`, `Side`, `new`, `post`, `amount`, `macro_rules`

## Structure

- Defines VAULT_LIMIT (512) account cap and is_settled(open) checking whether open account count is zero; doc comments reference Decoder::decode() and LEDGER_VERSION defined elsewhere.
- Defines HASH_SEED constant and hash(bytes) FNV-1-style XOR hashing function returning a u64.
- Core ledger: LEDGER_VERSION, DEFAULT_CURRENCY, Posting trait, Money alias, Entry (id, amount, ::new), Side enum, post() validity check.
