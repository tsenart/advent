const std = @import("std");
const print = std.debug.print;
const startsWith = std.mem.startsWith;
const isAsciiDigit = std.ascii.isDigit;

pub fn main() !void {
    var buf: [4096]u8 = undefined;
    var fbs = std.io.fixedBufferStream(&buf);
    const writer = fbs.writer();
    const stdin = std.io.getStdIn().reader();
    var sum: u64 = 0;

    while (true) {
        stdin.streamUntilDelimiter(writer, '\n', buf.len) catch |err| switch (err) {
            error.EndOfStream => if (fbs.getWritten().len == 0) break,
            else => {
                print("Error: {s}\n", .{@errorName(err)});
                return;
            },
        };

        const line = fbs.getWritten();
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

        fbs.reset();
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
