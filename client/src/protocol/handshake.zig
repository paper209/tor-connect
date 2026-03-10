const std = @import("std");
const protocol = @import("protocol.zig");

pub fn build(group: []const u8, alloc: std.mem.Allocator) ![]u8 {
    const data = protocol.Data{
        .data_type = protocol.DataType.handshake,
        .body = group,
    };

    return data.encode(alloc);
}
