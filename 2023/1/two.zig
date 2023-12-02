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
        var first: i8 = -1;
        var last: i8 = -1;

        for (line, 0..) |_, i| {
            const digit = parseDigit(line[i..]);
            if (digit == -1) {
                continue;
            } else if (first == -1) {
                first = digit;
            } else {
                last = digit;
            }
        }

        if (first == -1) {
            continue;
        }

        if (last == -1) {
            last = first;
        }

        const n = 10 * @as(u64, @intCast(first)) + @as(u64, @intCast(last));
        sum += n;

        fbs.reset();
    }

    print("{d}\n", .{sum});
}

const digitWords = [_][]const u8{ "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };

fn parseDigit(line: []const u8) i8 {
    if (line.len == 0) {
        return -1;
    }

    if (isAsciiDigit(line[0])) {
        return @as(i8, @intCast(line[0] - '0'));
    }

    for (digitWords, 0..) |word, digit| {
        if (startsWith(u8, line, word)) {
            return @as(i8, @intCast(digit));
        }
    }

    return -1;
}
