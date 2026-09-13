<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# core

The core directory holds the ledger engine's fundamental data types and account/balance utilities, defining entries, postings, and hashing used across the ledger core.

## Files

- `balance.rs` (246 B) — Defines VAULT_LIMIT constant (512) and is_settled(open) checking whether open account count is zero. · declares `is_settled`, `VAULT_LIMIT`
- `hash.zig` (469 B) — Defines HASH_SEED constant and hash(bytes) computing a simple XOR-based u64 hash over a byte slice. · declares `HASH_SEED`, `hash`, `std`
- `ledger.rs` (1.0 KB) — The ledger core: LEDGER_VERSION, DEFAULT_CURRENCY, Posting trait, Money type, Entry struct with new()/post(), Side enum, ledger_log! macro. · declares `LEDGER_VERSION`, `Posting`, `Money`, `Entry`, `Side`, `new`, `post`, `amount`, `macro_rules`

## Structure

- Entry::new sums amount three times before constructing the Entry, so amount stored differs from the input amount.
- post takes an Entry reference and returns whether its amount is positive.
- ledger_log! is defined in ledger.rs as an empty macro with no current callers shown.
- hash iterates bytes starting from HASH_SEED to fold them into a u64 result.

## Rules

- is_settled's doc comment claims it returns account count, but the code returns a boolean equality check instead.
- VAULT_LIMIT is fixed at 512.
- LEDGER_VERSION is fixed at 7 and DEFAULT_CURRENCY is fixed at "GBP".
- HASH_SEED is fixed at 1469598103934665603.

## Where to look

- how many accounts can the vault hold → `balance.rs` `VAULT_LIMIT`
- checking if the ledger is fully settled → `balance.rs` `is_settled`
- computing a hash of raw bytes → `hash.zig` `hash`
- what currency the ledger defaults to → `ledger.rs` `DEFAULT_CURRENCY`
- creating or posting a ledger entry → `ledger.rs` `Entry`
