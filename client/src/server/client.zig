const std = @import("std");
const protocol = @import("../protocol/protocol.zig");

pub fn sendHandshake(stream: std.net.Stream, group: []const u8) !bool {
    const handshake = try protocol.handshake.build(group, std.heap.page_allocator);
    try stream.writeAll(handshake);

    var buf: [4]u8 = undefined;
    _ = try stream.read(&buf);

    return protocol.isOk(buf[0..]);
}

pub fn sendKeepalive(stream: std.net.Stream) !void {
    const keepalive = try protocol.keepalive.build(std.heap.page_allocator);
    try stream.writeAll(keepalive);
}

pub fn handler(stream: std.net.Stream) !void {
    // handshake loop
    while (true) {
        if (try sendHandshake(stream, "test")) {
            break;
        }

        std.Thread.sleep(5 * std.time.ns_per_s);
    }

    // keepalive loop
    while (true) {
        try sendKeepalive(stream);

        var buf: [2]u8 = undefined;
        _ = try stream.read(&buf);
        switch (buf[1]) {
            @intFromEnum(protocol.DataType.keepalive) => {},

            @intFromEnum(protocol.DataType.proxy_list) => {
                var data_buf = try std.heap.page_allocator.alloc(u8, buf[0]);
                defer std.heap.page_allocator.free(data_buf);
                _ = try stream.read(data_buf[0..]);

                std.debug.print("{s}\n", .{data_buf});
            },
            else => {},
        }

        std.Thread.sleep(5 * std.time.ns_per_s);
    }
}
