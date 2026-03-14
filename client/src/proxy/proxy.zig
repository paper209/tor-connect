const std = @import("std");
const allocator = std.heap.page_allocator;

pub var tor_proxies: std.ArrayList([]const u8) = undefined;

pub fn init() !void {
    tor_proxies = try .initCapacity(allocator, 0);
}

pub fn update_proxies(data: []const u8) !void {
    tor_proxies.deinit(allocator);
    tor_proxies = try .initCapacity(allocator, 0);

    var proxies = std.mem.splitScalar(u8, data, ' ');
    while (proxies.next()) |proxy| {
        try tor_proxies.append(allocator, proxy);
    }

    for (tor_proxies.items) |proxy| {
        std.debug.print("{s}\n", .{proxy});
    }
}
