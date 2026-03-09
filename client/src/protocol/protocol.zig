const std = @import("std");

pub const DataType = enum(u8) {
    handshake = 0,
    keepalive = 1,
    proxy_list = 2,
};

// my protocol: [body_size(u8)][type(u8)][body([]u8)]
pub const Data = struct {
    data_type: DataType,
    data: []const u8,

    pub fn decode(data: []u8) Data {
        return Data{
            .data_type = @enumFromInt(data[1]),
            .data = data[2 .. data[0] + 2],
        };
    }

    pub fn encode(self: Data, alloc: std.mem.Allocator) ![]u8 {
        const body_size: u8 = self.data.len;
        const buf = try alloc.alloc(u8, body_size + 2);

        buf[0] = body_size; // body size (u8)
        buf[1] = self.data_type; // data type (u8)
        std.mem.copyForwards(u8, buf[2..], self.data);

        return buf;
    }
};
