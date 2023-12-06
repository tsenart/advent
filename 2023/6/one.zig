const std = @import("std");
const aoc = @import("aoc.zig");
const mem = std.mem;
const fmt = std.fmt;
const print = std.debug.print;
const ArrayList = std.ArrayList;

pub fn main() !void {
    const start = try std.time.Instant.now();

    var buffer: [2 * 1024 * 1024]u8 = undefined;
    var fba = std.heap.FixedBufferAllocator.init(&buffer);
    const alloc = fba.allocator();
    var input = try parseInput(alloc);

    var product: u64 = 1;
    for (input.time, 0..) |time, i| {
        var ways: u16 = 0;
        for (1..time + 1) |t|
            ways += @intFromBool(t * (time - t) > input.dist[i]);
        product *= ways;
    }

    const took = std.time.Instant.since(try std.time.Instant.now(), start);
    print("Answer {d}, took {s}\n", .{ product, std.fmt.fmtDuration(took) });
}

pub const Input = struct {
    time: []u16,
    dist: []u16,
};

pub fn parseInput(alloc: std.mem.Allocator) !Input {
    var stdin = std.io.bufferedReaderSize(1024, std.io.getStdIn().reader());
    var sc = aoc.fixedBufferScanner(stdin.reader(), '\n', 1024);

    var nums = ArrayList(u16).init(alloc);
    while (try sc.next()) |line| {
        const sep = mem.indexOfScalar(u8, line, ':') orelse continue;
        var ns = mem.tokenizeScalar(u8, line[sep + 1 ..], ' ');
        while (ns.next()) |num| {
            const n = try fmt.parseInt(u16, num, 10);
            try nums.append(n);
        }
    }

    const half = nums.items.len / 2;
    return Input{ .time = nums.items[0..half], .dist = nums.items[half..] };
}
