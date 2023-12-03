const std = @import("std");
const aoc = @import("aoc.zig");
const ascii = std.ascii;
const debug = std.debug;
const mem = std.mem;
const fmt = std.fmt;

pub fn main() !void {
    var stdin = std.io.bufferedReader(std.io.getStdIn().reader());
    var sc = aoc.fixedBufferScanner(stdin.reader(), '\n', 4096);

    const limits: [3 * maxRounds]u8 =
        ([_]u8{12} ** maxRounds) ++ // Red
        ([_]u8{13} ** maxRounds) ++ // Green
        ([_]u8{14} ** maxRounds); // Blue

    var sum: u64 = 0;
    while (try sc.next()) |line| {
        const game = parseGame(line) orelse return;
        const vec: @Vector(limits.len, u8) = game.colors;
        if (@reduce(.And, vec <= limits))
            sum += game.id;
    }

    debug.print("{d}\n", .{sum});
}

pub const maxRounds = 8;
pub const colors = [_][]const u8{ "red", "green", "blue" };
pub const Game = struct { id: u8, colors: [maxRounds * 3]u8 };

pub fn parseGame(line: []const u8) ?Game {
    const hi = mem.indexOf(u8, line, ": ") orelse return null;
    const lo = mem.lastIndexOfScalar(u8, line[0..hi], ' ') orelse return null;
    const id = fmt.parseUnsigned(u8, line[lo + 1 .. hi], 10) catch return null;

    var game = Game{ .id = id, .colors = [_]u8{0} ** (maxRounds * 3) };
    var rounds = mem.splitSequence(u8, line[hi + 2 ..], "; ");
    var i: usize = 0;

    while (rounds.next()) |round| {
        var draws = mem.splitSequence(u8, round, ", ");
        while (draws.next()) |draw| {
            const space = mem.indexOfScalar(u8, draw, ' ') orelse return null;
            const num = fmt.parseUnsigned(u8, draw[0..space], 10) catch return null;
            const color = for (colors, 0..) |color, j| {
                if (mem.eql(u8, draw[space + 1 ..], color)) break j;
            } else return null;
            game.colors[color * maxRounds + i] = num;
        }

        i += 1;
        if (i == maxRounds) {
            debug.print("too many rounds: {s}\n", .{line});
            return null;
        }
    }

    return game;
}
