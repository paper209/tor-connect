const std = @import("std");
const server = @import("server/server.zig");
const proxy = @import("proxy/proxy.zig");

const allocator = std.heap.page_allocator;

pub fn main() !void {
    const args = try std.process.argsAlloc(allocator);
    defer std.process.argsFree(allocator, args);

    if (args.len < 3) {
        std.debug.print("{s} [group_name] [proxy1] [proxy2]...\n", .{args[0]});
        return;
    }

    const group = args[1];
    try proxy.init(args[2..]);

    while (true) {
        for (proxy.tor_proxies.items) |p| {
            server.connect(group, p) catch |err| {
                std.debug.print("connect error: {}\n", .{err});
                std.Thread.sleep(10 * std.time.ns_per_s);
                continue;
            };
        }
    }
}
