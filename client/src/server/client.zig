const std = @import("std");
const protocol = @import("../protocol/protocol.zig");

pub fn sendHandshake(stream: std.net.Stream, group: []const u8) !bool {
    const handshake = try protocol.handshake.build(group, std.heap.page_allocator);
    try stream.writeAll(handshake);

    var buf: [4]u8 = undefined;
    _ = try stream.read(&buf);

    return protocol.isOk(buf[0..]);
}

pub fn sendKeepalive(stream: std.net.Stream) !bool {
    const keepalive = try protocol.keepalive.build(std.heap.page_allocator);
    try stream.writeAll(keepalive);

    var buf: [4]u8 = undefined;
    _ = try stream.read(&buf);

    return protocol.isOk(buf[0..]);
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
        if (try sendKeepalive(stream)) {
            std.Thread.sleep(3 * std.time.ns_per_s);
            continue;
        }

        std.Thread.sleep(5 * std.time.ns_per_s);
    }
}
