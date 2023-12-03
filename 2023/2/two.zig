const std = @import("std");
const aoc = @import("aoc.zig");
const one = @import("one.zig");
const debug = std.debug;

pub fn main() !void {
    var stdin = std.io.bufferedReader(std.io.getStdIn().reader());
    var sc = aoc.fixedBufferScanner(stdin.reader(), '\n', 4096);
    var sum: u64 = 0;
    const rounds = one.maxRounds;
    while (try sc.next()) |line| {
        const game = one.parseGame(line) orelse return;
        const r: @Vector(rounds, u8) = game.colors[0..rounds].*;
        const g: @Vector(rounds, u8) = game.colors[rounds .. 2 * rounds].*;
        const b: @Vector(rounds, u8) = game.colors[2 * rounds ..].*;
        sum +=
            @as(u64, @reduce(.Max, r)) *
            @as(u64, @reduce(.Max, g)) *
            @as(u64, @reduce(.Max, b));
    }
    debug.print("{d}\n", .{sum});
}
