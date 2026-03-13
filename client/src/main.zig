const std = @import("std");
const server = @import("server/server.zig");
const proxy = @import("proxy/proxy.zig");

pub fn main() !void {
    try proxy.init();

    const address = try std.net.Address.parseIp("127.0.0.1", 8080);
    while (true) {
        server.connect(address) catch |err| {
            std.debug.print("connect error: {}\n", .{err});
            std.Thread.sleep(10 * std.time.ns_per_s);
            continue;
        };
    }
}
