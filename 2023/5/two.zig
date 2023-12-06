const std = @import("std");
const aoc = @import("aoc.zig");
const one = @import("one.zig");
const print = std.debug.print;
const ArrayList = std.ArrayList;

pub fn main() !void {
    const start = try std.time.Instant.now();

    var buffer: [2 * 1024 * 1024]u8 = undefined;
    var fba = std.heap.FixedBufferAllocator.init(&buffer);
    const alloc = fba.allocator();

    const input = try one.parseInput(alloc);
    var i: usize = 0;

    var seeds = try ArrayList([2]i64).initCapacity(alloc, input.seeds.items.len / 2);
    while (i < input.seeds.items.len) {
        try seeds.append([2]i64{
            input.seeds.items[i],
            input.seeds.items[i] + input.seeds.items[i + 1] - 1,
        });
        i += 2;
    }

    var splits = try ArrayList([2]i64).initCapacity(alloc, 2);
    var min: i64 = std.math.maxInt(i64);
    for (seeds.items) |seed| {
        splits.items.len = 0;
        try splits.append(seed);
        for (input.maps.items) |map| {
            var j: usize = 0;
            while (j < splits.items.len) : (j += 1) {
                var s = &splits.items[j];
                for (map.items) |m| {
                    if (s[1] < m[2] or s[0] > m[3]) // Not covered
                        continue;

                    if (!(s[0] >= m[2] and s[1] <= m[3])) { // Not fully covered
                        if (s[0] < m[2]) { // Left overlap, split
                            try splits.append([2]i64{ s[0], m[2] - 1 });
                            s[0] = m[2];
                        } else { // Right overlap, split
                            try splits.append([2]i64{ m[3] + 1, s[1] });
                            s[1] = m[3];
                        }
                    }

                    const offset = m[0] - m[2];
                    s[0] += offset;
                    s[1] += offset;
                    break;
                }
            }
        }
        for (splits.items) |s| {
            if (s[0] < min) min = s[0];
        }
    }

    const took = std.time.Instant.since(try std.time.Instant.now(), start);
    print("Answer {d}, took {s}\n", .{ min, std.fmt.fmtDuration(took) });
}
