const std = @import("std");
const aoc = @import("aoc.zig");
const math = std.math;
const mem = std.mem;
const fmt = std.fmt;
const print = std.debug.print;

pub fn main() !void {
    const start = try std.time.Instant.now();
    const input = try parseInput();
    const time = input[0];
    const dist = input[1];

    // This implementation solves the problem by applying a mathematical approach
    // to a quadratic inequality t^2 - time * t + dist < 0.
    //
    // The roots of the quadratic equation (where the inequality changes its truth value)
    // are found using the standard formula for quadratic equations. The range of 't' values
    // that satisfy the inequality is then deduced from these roots.
    //
    // The approach is efficient, especially for large values of 'time' and 'dist', as it
    // avoids iterative or brute-force methods, instead directly calculating the required range.

    const discriminant = math.pow(u64, time, 2) - 4 * dist;
    const sq = math.sqrt(discriminant);
    const r1 = (time + sq + 1) / 2; // Round up
    const r2 = (time - sq) / 2;
    const ways = @max(r1, r2) - @min(r1, r2) + 1;

    const took = std.time.Instant.since(try std.time.Instant.now(), start);
    print("Answer {d}, took {s}\n", .{ ways, std.fmt.fmtDuration(took) });
}

pub fn parseInput() ![2]u64 {
    var stdin = std.io.bufferedReaderSize(1024, std.io.getStdIn().reader());
    var sc = aoc.fixedBufferScanner(stdin.reader(), '\n', 1024);
    var input: [2]u64 = mem.zeroes([2]u64);
    var i: usize = 0;
    while (try sc.next()) |line| {
        const sep = mem.indexOfScalar(u8, line, ':') orelse continue;
        var ns = mem.tokenizeScalar(u8, line[sep + 1 ..], ' ');
        while (ns.next()) |n|
            input[i] = input[i] *
                (std.math.pow(u64, 10, @as(u64, @intCast(n.len)))) +
                try fmt.parseInt(u64, n, 10);
        i += 1;
    }
    return input;
}
