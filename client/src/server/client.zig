const std = @import("std");

const proxy = @import("../proxy/proxy.zig");
const protocol = @import("../protocol/protocol.zig");

const allocator = std.heap.page_allocator;

pub fn sendHandshake(stream: std.net.Stream, group: []const u8) !bool {
    const handshake = try protocol.handshake.build(group, allocator);
    try stream.writeAll(handshake);

    var buf: [4]u8 = undefined;
    _ = try stream.read(&buf);

    return protocol.isOk(buf[1..]);
}

pub fn sendKeepalive(stream: std.net.Stream) !void {
    const keepalive = try protocol.keepalive.build(allocator);
    try stream.writeAll(keepalive);
}

pub fn handler(stream: std.net.Stream, group: []const u8) !void {
    // handshake loop
    while (true) {
        if (try sendHandshake(stream, group)) {
            break;
        }

        std.Thread.sleep(5 * std.time.ns_per_s);
    }

    // keepalive loop
    while (true) {
        try sendKeepalive(stream);

        var body_size: [1]u8 = undefined;
        _ = try stream.read(&body_size);

        const buf = try allocator.alloc(u8, body_size[0] + 1);
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
