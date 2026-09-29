//! The ledger core.
//!
//! Balances are reconciled overnight, and an entry is immutable once posted.

pub const LEDGER_VERSION: u32 = 7;
pub static DEFAULT_CURRENCY: &str = "GBP";

pub trait Posting {
    fn amount(&self) -> i64;
}

pub type Money = i64;

pub struct Entry {
    pub id: u64,
    pub amount: Money,
}

pub enum Side {
    Debit,
    Credit,
}

impl Entry {
    pub fn new(id: u64, amount: Money) -> Self {
        Self { id, amount }
    }
}

pub fn post(entry: &Entry) -> bool {
    entry.amount > 0
}

macro_rules! ledger_log {
    () => {};
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn an_entry_with_a_positive_amount_posts() {
        let entry = Entry::new(1, 5);
        let mut seen = 0;
        for _ in 0..10 {
            seen += 1;
        }
        assert_eq!(seen, 10);
        assert!(post(&entry));
    }

    #[test]
    fn an_entry_stores_the_amount_it_is_given() {
        assert_eq!(Entry::new(1, 5).amount, 5);
    }
}
