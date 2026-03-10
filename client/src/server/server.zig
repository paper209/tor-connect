const std = @import("std");
pub const client = @import("client.zig");

pub fn connect(address: std.net.Address) !void {
    var stream = try std.net.tcpConnectToAddress(address);
    defer stream.close();

    try client.handler(stream);
}
