// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: iam.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IAMService_Auth_FullMethodName                = "/basic.IAMService/Auth"
	IAMService_CheckAuthStatus_FullMethodName     = "/basic.IAMService/CheckAuthStatus"
	IAMService_ValidateAuthorities_FullMethodName = "/basic.IAMService/ValidateAuthorities"
	IAMService_CreateGroup_FullMethodName         = "/basic.IAMService/CreateGroup"
	IAMService_UpdateGroup_FullMethodName         = "/basic.IAMService/UpdateGroup"
	IAMService_GetGroup_FullMethodName            = "/basic.IAMService/GetGroup"
	IAMService_ListGroup_FullMethodName           = "/basic.IAMService/ListGroup"
	IAMService_CreateAuthority_FullMethodName     = "/basic.IAMService/CreateAuthority"
	IAMService_UpdateAuthority_FullMethodName     = "/basic.IAMService/UpdateAuthority"
	IAMService_GetAuthority_FullMethodName        = "/basic.IAMService/GetAuthority"
	IAMService_ListAuthority_FullMethodName       = "/basic.IAMService/ListAuthority"
	IAMService_CreateOrg_FullMethodName           = "/basic.IAMService/CreateOrg"
	IAMService_UpdateOrg_FullMethodName           = "/basic.IAMService/UpdateOrg"
	IAMService_GetOrg_FullMethodName              = "/basic.IAMService/GetOrg"
	IAMService_ListOrg_FullMethodName             = "/basic.IAMService/ListOrg"
	IAMService_CreateUser_FullMethodName          = "/basic.IAMService/CreateUser"
	IAMService_GetUser_FullMethodName             = "/basic.IAMService/GetUser"
	IAMService_UpdateUser_FullMethodName          = "/basic.IAMService/UpdateUser"
	IAMService_ListUser_FullMethodName            = "/basic.IAMService/ListUser"
	IAMService_CreateNamespace_FullMethodName     = "/basic.IAMService/CreateNamespace"
	IAMService_UpdateNamespace_FullMethodName     = "/basic.IAMService/UpdateNamespace"
	IAMService_GetNamespace_FullMethodName        = "/basic.IAMService/GetNamespace"
	IAMService_ListNamespace_FullMethodName       = "/basic.IAMService/ListNamespace"
)

// IAMServiceClient is the client API for IAMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IAMServiceClient interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	CheckAuthStatus(ctx context.Context, in *CheckAuthStatusRequest, opts ...grpc.CallOption) (*CheckAuthStatusResponse, error)
	ValidateAuthorities(ctx context.Context, in *ValidateAuthoritiesRequest, opts ...grpc.CallOption) (*ValidateAuthoritiesResponse, error)
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Group, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*Group, error)
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error)
	ListGroup(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupResponse, error)
	CreateAuthority(ctx context.Context, in *CreateAuthorityRequest, opts ...grpc.CallOption) (*Authority, error)
	UpdateAuthority(ctx context.Context, in *UpdateAuthorityRequest, opts ...grpc.CallOption) (*Authority, error)
	GetAuthority(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Authority, error)
	ListAuthority(ctx context.Context, in *ListAuthorityRequest, opts ...grpc.CallOption) (*ListAuthorityResponse, error)
	CreateOrg(ctx context.Context, in *CreateOrgRequest, opts ...grpc.CallOption) (*Org, error)
	UpdateOrg(ctx context.Context, in *UpdateOrgRequest, opts ...grpc.CallOption) (*Org, error)
	GetOrg(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Org, error)
	ListOrg(ctx context.Context, in *ListOrgRequest, opts ...grpc.CallOption) (*ListOrgResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error)
	UpdateNamespace(ctx context.Context, in *UpdateNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error)
	GetNamespace(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Namespace, error)
	ListNamespace(ctx context.Context, in *ListNamespaceRequest, opts ...grpc.CallOption) (*ListNamespaceResponse, error)
}

type iAMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIAMServiceClient(cc grpc.ClientConnInterface) IAMServiceClient {
	return &iAMServiceClient{cc}
}

func (c *iAMServiceClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, IAMService_Auth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) CheckAuthStatus(ctx context.Context, in *CheckAuthStatusRequest, opts ...grpc.CallOption) (*CheckAuthStatusResponse, error) {
	out := new(CheckAuthStatusResponse)
	err := c.cc.Invoke(ctx, IAMService_CheckAuthStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) ValidateAuthorities(ctx context.Context, in *ValidateAuthoritiesRequest, opts ...grpc.CallOption) (*ValidateAuthoritiesResponse, error) {
	out := new(ValidateAuthoritiesResponse)
	err := c.cc.Invoke(ctx, IAMService_ValidateAuthorities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, IAMService_CreateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, IAMService_UpdateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, IAMService_GetGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) ListGroup(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupResponse, error) {
	out := new(ListGroupResponse)
	err := c.cc.Invoke(ctx, IAMService_ListGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) CreateAuthority(ctx context.Context, in *CreateAuthorityRequest, opts ...grpc.CallOption) (*Authority, error) {
	out := new(Authority)
	err := c.cc.Invoke(ctx, IAMService_CreateAuthority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) UpdateAuthority(ctx context.Context, in *UpdateAuthorityRequest, opts ...grpc.CallOption) (*Authority, error) {
	out := new(Authority)
	err := c.cc.Invoke(ctx, IAMService_UpdateAuthority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) GetAuthority(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Authority, error) {
	out := new(Authority)
	err := c.cc.Invoke(ctx, IAMService_GetAuthority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) ListAuthority(ctx context.Context, in *ListAuthorityRequest, opts ...grpc.CallOption) (*ListAuthorityResponse, error) {
	out := new(ListAuthorityResponse)
	err := c.cc.Invoke(ctx, IAMService_ListAuthority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) CreateOrg(ctx context.Context, in *CreateOrgRequest, opts ...grpc.CallOption) (*Org, error) {
	out := new(Org)
	err := c.cc.Invoke(ctx, IAMService_CreateOrg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) UpdateOrg(ctx context.Context, in *UpdateOrgRequest, opts ...grpc.CallOption) (*Org, error) {
	out := new(Org)
	err := c.cc.Invoke(ctx, IAMService_UpdateOrg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) GetOrg(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Org, error) {
	out := new(Org)
	err := c.cc.Invoke(ctx, IAMService_GetOrg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) ListOrg(ctx context.Context, in *ListOrgRequest, opts ...grpc.CallOption) (*ListOrgResponse, error) {
	out := new(ListOrgResponse)
	err := c.cc.Invoke(ctx, IAMService_ListOrg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, IAMService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, IAMService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, IAMService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, IAMService_ListUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, IAMService_CreateNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) UpdateNamespace(ctx context.Context, in *UpdateNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, IAMService_UpdateNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) GetNamespace(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, IAMService_GetNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMServiceClient) ListNamespace(ctx context.Context, in *ListNamespaceRequest, opts ...grpc.CallOption) (*ListNamespaceResponse, error) {
	out := new(ListNamespaceResponse)
	err := c.cc.Invoke(ctx, IAMService_ListNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMServiceServer is the server API for IAMService service.
// All implementations must embed UnimplementedIAMServiceServer
// for forward compatibility
type IAMServiceServer interface {
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	CheckAuthStatus(context.Context, *CheckAuthStatusRequest) (*CheckAuthStatusResponse, error)
	ValidateAuthorities(context.Context, *ValidateAuthoritiesRequest) (*ValidateAuthoritiesResponse, error)
	CreateGroup(context.Context, *CreateGroupRequest) (*Group, error)
	UpdateGroup(context.Context, *UpdateGroupRequest) (*Group, error)
	GetGroup(context.Context, *GetGroupRequest) (*Group, error)
	ListGroup(context.Context, *ListGroupRequest) (*ListGroupResponse, error)
	CreateAuthority(context.Context, *CreateAuthorityRequest) (*Authority, error)
	UpdateAuthority(context.Context, *UpdateAuthorityRequest) (*Authority, error)
	GetAuthority(context.Context, *GetRequest) (*Authority, error)
	ListAuthority(context.Context, *ListAuthorityRequest) (*ListAuthorityResponse, error)
	CreateOrg(context.Context, *CreateOrgRequest) (*Org, error)
	UpdateOrg(context.Context, *UpdateOrgRequest) (*Org, error)
	GetOrg(context.Context, *GetRequest) (*Org, error)
	ListOrg(context.Context, *ListOrgRequest) (*ListOrgResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*Namespace, error)
	UpdateNamespace(context.Context, *UpdateNamespaceRequest) (*Namespace, error)
	GetNamespace(context.Context, *GetRequest) (*Namespace, error)
	ListNamespace(context.Context, *ListNamespaceRequest) (*ListNamespaceResponse, error)
	mustEmbedUnimplementedIAMServiceServer()
}

// UnimplementedIAMServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIAMServiceServer struct {
}

func (UnimplementedIAMServiceServer) Auth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedIAMServiceServer) CheckAuthStatus(context.Context, *CheckAuthStatusRequest) (*CheckAuthStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuthStatus not implemented")
}
func (UnimplementedIAMServiceServer) ValidateAuthorities(context.Context, *ValidateAuthoritiesRequest) (*ValidateAuthoritiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAuthorities not implemented")
}
func (UnimplementedIAMServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedIAMServiceServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedIAMServiceServer) GetGroup(context.Context, *GetGroupRequest) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedIAMServiceServer) ListGroup(context.Context, *ListGroupRequest) (*ListGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroup not implemented")
}
func (UnimplementedIAMServiceServer) CreateAuthority(context.Context, *CreateAuthorityRequest) (*Authority, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthority not implemented")
}
func (UnimplementedIAMServiceServer) UpdateAuthority(context.Context, *UpdateAuthorityRequest) (*Authority, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthority not implemented")
}
func (UnimplementedIAMServiceServer) GetAuthority(context.Context, *GetRequest) (*Authority, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthority not implemented")
}
func (UnimplementedIAMServiceServer) ListAuthority(context.Context, *ListAuthorityRequest) (*ListAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthority not implemented")
}
func (UnimplementedIAMServiceServer) CreateOrg(context.Context, *CreateOrgRequest) (*Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrg not implemented")
}
func (UnimplementedIAMServiceServer) UpdateOrg(context.Context, *UpdateOrgRequest) (*Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrg not implemented")
}
func (UnimplementedIAMServiceServer) GetOrg(context.Context, *GetRequest) (*Org, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrg not implemented")
}
func (UnimplementedIAMServiceServer) ListOrg(context.Context, *ListOrgRequest) (*ListOrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrg not implemented")
}
func (UnimplementedIAMServiceServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedIAMServiceServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedIAMServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedIAMServiceServer) ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedIAMServiceServer) CreateNamespace(context.Context, *CreateNamespaceRequest) (*Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamespace not implemented")
}
func (UnimplementedIAMServiceServer) UpdateNamespace(context.Context, *UpdateNamespaceRequest) (*Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNamespace not implemented")
}
func (UnimplementedIAMServiceServer) GetNamespace(context.Context, *GetRequest) (*Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespace not implemented")
}
func (UnimplementedIAMServiceServer) ListNamespace(context.Context, *ListNamespaceRequest) (*ListNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespace not implemented")
}
func (UnimplementedIAMServiceServer) mustEmbedUnimplementedIAMServiceServer() {}

// UnsafeIAMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IAMServiceServer will
// result in compilation errors.
type UnsafeIAMServiceServer interface {
	mustEmbedUnimplementedIAMServiceServer()
}

func RegisterIAMServiceServer(s grpc.ServiceRegistrar, srv IAMServiceServer) {
	s.RegisterService(&IAMService_ServiceDesc, srv)
}

func _IAMService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_CheckAuthStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).CheckAuthStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_CheckAuthStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).CheckAuthStatus(ctx, req.(*CheckAuthStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_ValidateAuthorities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateAuthoritiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).ValidateAuthorities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_ValidateAuthorities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).ValidateAuthorities(ctx, req.(*ValidateAuthoritiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_UpdateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_GetGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_ListGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).ListGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_ListGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).ListGroup(ctx, req.(*ListGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_CreateAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).CreateAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_CreateAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).CreateAuthority(ctx, req.(*CreateAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_UpdateAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).UpdateAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_UpdateAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).UpdateAuthority(ctx, req.(*UpdateAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_GetAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).GetAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_GetAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).GetAuthority(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_ListAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).ListAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_ListAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).ListAuthority(ctx, req.(*ListAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_CreateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).CreateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_CreateOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).CreateOrg(ctx, req.(*CreateOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_UpdateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).UpdateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_UpdateOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).UpdateOrg(ctx, req.(*UpdateOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_GetOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).GetOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_GetOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).GetOrg(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_ListOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).ListOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_ListOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).ListOrg(ctx, req.(*ListOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_CreateNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).CreateNamespace(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_UpdateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).UpdateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_UpdateNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).UpdateNamespace(ctx, req.(*UpdateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_GetNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).GetNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_GetNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).GetNamespace(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMService_ListNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMServiceServer).ListNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMService_ListNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMServiceServer).ListNamespace(ctx, req.(*ListNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IAMService_ServiceDesc is the grpc.ServiceDesc for IAMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IAMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "basic.IAMService",
	HandlerType: (*IAMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _IAMService_Auth_Handler,
		},
		{
			MethodName: "CheckAuthStatus",
			Handler:    _IAMService_CheckAuthStatus_Handler,
		},
		{
			MethodName: "ValidateAuthorities",
			Handler:    _IAMService_ValidateAuthorities_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _IAMService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _IAMService_UpdateGroup_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _IAMService_GetGroup_Handler,
		},
		{
			MethodName: "ListGroup",
			Handler:    _IAMService_ListGroup_Handler,
		},
		{
			MethodName: "CreateAuthority",
			Handler:    _IAMService_CreateAuthority_Handler,
		},
		{
			MethodName: "UpdateAuthority",
			Handler:    _IAMService_UpdateAuthority_Handler,
		},
		{
			MethodName: "GetAuthority",
			Handler:    _IAMService_GetAuthority_Handler,
		},
		{
			MethodName: "ListAuthority",
			Handler:    _IAMService_ListAuthority_Handler,
		},
		{
			MethodName: "CreateOrg",
			Handler:    _IAMService_CreateOrg_Handler,
		},
		{
			MethodName: "UpdateOrg",
			Handler:    _IAMService_UpdateOrg_Handler,
		},
		{
			MethodName: "GetOrg",
			Handler:    _IAMService_GetOrg_Handler,
		},
		{
			MethodName: "ListOrg",
			Handler:    _IAMService_ListOrg_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _IAMService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _IAMService_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _IAMService_UpdateUser_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _IAMService_ListUser_Handler,
		},
		{
			MethodName: "CreateNamespace",
			Handler:    _IAMService_CreateNamespace_Handler,
		},
		{
			MethodName: "UpdateNamespace",
			Handler:    _IAMService_UpdateNamespace_Handler,
		},
		{
			MethodName: "GetNamespace",
			Handler:    _IAMService_GetNamespace_Handler,
		},
		{
			MethodName: "ListNamespace",
			Handler:    _IAMService_ListNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iam.proto",
}
