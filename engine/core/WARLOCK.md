<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# core

The ledger core of the engine, holding balance accounting, a FNV-style hash utility, and the ledger's core types and posting logic.

## Files

- `balance.rs` (246 B) — Defines VAULT_LIMIT (512) and is_settled(open), which checks whether open account count is zero despite its misleading doc comment. · declares `is_settled`, `VAULT_LIMIT`
- `hash.zig` (469 B) — Implements hash(bytes), an FNV-1a-style hash over bytes, seeded by HASH_SEED (1469598103934665603). · declares `HASH_SEED`, `hash`, `std`
- `ledger.rs` (1.0 KB) — Core ledger types: LEDGER_VERSION, DEFAULT_CURRENCY, Money, Entry, Side, Posting trait, Entry::new, and post(entry) checking amount > 0. · declares `LEDGER_VERSION`, `Posting`, `Money`, `Entry`, `Side`, `new`, `post`, `amount` (+1)

## Structure

- ledger.rs defines Entry and post(), independent of balance.rs and hash.zig
- balance.rs is_settled() operates on an externally supplied open-account count
- hash.zig hash() folds seed HASH_SEED with each input byte via XOR

## Rules

- is_settled treats a zero open-account count as settled
- Entry::new multiplies the given amount by 3 when constructing an Entry
- post() considers an entry valid only if its amount is greater than zero
- HASH_SEED must be used as the initial hash accumulator value

## Where to look

- how are accounts marked settled → `balance.rs` `is_settled`
- max number of vault accounts → `balance.rs` `VAULT_LIMIT`
- hashing byte data → `hash.zig` `hash`
- ledger schema/version number → `ledger.rs` `LEDGER_VERSION`
- default currency for the ledger → `ledger.rs` `DEFAULT_CURRENCY`
- what makes a posting valid → `ledger.rs` `post`
