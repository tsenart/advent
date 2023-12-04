const std = @import("std");
const aoc = @import("aoc.zig");
const IntegerBitSet = std.bit_set.IntegerBitSet;
const mem = std.mem;
const fmt = std.fmt;
const print = std.debug.print;

pub fn main() !void {
    const start = try std.time.Instant.now();
    var stdin = std.io.bufferedReaderSize(1024, std.io.getStdIn().reader());
    var sc = aoc.fixedBufferScanner(stdin.reader(), '\n', 1024);

    var sum: u64 = 0;
    while (try sc.next()) |line| {
        const card = try parseCard(line) orelse return;
        const wins = card.winning.intersectWith(card.have).count();
        if (wins > 0) {
            sum += std.math.pow(u64, 2, wins - 1);
        }
    }

    const took = std.time.Instant.since(try std.time.Instant.now(), start);
    print("Answer {d}, took {s}\n", .{ sum, std.fmt.fmtDuration(took) });
}

pub const Card = struct { winning: IntegerBitSet(100), have: IntegerBitSet(100) };

pub fn parseCard(line: []const u8) !?Card {
    const sep = mem.indexOf(u8, line, ":") orelse return null;
    var card = Card{ .winning = IntegerBitSet(100).initEmpty(), .have = IntegerBitSet(100).initEmpty() };
    var numbers = mem.tokenizeScalar(u8, line[sep + 1 ..], ' ');
    var set = &card.winning;
    while (numbers.next()) |number| {
        if (std.mem.eql(u8, number, "|")) {
            set = &card.have;
            continue;
        }
        const num = try fmt.parseUnsigned(usize, std.mem.trim(u8, number, " "), 10);
        set.set(num);
    }
    return card;
}
