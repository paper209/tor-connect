const std = @import("std");

const config = @import("../config.zig");
pub const client = @import("client.zig");
pub const proxy = @import("../proxy/proxy.zig");

pub fn connect(group: []const u8, p: []const u8) !void {
    var stream = try proxy.socks5.connect(p, config.server_address, config.server_port);
    defer stream.close();

    try client.handler(stream, group);
}
