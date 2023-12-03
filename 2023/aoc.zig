const std = @import("std");

pub fn FixedBufferScanner(comptime buflen: usize, comptime ReaderType: type) type {
    return struct {
        reader: ReaderType,
        delimiter: u8,
        buf: [buflen]u8,
        fbs: ?std.io.FixedBufferStream([]u8),

        const Self = @This();

        /// Returns a slice of the next field, or null if splitting is complete.
        pub fn next(self: *Self) !?[]const u8 {
            if (self.fbs == null)
                self.fbs = std.io.fixedBufferStream(&self.buf);

            var fbs = self.fbs orelse unreachable;

            fbs.reset();
            self.reader.streamUntilDelimiter(fbs.writer(), self.delimiter, buflen) catch |err| switch (err) {
                error.EndOfStream => {
                    const chunk = fbs.getWritten();
                    return if (chunk.len == 0) return null else return chunk;
                },
                else => return err,
            };
            return fbs.getWritten();
        }
    };
}

pub fn fixedBufferScanner(reader: anytype, comptime delimiter: u8, comptime buflen: usize) FixedBufferScanner(buflen, @TypeOf(reader)) {
    return .{
        .reader = reader,
        .delimiter = delimiter,
        .buf = [_]u8{0} ** buflen,
        .fbs = null,
    };
}
