// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/user_services.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import context "context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserRegistry service

type UserRegistryClient interface {
	// Register a new user. This method may be restricted by network settings.
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Get the user with the given identifiers, selecting the fields given by the
	// field mask. The method may return more or less fields, depending on the rights
	// of the caller.
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Create a temporary password that can be used for updating a forgotten password.
	// The generated password is sent to the user's email address.
	CreateTemporaryPassword(ctx context.Context, in *CreateTemporaryPasswordRequest, opts ...grpc.CallOption) (*types.Empty, error)
	UpdatePassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Delete(ctx context.Context, in *UserIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type userRegistryClient struct {
	cc *grpc.ClientConn
}

func NewUserRegistryClient(cc *grpc.ClientConn) UserRegistryClient {
	return &userRegistryClient{cc}
}

func (c *userRegistryClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) CreateTemporaryPassword(ctx context.Context, in *CreateTemporaryPasswordRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/CreateTemporaryPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) UpdatePassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) Delete(ctx context.Context, in *UserIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserRegistry service

type UserRegistryServer interface {
	// Register a new user. This method may be restricted by network settings.
	Create(context.Context, *CreateUserRequest) (*User, error)
	// Get the user with the given identifiers, selecting the fields given by the
	// field mask. The method may return more or less fields, depending on the rights
	// of the caller.
	Get(context.Context, *GetUserRequest) (*User, error)
	Update(context.Context, *UpdateUserRequest) (*User, error)
	// Create a temporary password that can be used for updating a forgotten password.
	// The generated password is sent to the user's email address.
	CreateTemporaryPassword(context.Context, *CreateTemporaryPasswordRequest) (*types.Empty, error)
	UpdatePassword(context.Context, *UpdateUserPasswordRequest) (*types.Empty, error)
	Delete(context.Context, *UserIdentifiers) (*types.Empty, error)
}

func RegisterUserRegistryServer(s *grpc.Server, srv UserRegistryServer) {
	s.RegisterService(&_UserRegistry_serviceDesc, srv)
}

func _UserRegistry_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).Get(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_CreateTemporaryPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemporaryPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).CreateTemporaryPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/CreateTemporaryPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).CreateTemporaryPassword(ctx, req.(*CreateTemporaryPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).UpdatePassword(ctx, req.(*UpdateUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).Delete(ctx, req.(*UserIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.UserRegistry",
	HandlerType: (*UserRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserRegistry_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserRegistry_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserRegistry_Update_Handler,
		},
		{
			MethodName: "CreateTemporaryPassword",
			Handler:    _UserRegistry_CreateTemporaryPassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserRegistry_UpdatePassword_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/user_services.proto",
}

// Client API for UserAccess service

type UserAccessClient interface {
	ListRights(ctx context.Context, in *UserIdentifiers, opts ...grpc.CallOption) (*Rights, error)
	CreateAPIKey(ctx context.Context, in *CreateUserAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
	ListAPIKeys(ctx context.Context, in *UserIdentifiers, opts ...grpc.CallOption) (*APIKeys, error)
	// Update the rights of an existing user API key. To generate an API key,
	// the CreateAPIKey should be used. To delete an API key, update it
	// with zero rights.
	UpdateAPIKey(ctx context.Context, in *UpdateUserAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error)
}

type userAccessClient struct {
	cc *grpc.ClientConn
}

func NewUserAccessClient(cc *grpc.ClientConn) UserAccessClient {
	return &userAccessClient{cc}
}

func (c *userAccessClient) ListRights(ctx context.Context, in *UserIdentifiers, opts ...grpc.CallOption) (*Rights, error) {
	out := new(Rights)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserAccess/ListRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccessClient) CreateAPIKey(ctx context.Context, in *CreateUserAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserAccess/CreateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccessClient) ListAPIKeys(ctx context.Context, in *UserIdentifiers, opts ...grpc.CallOption) (*APIKeys, error) {
	out := new(APIKeys)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserAccess/ListAPIKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccessClient) UpdateAPIKey(ctx context.Context, in *UpdateUserAPIKeyRequest, opts ...grpc.CallOption) (*APIKey, error) {
	out := new(APIKey)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserAccess/UpdateAPIKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserAccess service

type UserAccessServer interface {
	ListRights(context.Context, *UserIdentifiers) (*Rights, error)
	CreateAPIKey(context.Context, *CreateUserAPIKeyRequest) (*APIKey, error)
	ListAPIKeys(context.Context, *UserIdentifiers) (*APIKeys, error)
	// Update the rights of an existing user API key. To generate an API key,
	// the CreateAPIKey should be used. To delete an API key, update it
	// with zero rights.
	UpdateAPIKey(context.Context, *UpdateUserAPIKeyRequest) (*APIKey, error)
}

func RegisterUserAccessServer(s *grpc.Server, srv UserAccessServer) {
	s.RegisterService(&_UserAccess_serviceDesc, srv)
}

func _UserAccess_ListRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessServer).ListRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserAccess/ListRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessServer).ListRights(ctx, req.(*UserIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccess_CreateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessServer).CreateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserAccess/CreateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessServer).CreateAPIKey(ctx, req.(*CreateUserAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccess_ListAPIKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessServer).ListAPIKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserAccess/ListAPIKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessServer).ListAPIKeys(ctx, req.(*UserIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccess_UpdateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessServer).UpdateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserAccess/UpdateAPIKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessServer).UpdateAPIKey(ctx, req.(*UpdateUserAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.UserAccess",
	HandlerType: (*UserAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRights",
			Handler:    _UserAccess_ListRights_Handler,
		},
		{
			MethodName: "CreateAPIKey",
			Handler:    _UserAccess_CreateAPIKey_Handler,
		},
		{
			MethodName: "ListAPIKeys",
			Handler:    _UserAccess_ListAPIKeys_Handler,
		},
		{
			MethodName: "UpdateAPIKey",
			Handler:    _UserAccess_UpdateAPIKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/user_services.proto",
}

// Client API for UserInvitationRegistry service

type UserInvitationRegistryClient interface {
	Send(ctx context.Context, in *SendInvitationRequest, opts ...grpc.CallOption) (*types.Empty, error)
	List(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*Invitations, error)
	Delete(ctx context.Context, in *DeleteInvitationRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type userInvitationRegistryClient struct {
	cc *grpc.ClientConn
}

func NewUserInvitationRegistryClient(cc *grpc.ClientConn) UserInvitationRegistryClient {
	return &userInvitationRegistryClient{cc}
}

func (c *userInvitationRegistryClient) Send(ctx context.Context, in *SendInvitationRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserInvitationRegistry/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInvitationRegistryClient) List(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*Invitations, error) {
	out := new(Invitations)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserInvitationRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInvitationRegistryClient) Delete(ctx context.Context, in *DeleteInvitationRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UserInvitationRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserInvitationRegistry service

type UserInvitationRegistryServer interface {
	Send(context.Context, *SendInvitationRequest) (*types.Empty, error)
	List(context.Context, *types.Empty) (*Invitations, error)
	Delete(context.Context, *DeleteInvitationRequest) (*types.Empty, error)
}

func RegisterUserInvitationRegistryServer(s *grpc.Server, srv UserInvitationRegistryServer) {
	s.RegisterService(&_UserInvitationRegistry_serviceDesc, srv)
}

func _UserInvitationRegistry_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInvitationRegistryServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserInvitationRegistry/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInvitationRegistryServer).Send(ctx, req.(*SendInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInvitationRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInvitationRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserInvitationRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInvitationRegistryServer).List(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInvitationRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInvitationRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UserInvitationRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInvitationRegistryServer).Delete(ctx, req.(*DeleteInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInvitationRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.UserInvitationRegistry",
	HandlerType: (*UserInvitationRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _UserInvitationRegistry_Send_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserInvitationRegistry_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserInvitationRegistry_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/user_services.proto",
}

func init() {
	proto.RegisterFile("lorawan-stack/api/user_services.proto", fileDescriptor_user_services_c96c7a28e66a0811)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/user_services.proto", fileDescriptor_user_services_c96c7a28e66a0811)
}

var fileDescriptor_user_services_c96c7a28e66a0811 = []byte{
	// 759 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x3f, 0x4c, 0xeb, 0x46,
	0x18, 0xf7, 0xbd, 0xf7, 0x9a, 0xe1, 0x5e, 0x84, 0xfa, 0x4e, 0x4f, 0x2f, 0x7d, 0xe6, 0xe9, 0x7b,
	0xc2, 0x2d, 0x42, 0x45, 0xe4, 0x2c, 0x01, 0x13, 0x1b, 0xfd, 0x23, 0x84, 0xda, 0x81, 0x52, 0x58,
	0x5a, 0x89, 0xc8, 0x49, 0x0e, 0xe7, 0x1a, 0xb0, 0x5d, 0xdf, 0x05, 0x14, 0x21, 0x2a, 0xc4, 0x84,
	0xd4, 0xa5, 0x52, 0x87, 0x76, 0xac, 0x3a, 0x51, 0xa9, 0x03, 0x23, 0x23, 0x23, 0x23, 0x52, 0x17,
	0x46, 0x62, 0x77, 0x60, 0xaa, 0x18, 0x19, 0x2b, 0xdf, 0xc5, 0x09, 0x71, 0x62, 0xc8, 0x66, 0xdf,
	0xf7, 0xdd, 0xef, 0xdf, 0xdd, 0x67, 0xe3, 0xe9, 0x1d, 0x3f, 0x74, 0xf6, 0x1d, 0xaf, 0x2c, 0xa4,
	0x53, 0x6b, 0xda, 0x4e, 0xc0, 0xed, 0x96, 0x60, 0x61, 0x45, 0xb0, 0x70, 0x8f, 0xd7, 0x98, 0xa0,
	0x41, 0xe8, 0x4b, 0x9f, 0x4c, 0x48, 0xe9, 0xd1, 0x6e, 0x2b, 0xdd, 0x5b, 0x30, 0xcb, 0x2e, 0x97,
	0x8d, 0x56, 0x95, 0xd6, 0xfc, 0x5d, 0xdb, 0xf5, 0x5d, 0xdf, 0x56, 0x6d, 0xd5, 0xd6, 0xb6, 0x7a,
	0x53, 0x2f, 0xea, 0x49, 0x6f, 0x37, 0xdf, 0xb9, 0xbe, 0xef, 0xee, 0x30, 0x05, 0xef, 0x78, 0x9e,
	0x2f, 0x1d, 0xc9, 0x7d, 0xaf, 0x0b, 0x6e, 0x4e, 0x76, 0xab, 0x3d, 0x0c, 0xb6, 0x1b, 0xc8, 0x76,
	0xb7, 0xf8, 0xf1, 0xb0, 0x40, 0x5e, 0x67, 0x9e, 0xe4, 0xdb, 0x9c, 0x85, 0x29, 0x02, 0x0c, 0x37,
	0x85, 0xdc, 0x6d, 0xc8, 0xb4, 0xfe, 0x6e, 0xb4, 0x4b, 0x5d, 0x9d, 0xff, 0xeb, 0x03, 0x5c, 0xdc,
	0x14, 0x2c, 0x5c, 0x67, 0x2e, 0x17, 0x32, 0x6c, 0x93, 0x0d, 0x5c, 0xf8, 0x3c, 0x64, 0x8e, 0x64,
	0x64, 0x8a, 0x0e, 0x1a, 0xa7, 0x7a, 0x5d, 0x77, 0xff, 0xd8, 0x62, 0x42, 0x9a, 0xaf, 0xb3, 0x2d,
	0x49, 0xd1, 0x7a, 0x75, 0xfc, 0xcf, 0xbf, 0xbf, 0x3e, 0x7b, 0x69, 0x15, 0x14, 0x91, 0x58, 0x42,
	0xb3, 0x64, 0x0b, 0x3f, 0x5f, 0x61, 0x92, 0x40, 0xb6, 0x7f, 0x85, 0xc9, 0xa7, 0xf1, 0xa6, 0x14,
	0xde, 0x24, 0x79, 0xab, 0xf1, 0xec, 0x03, 0x75, 0x4a, 0xbc, 0x2e, 0x68, 0xf7, 0xe1, 0x90, 0xb8,
	0xb8, 0xb0, 0x19, 0xd4, 0x47, 0xaa, 0xd6, 0xeb, 0x4f, 0xb3, 0x7c, 0xa2, 0x58, 0xc0, 0x1c, 0x60,
	0xa1, 0x0f, 0x59, 0x12, 0x23, 0xbf, 0x21, 0x5c, 0xd2, 0x39, 0x6c, 0xb0, 0xdd, 0xc0, 0x0f, 0x9d,
	0xb0, 0xbd, 0xe6, 0x08, 0xb1, 0xef, 0x87, 0x75, 0x42, 0x47, 0x07, 0x36, 0xd4, 0x98, 0xea, 0x78,
	0x43, 0xf5, 0xe1, 0xd3, 0xf4, 0xf0, 0xe9, 0x97, 0xc9, 0xe1, 0x5b, 0x8b, 0x4a, 0x09, 0xb5, 0xe6,
	0x72, 0xfd, 0xda, 0x32, 0xc5, 0xac, 0x04, 0x29, 0xfb, 0x31, 0xc2, 0x13, 0xda, 0x6b, 0x4f, 0xd0,
	0xa7, 0xf9, 0x59, 0x8c, 0xab, 0xa5, 0xac, 0xb4, 0xcc, 0x98, 0x56, 0xbe, 0x96, 0x54, 0x41, 0x12,
	0xcf, 0xf7, 0xb8, 0xf0, 0x05, 0xdb, 0x61, 0x92, 0x91, 0xf7, 0xa3, 0x42, 0x5e, 0xed, 0xdf, 0xde,
	0x5c, 0xc6, 0x8f, 0x14, 0x23, 0x99, 0xfd, 0x30, 0xc3, 0x78, 0x38, 0xff, 0xdf, 0x73, 0x8c, 0x13,
	0x94, 0xe5, 0x5a, 0x8d, 0x09, 0x41, 0xb6, 0x31, 0xfe, 0x9a, 0x0b, 0xb9, 0xae, 0x2e, 0xfb, 0x38,
	0x7c, 0x99, 0x06, 0xbd, 0xd1, 0x7a, 0xaf, 0xf8, 0xde, 0x92, 0x52, 0x96, 0xaf, 0x3b, 0x46, 0xe4,
	0x27, 0x5c, 0xd4, 0x07, 0xb9, 0xbc, 0xb6, 0xfa, 0x15, 0x6b, 0x93, 0x99, 0xfc, 0xb9, 0xd0, 0x1d,
	0xfd, 0x4c, 0x33, 0x8d, 0xba, 0x9c, 0x66, 0x6a, 0x3d, 0x92, 0xa9, 0x13, 0xf0, 0x72, 0x93, 0xb5,
	0xd5, 0xec, 0xfc, 0x80, 0x5f, 0x26, 0x3e, 0xf5, 0xe6, 0x31, 0x8c, 0x96, 0x46, 0xd3, 0x8a, 0xdc,
	0x39, 0xea, 0xd3, 0x91, 0x9f, 0x11, 0x2e, 0xea, 0x4b, 0x92, 0x67, 0xb6, 0x7f, 0x85, 0xc6, 0x33,
	0xbb, 0xa4, 0x48, 0x17, 0x4d, 0xfb, 0x69, 0xb3, 0xf6, 0x81, 0x13, 0xf0, 0x4a, 0x93, 0xb5, 0xa9,
	0x1e, 0xb6, 0xf9, 0xbf, 0x9f, 0xe1, 0x37, 0xca, 0x9d, 0xb7, 0xc7, 0xf5, 0x67, 0xb3, 0xf7, 0x99,
	0xda, 0xc2, 0x2f, 0xbe, 0x65, 0x5e, 0x9d, 0x4c, 0x67, 0x69, 0x93, 0xd5, 0x87, 0xfd, 0x8f, 0x5f,
	0xef, 0x92, 0x52, 0xf7, 0xca, 0x2a, 0xda, 0xbc, 0xb7, 0x47, 0x85, 0xfe, 0x0d, 0x7e, 0x91, 0x84,
	0x4e, 0x72, 0x36, 0x9a, 0x93, 0x59, 0xde, 0x3e, 0xa7, 0xb0, 0x5e, 0x2b, 0xd4, 0x09, 0x32, 0x80,
	0x4a, 0x2a, 0xbd, 0xd9, 0x18, 0x0a, 0x55, 0xaf, 0x8f, 0x2f, 0xbb, 0x4b, 0x30, 0x3b, 0x40, 0xf0,
	0xd9, 0x9f, 0xe8, 0xb2, 0x03, 0xe8, 0xaa, 0x03, 0xe8, 0xba, 0x03, 0xc6, 0x4d, 0x07, 0x8c, 0xdb,
	0x0e, 0x18, 0x77, 0x1d, 0x30, 0xee, 0x3b, 0x80, 0x8e, 0x22, 0x40, 0x27, 0x11, 0x18, 0xa7, 0x11,
	0xa0, 0xb3, 0x08, 0x8c, 0xf3, 0x08, 0x8c, 0x8b, 0x08, 0x8c, 0xcb, 0x08, 0xd0, 0x55, 0x04, 0xe8,
	0x3a, 0x02, 0xe3, 0x26, 0x02, 0x74, 0x1b, 0x81, 0x71, 0x17, 0x01, 0xba, 0x8f, 0xc0, 0x38, 0x8a,
	0xc1, 0x38, 0x89, 0x01, 0xfd, 0x12, 0x83, 0xf1, 0x7b, 0x0c, 0xe8, 0x8f, 0x18, 0x8c, 0xd3, 0x18,
	0x8c, 0xb3, 0x18, 0xd0, 0x79, 0x0c, 0xe8, 0x22, 0x06, 0xf4, 0xdd, 0x9c, 0xeb, 0x53, 0xd9, 0x60,
	0xb2, 0xc1, 0x3d, 0x57, 0x50, 0x8f, 0xc9, 0x7d, 0x3f, 0x6c, 0xda, 0x83, 0x7f, 0x9d, 0xa0, 0xe9,
	0xda, 0x52, 0x7a, 0x41, 0xb5, 0x5a, 0x50, 0x56, 0x16, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xb4,
	0x4e, 0x68, 0x8c, 0x7d, 0x07, 0x00, 0x00,
}
