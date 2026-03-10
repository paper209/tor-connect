const std = @import("std");

pub const DataType = enum(u8) {
    handshake = 0,
    keepalive = 1,
    proxy_list = 2,
};

// my protocol: [body_size(u8)][type(u8)][body([]u8)]
pub const Data = struct {
    data_type: DataType,
    body: []const u8,

    pub fn decode(buf: []const u8) Data {
        return Data{
            .data_type = @enumFromInt(buf[1]),
            .body = buf[2 .. buf[0] + 2],
        };
    }

    pub fn encode(self: Data, alloc: std.mem.Allocator) ![]u8 {
        const body_size: u8 = self.body.len;
        const buf = try alloc.alloc(u8, body_size + 2);

        buf[0] = body_size; // body size (u8)
        buf[1] = self.data_type; // data type (u8)
        std.mem.copyForwards(u8, buf[2..], self.body);

        return buf;
    }
};

pub fn isOk(buf: []const u8) bool {
    const data: Data = .decode(buf);
    return std.mem.eql([]const u8, data.body, "ok");
}

pub fn buildHandshake(alloc: std.mem.Allocator) ![]u8 {
    const data = Data{
        .data_type = DataType.handshake,
        .body = []const u8{},
    };

    return data.encode(alloc);
}

pub fn buildKeepAlive(alloc: std.mem.Allocator) ![]u8 {
    const data = Data{
        .data_type = DataType.keepalive,
        .body = []const u8{},
    };

    return data.encode(alloc);
}
