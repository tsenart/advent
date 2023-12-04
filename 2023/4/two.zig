const std = @import("std");
const aoc = @import("aoc.zig");
const one = @import("one.zig");
const mem = std.mem;
const fmt = std.fmt;
const print = std.debug.print;

pub fn main() !void {
    const start = try std.time.Instant.now();

    var stdin = std.io.bufferedReaderSize(1024, std.io.getStdIn().reader());
    var sc = aoc.fixedBufferScanner(stdin.reader(), '\n', 1024);

    var copies: [200]u32 = mem.zeroes([200]u32);
    var id: u8 = 0;
    var sum: u64 = 0;
    while (try sc.next()) |line| {
        const card = try one.parseCard(line) orelse return;
        const wins = card.winning.intersectWith(card.have).count();
        copies[id] += 1;
        for (1..wins + 1) |i| {
            copies[id + i] += copies[id];
        }
        sum += copies[id];
        id += 1;
    }

    const took = std.time.Instant.since(try std.time.Instant.now(), start);
    print("Answer {d}, took {s}\n", .{ sum, std.fmt.fmtDuration(took) });
}
