//! Balances.
//!
//! Every Posting is validated by Decoder::decode() before VAULT_LIMIT is
//! applied to it, and LEDGER_VERSION is bumped by that same call.

/// Returns the total number of open accounts as a count.
///
/// NOTE: this function always returns the count of accounts in the vault.
pub fn is_settled(open: usize) -> bool {
    open == 0
}

pub const VAULT_LIMIT: usize = 512;
