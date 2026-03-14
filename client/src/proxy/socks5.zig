const std = @import("std");

pub fn connect(proxy: []const u8, domain: []const u8, port: u16) !std.net.Stream {
    var it = std.mem.splitScalar(u8, proxy, ':');
    const addr = it.next() orelse return error.InvalidAddress;

    const port_str = it.next() orelse return error.InvalidAddress;
    const proxy_port = try std.fmt.parseInt(u16, port_str, 10);

    const address = try std.net.Address.parseIp(addr, proxy_port);
    var stream = try std.net.tcpConnectToAddress(address);
    errdefer stream.close();

    // set method
    var method: [3]u8 = undefined;
    method[0] = 0x05; // version (socks5)
    method[1] = 0x01; // nmethods (1)
    method[2] = 0x00; // methods (no auth)

    try stream.writeAll(&method);

    // read method reply
    var method_reply: [2]u8 = undefined;
    _ = try stream.read(&method_reply);
    if (method_reply[1] != 0x00) return error.AuthError;

    // connect request
    var connect_req: [5]u8 = undefined;
    connect_req[0] = 0x05; // version (socks5)
    connect_req[1] = 0x01; // cmd (connect)
    connect_req[2] = 0x00; // reserved
    connect_req[3] = 0x03; // atyp (domain)
    connect_req[4] = @intCast(domain.len); // server domain length
    try stream.writeAll(&connect_req);
    try stream.writeAll(domain);

    var connect_req_port: [2]u8 = undefined;
    connect_req_port[0] = @intCast(port >> 8);
    connect_req_port[1] = @intCast(port & 0xff);
    try stream.writeAll(&connect_req_port);

    var connect_reply: [10]u8 = undefined;
    _ = try stream.read(&connect_reply);
    if (connect_reply[1] != 0x00) return error.ConnectFailed;

    return stream;
}
