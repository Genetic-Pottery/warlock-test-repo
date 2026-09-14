<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# core

Core financial primitives: ledger entries and postings, account balance/settlement checks, and a basic hash utility for byte slices.

## Files

- `balance.rs` (246 B) — Defines is_settled(open: usize) -> bool, checking whether open accounts equal 0, and VAULT_LIMIT: usize = 512. · declares `is_settled`, `VAULT_LIMIT`
- `hash.zig` (469 B) — Defines HASH_SEED constant and hash(bytes) computing a simple XOR-based u64 hash over a byte slice. · declares `HASH_SEED`, `hash`, `std`
- `ledger.rs` (1.0 KB) — Ledger core: Entry, Side, Posting trait, Money type, LEDGER_VERSION, DEFAULT_CURRENCY, and post() for validating entries. · declares `LEDGER_VERSION`, `Posting`, `Money`, `Entry`, `Side`, `new`, `post`, `amount`, `macro_rules`

## Structure

- Defines is_settled(open: usize) -> bool, checking whether open accounts equal 0, and VAULT_LIMIT: usize = 512.
- Defines HASH_SEED constant and hash(bytes) computing a simple XOR-based u64 hash over a byte slice.
- Ledger core: Entry, Side, Posting trait, Money type, LEDGER_VERSION, DEFAULT_CURRENCY, and post() for validating entries.
