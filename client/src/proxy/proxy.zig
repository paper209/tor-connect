const std = @import("std");
const allocator = std.heap.page_allocator;

pub const socks5 = @import("socks5.zig");

pub var tor_proxies: std.ArrayList([]const u8) = undefined;

pub fn init(proxies: []const []const u8) !void {
    tor_proxies = try .initCapacity(allocator, 0);
    for (proxies) |proxy| {
        try tor_proxies.append(allocator, proxy);
    }
}

pub fn update_proxies(data: []const u8) !void {
    tor_proxies.deinit(allocator);
    tor_proxies = try .initCapacity(allocator, 0);

    var proxies = std.mem.splitScalar(u8, data, ' ');
    while (proxies.next()) |proxy| {
        const proxy_copy = try allocator.dupe(u8, proxy);
        try tor_proxies.append(allocator, proxy_copy);
    }
}
