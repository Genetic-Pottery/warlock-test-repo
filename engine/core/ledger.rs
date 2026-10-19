//! The ledger core.
//!
//! Balances are reconciled overnight, and an entry is immutable once posted.

mod hash;

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

/// One entry in a [`Chain`], linked to the link before it by hash.
pub struct Link {
    pub entry: Entry,
    pub prev_hash: u64,
    pub hash: u64,
}

/// An append-only, hash-linked sequence of entries.
pub struct Chain {
    links: Vec<Link>,
}

/// Returns the hash of a link: `prev_hash`, `id`, then `amount`, each as 8
/// little-endian bytes.
fn link_hash(prev_hash: u64, entry: &Entry) -> u64 {
    let mut bytes = [0u8; 24];
    bytes[0..8].copy_from_slice(&prev_hash.to_le_bytes());
    bytes[8..16].copy_from_slice(&entry.id.to_le_bytes());
    bytes[16..24].copy_from_slice(&entry.amount.to_le_bytes());
    hash::hash(&bytes)
}

impl Chain {
    /// Returns an empty chain.
    pub fn new() -> Self {
        Self { links: Vec::new() }
    }

    /// Appends `entry` and returns the new link's hash.
    ///
    /// Accepts any amount, including negative amounts from `Entry::reverse`.
    pub fn append(&mut self, entry: Entry) -> u64 {
        let prev_hash = self.links.last().map_or(0, |link| link.hash);
        let hash = link_hash(prev_hash, &entry);
        self.links.push(Link { entry, prev_hash, hash });
        hash
    }

    /// Returns `Err(i)` for the lowest index `i` whose stored `prev_hash` or
    /// `hash` doesn't match what it recomputes, or `Ok(())` if none.
    pub fn verify(&self) -> Result<(), usize> {
        let mut prev_hash = 0;
        for (i, link) in self.links.iter().enumerate() {
            if link.prev_hash != prev_hash || link.hash != link_hash(prev_hash, &link.entry) {
                return Err(i);
            }
            prev_hash = link.hash;
        }
        Ok(())
    }
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

    fn three_link_chain() -> Chain {
        let mut chain = Chain::new();
        let first = Entry::new(1, 5);
        let reversed = first.reverse(2).expect("5 negates");
        chain.append(first);
        chain.append(reversed);
        chain.append(Entry::new(3, 9));
        chain
    }

    #[test]
    fn an_empty_chain_verifies() {
        assert!(Chain::new().links.is_empty());
        assert_eq!(Chain::new().verify(), Ok(()));
    }

    #[test]
    fn a_valid_three_link_chain_verifies() {
        let mut chain = Chain::new();
        let h0 = chain.append(Entry::new(1, 5));
        let h1 = chain.append(Entry::new(1, 5).reverse(2).expect("5 negates"));
        let h2 = chain.append(Entry::new(3, 9));

        assert_eq!(chain.links.len(), 3);
        assert_eq!(chain.links[0].prev_hash, 0);
        assert_eq!(chain.links[1].prev_hash, h0);
        assert_eq!(chain.links[2].prev_hash, h1);
        assert_eq!(chain.links[0].hash, h0);
        assert_eq!(chain.links[1].hash, h1);
        assert_eq!(chain.links[2].hash, h2);
        assert_eq!(chain.links[1].entry.amount, -5);

        let mut bytes = Vec::new();
        bytes.extend_from_slice(&0u64.to_le_bytes());
        bytes.extend_from_slice(&1u64.to_le_bytes());
        bytes.extend_from_slice(&5i64.to_le_bytes());
        assert_eq!(h0, hash::hash(&bytes));

        assert_eq!(chain.verify(), Ok(()));
    }

    #[test]
    fn an_edited_amount_fails_at_its_link() {
        let mut chain = three_link_chain();
        chain.links[1].entry.amount = 500;
        assert_eq!(chain.verify(), Err(1));
    }

    #[test]
    fn an_edited_prev_hash_fails_at_its_link() {
        let mut chain = three_link_chain();
        chain.links[2].prev_hash ^= 1;
        assert_eq!(chain.verify(), Err(2));
    }
}
