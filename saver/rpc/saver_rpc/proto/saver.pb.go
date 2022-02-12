// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: saver.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item string `protobuf:"bytes,1,opt,name=Item,proto3" json:"Item,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_saver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_saver_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

var File_saver_proto protoreflect.FileDescriptor

var file_saver_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x61, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x32, 0x31, 0x0a, 0x05, 0x73, 0x61, 0x76,
	0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_saver_proto_rawDescOnce sync.Once
	file_saver_proto_rawDescData = file_saver_proto_rawDesc
)

func file_saver_proto_rawDescGZIP() []byte {
	file_saver_proto_rawDescOnce.Do(func() {
		file_saver_proto_rawDescData = protoimpl.X.CompressGZIP(file_saver_proto_rawDescData)
	})
	return file_saver_proto_rawDescData
}

var file_saver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_saver_proto_goTypes = []interface{}{
	(*Request)(nil),       // 0: Request
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_saver_proto_depIdxs = []int32{
	0, // 0: saver.Save:input_type -> Request
	1, // 1: saver.Save:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_saver_proto_init() }
func file_saver_proto_init() {
	if File_saver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_saver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_saver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_saver_proto_goTypes,
		DependencyIndexes: file_saver_proto_depIdxs,
		MessageInfos:      file_saver_proto_msgTypes,
	}.Build()
	File_saver_proto = out.File
	file_saver_proto_rawDesc = nil
	file_saver_proto_goTypes = nil
	file_saver_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SaverClient is the client API for Saver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SaverClient interface {
	Save(ctx context.Context, in *Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type saverClient struct {
	cc grpc.ClientConnInterface
}

func NewSaverClient(cc grpc.ClientConnInterface) SaverClient {
	return &saverClient{cc}
}

func (c *saverClient) Save(ctx context.Context, in *Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/saver/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SaverServer is the server API for Saver service.
type SaverServer interface {
	Save(context.Context, *Request) (*emptypb.Empty, error)
}

// UnimplementedSaverServer can be embedded to have forward compatible implementations.
type UnimplementedSaverServer struct {
}

func (*UnimplementedSaverServer) Save(context.Context, *Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}

func RegisterSaverServer(s *grpc.Server, srv SaverServer) {
	s.RegisterService(&_Saver_serviceDesc, srv)
}

func _Saver_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaverServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saver/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaverServer).Save(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Saver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "saver",
	HandlerType: (*SaverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Saver_Save_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "saver.proto",
}
