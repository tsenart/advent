const std = @import("std");
const aoc = @import("aoc.zig");
const print = std.debug.print;
const isDigit = std.ascii.isDigit;

pub fn main() !void {
    const start = try std.time.Instant.now();

    var buffer: [2 * 1024 * 1024]u8 = undefined;
    var fba = std.heap.FixedBufferAllocator.init(&buffer);
    const alloc = fba.allocator();

    const size = 256;
    var stdin = std.io.bufferedReaderSize(size, std.io.getStdIn().reader());
    var sc = aoc.fixedBufferScanner(stdin.reader(), '\n', size);

    const digits = "0123456789";
    const number = struct {
        n: u16,
        counted: bool,
    };

    var cols: usize = 0;
    var row: usize = 0;
    var nums = std.AutoHashMap(usize, *number).init(alloc);
    var symbols = std.ArrayList([2]i32).init(alloc);
    defer nums.deinit();
    defer symbols.deinit();

    while (try sc.next()) |line| {
        if (cols == 0) cols = line.len;

        var col: usize = 0;
        while (col < line.len) {
            const c = line[col];
            if (c == '.') {
                col += 1;
                continue;
            }

            if (isDigit(c)) {
                const end = std.mem.indexOfNonePos(u8, line, col + 1, digits) orelse cols;
                const n = try std.fmt.parseInt(u16, line[col..end], 10);
                var num = try alloc.create(number);
                num.n = n;
                for (col..end) |i| {
                    try nums.put(row * cols + i, num);
                }
                col = end;
                continue;
            }

            if (c == '*')
                try symbols.append([2]i32{ @intCast(row), @intCast(col) });
            col += 1;
        }

        row += 1;
    }

    const rows = row;
    var sum: u64 = 0;

    for (symbols.items) |symbol| {
        const r = symbol[0]; // row
        const c = symbol[1]; // col
        const adjacent = [_][2]i32{
            [2]i32{ r, c - 1 },
            [2]i32{ r, c + 1 },
            [2]i32{ r - 1, c },
            [2]i32{ r - 1, c - 1 },
            [2]i32{ r - 1, c + 1 },
            [2]i32{ r + 1, c },
            [2]i32{ r + 1, c - 1 },
            [2]i32{ r + 1, c + 1 },
        };

        var ratio: u64 = 1;
        var added: u64 = 0;
        for (adjacent) |pos| {
            if (pos[0] < 0 or pos[0] > rows or pos[1] < 0 or pos[1] > cols)
                continue;

            const i: usize = @as(usize, @intCast(pos[0])) * cols +
                @as(usize, @intCast(pos[1]));

            if (nums.getEntry(i)) |e| {
                if (!e.value_ptr.*.counted) {
                    ratio *= e.value_ptr.*.n;
                    e.value_ptr.*.counted = true;
                    added += 1;
                }
            }
        }

        if (added == 2)
            sum += ratio;
    }

    const took = std.time.Instant.since(try std.time.Instant.now(), start);
    print("Answer {d}, took {s}\n", .{ sum, std.fmt.fmtDuration(took) });
}
