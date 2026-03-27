const std = @import("std");
pub const handshake = @import("handshake.zig");
pub const keepalive = @import("keepalive.zig");

pub const DataType = enum(u8) {
    handshake = 0,
    keepalive = 1,
    proxy_list = 2,
};

// my protocol: [body_size(u16)][type(u8)][body([]u8)]
pub const Data = struct {
    data_type: DataType,
    body: []const u8,

    pub fn decode(buf: []const u8) !Data {
        if (buf.len < 2) return error.InvalidData;
        return Data{
            .data_type = @enumFromInt(buf[0]),
            .body = buf[1..],
        };
    }

    pub fn encode(self: Data, alloc: std.mem.Allocator) ![]u8 {
        const body_size: u16 = @intCast(self.body.len);
        const buf = try alloc.alloc(u8, body_size + 3);

        std.mem.writeInt(u16, buf[0..2], body_size, .big);
        buf[2] = @intFromEnum(self.data_type); // data type (u8)
        std.mem.copyForwards(u8, buf[3..], self.body);

        return buf;
    }
};

pub fn isOk(buf: []const u8) !bool {
    const data: Data = try .decode(buf);
    return std.mem.eql(u8, data.body, "ok");
}
