const std = @import("std");
const protocol = @import("protocol.zig");

pub fn build(alloc: std.mem.Allocator) ![]u8 {
    const data = protocol.Data{
        .data_type = protocol.DataType.keepalive,
        .body = "",
    };

    return data.encode(alloc);
}
