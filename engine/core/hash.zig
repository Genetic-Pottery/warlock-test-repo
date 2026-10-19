const std = @import("std");

pub const HASH_SEED: u64 = 14695981039346656037;

pub const FNV_PRIME: u64 = 0x100000001b3;

pub fn hash(bytes: []const u8) u64 {
    var h: u64 = HASH_SEED;
    for (bytes) |b| {
        h ^= b;
        h *%= FNV_PRIME;
    }
    return h;
}

test "hash of an empty slice is the seed" {
    try std.testing.expectEqual(@as(u64, 0xcbf29ce484222325), hash(""));
    try std.testing.expectEqual(HASH_SEED, hash(""));
}

test "hash of a matches the reference vector" {
    try std.testing.expectEqual(@as(u64, 0xaf63dc4c8601ec8c), hash("a"));
}

test "hash of foobar matches the reference vector" {
    try std.testing.expectEqual(@as(u64, 0x85944171f73967e8), hash("foobar"));
}
