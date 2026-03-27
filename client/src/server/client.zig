const std = @import("std");

const proxy = @import("../proxy/proxy.zig");
const protocol = @import("../protocol/protocol.zig");

const allocator = std.heap.page_allocator;

pub fn sendHandshake(stream: std.net.Stream, group: []const u8) !bool {
    const handshake = try protocol.handshake.build(group, allocator);
    try stream.writeAll(handshake);

    var buf: [5]u8 = undefined;
    _ = try stream.read(&buf);

    return protocol.isOk(buf[2..]);
}

pub fn sendKeepalive(stream: std.net.Stream) !void {
    const keepalive = try protocol.keepalive.build(allocator);
    try stream.writeAll(keepalive);
}

pub fn handler(stream: std.net.Stream, group: []const u8) !void {
    // handshake loop
    const stat = try sendHandshake(stream, group);
    if (!stat) return error.HandshakeError;

    // keepalive loop
    while (true) {
        try sendKeepalive(stream);

        var size_buf: [2]u8 = undefined;
        _ = try stream.read(&size_buf);

        const body_size = std.mem.readInt(u16, size_buf[0..], .big);

        const buf = try allocator.alloc(u8, body_size + 1);
        _ = try stream.read(buf);

        const data: protocol.Data = try .decode(buf);
        switch (data.data_type) {
            protocol.DataType.keepalive => {},
            protocol.DataType.proxy_list => {
                try proxy.update_proxies(data.body);
            },
            else => {},
        }

        allocator.free(buf);
        std.Thread.sleep(5 * std.time.ns_per_s);
    }
}
