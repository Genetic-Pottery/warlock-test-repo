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

    /// Returns a new entry under `id` carrying the negated amount.
    ///
    /// Returns `None` when the amount cannot be negated (`i64::MIN`).
    pub fn reverse(&self, id: u64) -> Option<Entry> {
        self.amount.checked_neg().map(|amount| Entry { id, amount })
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

    #[test]
    fn reversing_an_entry_negates_the_amount_under_the_new_id() {
        let entry = Entry::new(1, 5);
        let reversed = entry.reverse(2).expect("5 negates");
        assert_eq!(reversed.id, 2);
        assert_eq!(reversed.amount, -5);
        assert_eq!(entry.id, 1);
        assert_eq!(entry.amount, 5);

        let credit = Entry::new(3, -7);
        let reversed = credit.reverse(4).expect("-7 negates");
        assert_eq!(reversed.id, 4);
        assert_eq!(reversed.amount, 7);
    }

    #[test]
    fn reversing_the_minimum_amount_is_none() {
        assert!(Entry::new(1, i64::MIN).reverse(2).is_none());
    }
}
