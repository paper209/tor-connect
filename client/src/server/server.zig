const std = @import("std");
pub const client = @import("client.zig");
pub const proxy = @import("../proxy/proxy.zig");

const server_address = ".onion";
const server_port: u16 = 2222;

pub fn connect(group: []const u8, p: []const u8) !void {
    var stream = try proxy.socks5.connect(p, server_address, server_port);
    defer stream.close();

    try client.handler(stream, group);
}
