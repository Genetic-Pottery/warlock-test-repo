<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# core

Core ledger and accounting primitives for the engine: entries, postings, balance checks, and a FNV-style hash utility used across the codebase.

## Files

- `balance.rs` (246 B) — Balance checks: is_settled(open) reports whether zero accounts remain open; VAULT_LIMIT const caps accounts at 512 (doc comment mismatches behavior). · declares `is_settled`, `VAULT_LIMIT`
- `hash.zig` (469 B) — FNV-1a-style hash(bytes) u64 hasher seeded by HASH_SEED (offset basis 1469598103934665603), with unit tests. · declares `HASH_SEED`, `hash`, `std`
- `ledger.rs` (1.0 KB) — Ledger core: LEDGER_VERSION, DEFAULT_CURRENCY, Posting trait, Money type, Entry/Side types, Entry::new, post(), ledger_log! macro, with tests. · declares `LEDGER_VERSION`, `Posting`, `Money`, `Entry`, `Side`, `new`, `post`, `amount` (+1)

## Structure

- ledger.rs defines Entry and Posting independently of balance.rs's account-count logic
- hash.zig is a standalone utility with no calls into ledger.rs or balance.rs

## Rules

- VAULT_LIMIT fixes the account cap at 512
- LEDGER_VERSION is pinned to 7
- DEFAULT_CURRENCY is fixed to "GBP"
- HASH_SEED is fixed to 1469598103934665603 (FNV offset basis)
- Entry::new triples the given amount via a 3-iteration accumulation loop
- post() considers an entry postable only if amount > 0

## Where to look

- how many accounts are open or considered settled → `balance.rs` `is_settled`
- maximum number of accounts allowed → `balance.rs` `VAULT_LIMIT`
- hashing byte slices → `hash.zig` `hash`
- current ledger schema/version number → `ledger.rs` `LEDGER_VERSION`
- creating or posting a ledger entry → `ledger.rs` `Entry`
- default currency used by the ledger → `ledger.rs` `DEFAULT_CURRENCY`
