// nolint
// The file is auto-generated, to test the auto-generated grpc go file
// Google tests needed here
package grpc

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/runtime/protoimpl"
	"testing"
)

func TestIPStringRequest_Descriptor(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ip            string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
		want1  []int
	}{
		{"Autogenerated test", fields{}, []byte{0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x1, 0x90, 0x0, 0x6f, 0xff, 0xa, 0xa, 0x32, 0x62, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0xa, 0xf, 0x49, 0x50, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0xe, 0xa, 0x2, 0x69, 0x70, 0x18, 0x1, 0x20, 0x1, 0x28, 0x9, 0x52, 0x2, 0x69, 0x70, 0x22, 0x19, 0xa, 0x7, 0x4f, 0x4b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0xe, 0xa, 0x2, 0x6f, 0x6b, 0x18, 0x1, 0x20, 0x1, 0x28, 0x8, 0x52, 0x2, 0x6f, 0x6b, 0x32, 0x2c, 0xa, 0x6, 0x49, 0x50, 0x32, 0x62, 0x61, 0x6e, 0x12, 0x22, 0xa, 0x2, 0x49, 0x50, 0x12, 0x10, 0x2e, 0x49, 0x50, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x8, 0x2e, 0x4f, 0x4b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0, 0x42, 0xe, 0x5a, 0xc, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x46, 0x6f, 0x72, 0x67, 0x52, 0x50, 0x43, 0x62, 0x6, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x8b, 0x19, 0x32, 0x4c, 0x90, 0x0, 0x0, 0x0}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := &IPStringRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Ip:            tt.fields.Ip,
			}
			got, got1 := ip.Descriptor()
			assert.Equalf(t, tt.want, got, "Descriptor()")
			assert.Equalf(t, tt.want1, got1, "Descriptor()")
		})
	}
}

func TestIPStringRequest_GetIp(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ip            string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Autogenerated test", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &IPStringRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Ip:            tt.fields.Ip,
			}
			assert.Equalf(t, tt.want, x.GetIp(), "GetIp()")
		})
	}
}

func TestIPStringRequest_ProtoMessage(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ip            string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Autogenerated test", fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := &IPStringRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Ip:            tt.fields.Ip,
			}
			ip.ProtoMessage()
		})
	}
}

func TestIPStringRequest_ProtoReflect(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ip            string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Autogenerated test", fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Google programmer needed
		})
	}
}

func TestIPStringRequest_Reset(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ip            string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Autogenerated test", fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &IPStringRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Ip:            tt.fields.Ip,
			}
			x.Reset()
		})
	}
}

func TestIPStringRequest_String(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ip            string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Autogenerated test", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &IPStringRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Ip:            tt.fields.Ip,
			}
			assert.Equalf(t, tt.want, x.String(), "String()")
		})
	}
}

func TestOKReply_Descriptor(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ok            bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
		want1  []int
	}{
		{"Autogenerated test", fields{}, []byte{0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x1, 0x90, 0x0, 0x6f, 0xff, 0xa, 0xa, 0x32, 0x62, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0xa, 0xf, 0x49, 0x50, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0xe, 0xa, 0x2, 0x69, 0x70, 0x18, 0x1, 0x20, 0x1, 0x28, 0x9, 0x52, 0x2, 0x69, 0x70, 0x22, 0x19, 0xa, 0x7, 0x4f, 0x4b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0xe, 0xa, 0x2, 0x6f, 0x6b, 0x18, 0x1, 0x20, 0x1, 0x28, 0x8, 0x52, 0x2, 0x6f, 0x6b, 0x32, 0x2c, 0xa, 0x6, 0x49, 0x50, 0x32, 0x62, 0x61, 0x6e, 0x12, 0x22, 0xa, 0x2, 0x49, 0x50, 0x12, 0x10, 0x2e, 0x49, 0x50, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x8, 0x2e, 0x4f, 0x4b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0, 0x42, 0xe, 0x5a, 0xc, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x46, 0x6f, 0x72, 0x67, 0x52, 0x50, 0x43, 0x62, 0x6, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x8b, 0x19, 0x32, 0x4c, 0x90, 0x0, 0x0, 0x0}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := &OKReply{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Ok:            tt.fields.Ok,
			}
			got, got1 := ok.Descriptor()
			assert.Equalf(t, tt.want, got, "Descriptor()")
			assert.Equalf(t, tt.want1, got1, "Descriptor()")
		})
	}
}

func TestOKReply_GetOk(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ok            bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Autogenerated test", fields{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &OKReply{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Ok:            tt.fields.Ok,
			}
			assert.Equalf(t, tt.want, x.GetOk(), "GetOk()")
		})
	}
}

func TestOKReply_ProtoMessage(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ok            bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Autogenerated test", fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := &OKReply{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Ok:            tt.fields.Ok,
			}
			ok.ProtoMessage()
		})
	}
}

func TestOKReply_ProtoReflect(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ok            bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Autogenerated test", fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Google programmer needed
		})
	}
}

func TestOKReply_Reset(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ok            bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Autogenerated test", fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &OKReply{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Ok:            tt.fields.Ok,
			}
			x.Reset()
		})
	}
}

func TestOKReply_String(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Ok            bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Autogenerated test", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &OKReply{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Ok:            tt.fields.Ok,
			}
			assert.Equalf(t, tt.want, x.String(), "String()")
		})
	}
}

func Test_file__2ban_proto_init(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Autogenerated test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file__2ban_proto_init()
		})
	}
}

func Test_file__2ban_proto_rawDescGZIP(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{"Autogenerated test", []byte{0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0x1, 0x90, 0x0, 0x6f, 0xff, 0xa, 0xa, 0x32, 0x62, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0xa, 0xf, 0x49, 0x50, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0xe, 0xa, 0x2, 0x69, 0x70, 0x18, 0x1, 0x20, 0x1, 0x28, 0x9, 0x52, 0x2, 0x69, 0x70, 0x22, 0x19, 0xa, 0x7, 0x4f, 0x4b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0xe, 0xa, 0x2, 0x6f, 0x6b, 0x18, 0x1, 0x20, 0x1, 0x28, 0x8, 0x52, 0x2, 0x6f, 0x6b, 0x32, 0x2c, 0xa, 0x6, 0x49, 0x50, 0x32, 0x62, 0x61, 0x6e, 0x12, 0x22, 0xa, 0x2, 0x49, 0x50, 0x12, 0x10, 0x2e, 0x49, 0x50, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x8, 0x2e, 0x4f, 0x4b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0, 0x42, 0xe, 0x5a, 0xc, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x46, 0x6f, 0x72, 0x67, 0x52, 0x50, 0x43, 0x62, 0x6, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x8b, 0x19, 0x32, 0x4c, 0x90, 0x0, 0x0, 0x0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, file__2ban_proto_rawDescGZIP(), "file__2ban_proto_rawDescGZIP()")
		})
	}
}
