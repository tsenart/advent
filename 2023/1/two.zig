const std = @import("std");
const aoc = @import("aoc.zig");
const print = std.debug.print;
const startsWith = std.mem.startsWith;
const isAsciiDigit = std.ascii.isDigit;

pub fn main() !void {
    var stdin = std.io.bufferedReader(std.io.getStdIn().reader());
    var sc = aoc.fixedBufferScanner(stdin.reader(), '\n', 4096);
    var sum: u64 = 0;
    while (try sc.next()) |line| {
        var first: ?u8 = null;
        var last: ?u8 = null;

        for (line, 0..) |_, i| {
            if (parseDigit(line[i..])) |digit| {
                if (first == null) {
                    first = digit;
                } else {
                    last = digit;
                }
            }
        }

        const f = first orelse continue;
        const l = last orelse f;
        sum += 10 * f + l;
    }

    print("{d}\n", .{sum});
}

const digitWords = [_][]const u8{ "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };

fn parseDigit(line: []const u8) ?u8 {
    if (line.len == 0) return null;
    if (isAsciiDigit(line[0])) return line[0] - '0';
    for (digitWords, 0..) |word, digit| {
        if (startsWith(u8, line, word)) return @intCast(digit);
    }
    return null;
}
