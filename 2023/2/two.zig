const std = @import("std");
const aoc = @import("aoc.zig");
const one = @import("one.zig");
const debug = std.debug;

pub fn main() !void {
    const start = try std.time.Instant.now();

    var stdin = std.io.bufferedReader(std.io.getStdIn().reader());
    var sc = aoc.fixedBufferScanner(stdin.reader(), '\n', 4096);
    var sum: u64 = 0;
    const rounds = one.maxRounds;
    while (try sc.next()) |line| {
        const game = one.parseGame(line) orelse return;
        var power: u64 = 1;
        for (0..3) |i| {
            power *= std.mem.max(u8, game.colors[i * rounds .. (i + 1) * rounds]);
        }
        sum += power;
    }

    const took = std.time.Instant.since(try std.time.Instant.now(), start);
    debug.print("Answer {d}, took {s}\n", .{ sum, std.fmt.fmtDuration(took) });
}
