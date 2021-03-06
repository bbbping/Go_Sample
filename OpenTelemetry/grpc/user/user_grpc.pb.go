// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: user_grpc.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InsertReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *InsertReq) Reset() {
	*x = InsertReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertReq) ProtoMessage() {}

func (x *InsertReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertReq.ProtoReflect.Descriptor instead.
func (*InsertReq) Descriptor() ([]byte, []int) {
	return file_user_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *InsertReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type InsertRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *InsertRes) Reset() {
	*x = InsertRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRes) ProtoMessage() {}

func (x *InsertRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertRes.ProtoReflect.Descriptor instead.
func (*InsertRes) Descriptor() ([]byte, []int) {
	return file_user_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *InsertRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QueryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *QueryReq) Reset() {
	*x = QueryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReq) ProtoMessage() {}

func (x *QueryReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReq.ProtoReflect.Descriptor instead.
func (*QueryReq) Descriptor() ([]byte, []int) {
	return file_user_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *QueryReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QueryRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *QueryRes) Reset() {
	*x = QueryRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRes) ProtoMessage() {}

func (x *QueryRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRes.ProtoReflect.Descriptor instead.
func (*QueryRes) Descriptor() ([]byte, []int) {
	return file_user_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *QueryRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QueryRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_user_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRes) Reset() {
	*x = DeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRes) ProtoMessage() {}

func (x *DeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRes.ProtoReflect.Descriptor instead.
func (*DeleteRes) Descriptor() ([]byte, []int) {
	return file_user_grpc_proto_rawDescGZIP(), []int{5}
}

var File_user_grpc_proto protoreflect.FileDescriptor

var file_user_grpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49,
	0x64, 0x22, 0x2e, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x1b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x0b,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x32, 0x8d, 0x01, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x0f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x29, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e,
	0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_grpc_proto_rawDescOnce sync.Once
	file_user_grpc_proto_rawDescData = file_user_grpc_proto_rawDesc
)

func file_user_grpc_proto_rawDescGZIP() []byte {
	file_user_grpc_proto_rawDescOnce.Do(func() {
		file_user_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_grpc_proto_rawDescData)
	})
	return file_user_grpc_proto_rawDescData
}

var file_user_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_grpc_proto_goTypes = []interface{}{
	(*InsertReq)(nil), // 0: user.InsertReq
	(*InsertRes)(nil), // 1: user.InsertRes
	(*QueryReq)(nil),  // 2: user.QueryReq
	(*QueryRes)(nil),  // 3: user.QueryRes
	(*DeleteReq)(nil), // 4: user.DeleteReq
	(*DeleteRes)(nil), // 5: user.DeleteRes
}
var file_user_grpc_proto_depIdxs = []int32{
	0, // 0: user.User.Insert:input_type -> user.InsertReq
	2, // 1: user.User.Query:input_type -> user.QueryReq
	4, // 2: user.User.Delete:input_type -> user.DeleteReq
	1, // 3: user.User.Insert:output_type -> user.InsertRes
	3, // 4: user.User.Query:output_type -> user.QueryRes
	5, // 5: user.User.Delete:output_type -> user.DeleteRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_grpc_proto_init() }
func file_user_grpc_proto_init() {
	if File_user_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertReq); i {
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
		file_user_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertRes); i {
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
		file_user_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReq); i {
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
		file_user_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRes); i {
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
		file_user_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_user_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRes); i {
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
			RawDescriptor: file_user_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_grpc_proto_goTypes,
		DependencyIndexes: file_user_grpc_proto_depIdxs,
		MessageInfos:      file_user_grpc_proto_msgTypes,
	}.Build()
	File_user_grpc_proto = out.File
	file_user_grpc_proto_rawDesc = nil
	file_user_grpc_proto_goTypes = nil
	file_user_grpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Insert(ctx context.Context, in *InsertReq, opts ...grpc.CallOption) (*InsertRes, error)
	Query(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryRes, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Insert(ctx context.Context, in *InsertReq, opts ...grpc.CallOption) (*InsertRes, error) {
	out := new(InsertRes)
	err := c.cc.Invoke(ctx, "/user.User/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Query(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryRes, error) {
	out := new(QueryRes)
	err := c.cc.Invoke(ctx, "/user.User/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error) {
	out := new(DeleteRes)
	err := c.cc.Invoke(ctx, "/user.User/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Insert(context.Context, *InsertReq) (*InsertRes, error)
	Query(context.Context, *QueryReq) (*QueryRes, error)
	Delete(context.Context, *DeleteReq) (*DeleteRes, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Insert(context.Context, *InsertReq) (*InsertRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (*UnimplementedUserServer) Query(context.Context, *QueryReq) (*QueryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedUserServer) Delete(context.Context, *DeleteReq) (*DeleteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Insert(ctx, req.(*InsertReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Query(ctx, req.(*QueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _User_Insert_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _User_Query_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _User_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_grpc.proto",
}
