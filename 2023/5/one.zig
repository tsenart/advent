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

    var min: i64 = std.math.maxInt(i64);
    for (input.seeds.items) |seed| {
        var n: i64 = @intCast(seed);
        for (input.maps.items) |map| {
            for (map.items) |m| {
                if (n >= m[2] and n <= m[3]) {
                    n = n - m[2] + m[0];
                    break;
                }
            }
        }
        if (n < min) min = n;
    }

    const took = std.time.Instant.since(try std.time.Instant.now(), start);
    print("Answer {d}, took {s}\n", .{ min, std.fmt.fmtDuration(took) });
}

pub const Input = struct {
    seeds: ArrayList(i64),
    maps: ArrayList(ArrayList([4]i64)),
};

pub fn parseInput(alloc: std.mem.Allocator) !Input {
    var stdin = std.io.bufferedReaderSize(1024, std.io.getStdIn().reader());
    var sc = aoc.fixedBufferScanner(stdin.reader(), '\n', 1024);

    var seeds = ArrayList(i64).init(alloc);
    var maps = ArrayList(ArrayList([4]i64)).init(alloc);

    while (try sc.next()) |line| {
        if (line.len == 0) continue;

        if (mem.startsWith(u8, line, "seeds: ")) {
            var ns = mem.tokenizeScalar(u8, line[7..], ' ');
            while (ns.next()) |num| {
                const seed = try fmt.parseInt(i64, num, 10);
                try seeds.append(seed);
            }
            continue;
        }

        if (mem.indexOfScalar(u8, line, ':')) |_| {
            const map = ArrayList([4]i64).init(alloc);
            try maps.append(map);
            continue;
        }

        var m: [4]i64 = mem.zeroes([4]i64);
        var ns = mem.tokenizeScalar(u8, line, ' ');
        var i: usize = 0;
        while (ns.next()) |num| {
            const n = try fmt.parseInt(i64, num, 10);
            if (i < 2) {
                m[i * 2] = n;
            } else {
                m[1] = m[0] + n - 1;
                m[3] = m[2] + n - 1;
            }
            i += 1;
        }

        try maps.items[maps.items.len - 1].append(m);
    }

    return Input{ .seeds = seeds, .maps = maps };
}
