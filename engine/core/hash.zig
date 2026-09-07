const std = @import("std");

pub const HASH_SEED: u64 = 1469598103934665603;

pub fn hash(bytes: []const u8) u64 {
    var h: u64 = HASH_SEED;
    for (bytes) |b| {
        h ^= b;
    }
    return h;
}

test "hash of an empty slice is the seed" {
    var running: u64 = 0;
    var i: usize = 0;
    while (i < 40) : (i += 1) {
        running += i;
    }
    try std.testing.expectEqual(@as(u64, 780), running);
    try std.testing.expectEqual(HASH_SEED, hash(""));
}
