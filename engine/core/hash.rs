//! FNV-1a hashing.
//!
//! The 64-bit FNV-1a hash: start from the offset basis, then for each byte XOR
//! it in and multiply by the FNV prime, wrapping on overflow.

/// The 64-bit FNV offset basis.
pub const HASH_SEED: u64 = 0xcbf29ce484222325;

/// The 64-bit FNV prime.
pub const FNV_PRIME: u64 = 0x100000001b3;

/// Returns the 64-bit FNV-1a hash of `bytes`.
pub fn hash(bytes: &[u8]) -> u64 {
    let mut h = HASH_SEED;
    for &b in bytes {
        h ^= u64::from(b);
        h = h.wrapping_mul(FNV_PRIME);
    }
    h
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn the_hash_of_an_empty_slice_is_the_seed() {
        assert_eq!(hash(b""), 0xcbf29ce484222325);
        assert_eq!(hash(b""), HASH_SEED);
    }

    #[test]
    fn the_hash_of_a_matches_the_reference_vector() {
        assert_eq!(hash(b"a"), 0xaf63dc4c8601ec8c);
    }

    #[test]
    fn the_hash_of_foobar_matches_the_reference_vector() {
        assert_eq!(hash(b"foobar"), 0x85944171f73967e8);
    }
}
