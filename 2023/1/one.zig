const std = @import("std");
const isDigit = std.ascii.isDigit;
const print = std.debug.print;

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

        var l: usize = 0;
        var r: usize = line.len - 1;
        while (l <= r) {
            if (isDigit(line[l]) and isDigit(line[r])) break;
            if (!isDigit(line[l])) l += 1;
            if (!isDigit(line[r])) r -= 1;
        }

        if (l > r) continue;

        sum += 10 * (line[l] - '0') + (line[r] - '0');
        fbs.reset();
    }

    print("{d}\n", .{sum});
}
