const std = @import("std");
const aoc = @import("aoc.zig");
const isDigit = std.ascii.isDigit;
const print = std.debug.print;

pub fn main() !void {
    var sum: u64 = 0;
    var stdin = std.io.bufferedReader(std.io.getStdIn().reader());
    var sc = aoc.fixedBufferScanner(stdin.reader(), '\n', 4096);
    while (try sc.next()) |line| {
        var l: usize = 0;
        var r: usize = line.len - 1;
        while (l <= r) {
            if (isDigit(line[l]) and isDigit(line[r])) break;
            if (!isDigit(line[l])) l += 1;
            if (!isDigit(line[r])) r -= 1;
        }
        if (l <= r)
            sum += 10 * (line[l] - '0') + (line[r] - '0');
    }
    print("{d}\n", .{sum});
}
