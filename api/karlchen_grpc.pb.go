// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// DokoClient is the client API for Doko service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DokoClient interface {
	Register(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*RegisterReply, error)
	CheckLogin(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserName, error)
	CreateTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TableData, error)
	StartTable(ctx context.Context, in *TableId, opts ...grpc.CallOption) (*MatchState, error)
	JoinTable(ctx context.Context, in *JoinTableRequest, opts ...grpc.CallOption) (*TableState, error)
	PlayCard(ctx context.Context, in *PlayCardRequest, opts ...grpc.CallOption) (*PlayedCard, error)
	PlaceBid(ctx context.Context, in *PlaceBidRequest, opts ...grpc.CallOption) (*Bid, error)
	Declare(ctx context.Context, in *DeclareRequest, opts ...grpc.CallOption) (*Declaration, error)
	StartNextMatch(ctx context.Context, in *TableId, opts ...grpc.CallOption) (*MatchState, error)
	StartSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Doko_StartSessionClient, error)
}

type dokoClient struct {
	cc grpc.ClientConnInterface
}

func NewDokoClient(cc grpc.ClientConnInterface) DokoClient {
	return &dokoClient{cc}
}

func (c *dokoClient) Register(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/api.Doko/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokoClient) CheckLogin(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserName, error) {
	out := new(UserName)
	err := c.cc.Invoke(ctx, "/api.Doko/CheckLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokoClient) CreateTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TableData, error) {
	out := new(TableData)
	err := c.cc.Invoke(ctx, "/api.Doko/CreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokoClient) StartTable(ctx context.Context, in *TableId, opts ...grpc.CallOption) (*MatchState, error) {
	out := new(MatchState)
	err := c.cc.Invoke(ctx, "/api.Doko/StartTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokoClient) JoinTable(ctx context.Context, in *JoinTableRequest, opts ...grpc.CallOption) (*TableState, error) {
	out := new(TableState)
	err := c.cc.Invoke(ctx, "/api.Doko/JoinTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokoClient) PlayCard(ctx context.Context, in *PlayCardRequest, opts ...grpc.CallOption) (*PlayedCard, error) {
	out := new(PlayedCard)
	err := c.cc.Invoke(ctx, "/api.Doko/PlayCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokoClient) PlaceBid(ctx context.Context, in *PlaceBidRequest, opts ...grpc.CallOption) (*Bid, error) {
	out := new(Bid)
	err := c.cc.Invoke(ctx, "/api.Doko/PlaceBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokoClient) Declare(ctx context.Context, in *DeclareRequest, opts ...grpc.CallOption) (*Declaration, error) {
	out := new(Declaration)
	err := c.cc.Invoke(ctx, "/api.Doko/Declare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokoClient) StartNextMatch(ctx context.Context, in *TableId, opts ...grpc.CallOption) (*MatchState, error) {
	out := new(MatchState)
	err := c.cc.Invoke(ctx, "/api.Doko/StartNextMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dokoClient) StartSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Doko_StartSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Doko_ServiceDesc.Streams[0], "/api.Doko/StartSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &dokoStartSessionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Doko_StartSessionClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type dokoStartSessionClient struct {
	grpc.ClientStream
}

func (x *dokoStartSessionClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DokoServer is the server API for Doko service.
// All implementations must embed UnimplementedDokoServer
// for forward compatibility
type DokoServer interface {
	Register(context.Context, *UserName) (*RegisterReply, error)
	CheckLogin(context.Context, *Empty) (*UserName, error)
	CreateTable(context.Context, *Empty) (*TableData, error)
	StartTable(context.Context, *TableId) (*MatchState, error)
	JoinTable(context.Context, *JoinTableRequest) (*TableState, error)
	PlayCard(context.Context, *PlayCardRequest) (*PlayedCard, error)
	PlaceBid(context.Context, *PlaceBidRequest) (*Bid, error)
	Declare(context.Context, *DeclareRequest) (*Declaration, error)
	StartNextMatch(context.Context, *TableId) (*MatchState, error)
	StartSession(*Empty, Doko_StartSessionServer) error
	mustEmbedUnimplementedDokoServer()
}

// UnimplementedDokoServer must be embedded to have forward compatible implementations.
type UnimplementedDokoServer struct {
}

func (UnimplementedDokoServer) Register(context.Context, *UserName) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedDokoServer) CheckLogin(context.Context, *Empty) (*UserName, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckLogin not implemented")
}
func (UnimplementedDokoServer) CreateTable(context.Context, *Empty) (*TableData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}
func (UnimplementedDokoServer) StartTable(context.Context, *TableId) (*MatchState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartTable not implemented")
}
func (UnimplementedDokoServer) JoinTable(context.Context, *JoinTableRequest) (*TableState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinTable not implemented")
}
func (UnimplementedDokoServer) PlayCard(context.Context, *PlayCardRequest) (*PlayedCard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayCard not implemented")
}
func (UnimplementedDokoServer) PlaceBid(context.Context, *PlaceBidRequest) (*Bid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBid not implemented")
}
func (UnimplementedDokoServer) Declare(context.Context, *DeclareRequest) (*Declaration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Declare not implemented")
}
func (UnimplementedDokoServer) StartNextMatch(context.Context, *TableId) (*MatchState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartNextMatch not implemented")
}
func (UnimplementedDokoServer) StartSession(*Empty, Doko_StartSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method StartSession not implemented")
}
func (UnimplementedDokoServer) mustEmbedUnimplementedDokoServer() {}

// UnsafeDokoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DokoServer will
// result in compilation errors.
type UnsafeDokoServer interface {
	mustEmbedUnimplementedDokoServer()
}

func RegisterDokoServer(s grpc.ServiceRegistrar, srv DokoServer) {
	s.RegisterService(&Doko_ServiceDesc, srv)
}

func _Doko_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokoServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Doko/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokoServer).Register(ctx, req.(*UserName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doko_CheckLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokoServer).CheckLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Doko/CheckLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokoServer).CheckLogin(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doko_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokoServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Doko/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokoServer).CreateTable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doko_StartTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TableId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokoServer).StartTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Doko/StartTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokoServer).StartTable(ctx, req.(*TableId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doko_JoinTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokoServer).JoinTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Doko/JoinTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokoServer).JoinTable(ctx, req.(*JoinTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doko_PlayCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokoServer).PlayCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Doko/PlayCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokoServer).PlayCard(ctx, req.(*PlayCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doko_PlaceBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceBidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokoServer).PlaceBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Doko/PlaceBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokoServer).PlaceBid(ctx, req.(*PlaceBidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doko_Declare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokoServer).Declare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Doko/Declare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokoServer).Declare(ctx, req.(*DeclareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doko_StartNextMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TableId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DokoServer).StartNextMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Doko/StartNextMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DokoServer).StartNextMatch(ctx, req.(*TableId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doko_StartSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DokoServer).StartSession(m, &dokoStartSessionServer{stream})
}

type Doko_StartSessionServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type dokoStartSessionServer struct {
	grpc.ServerStream
}

func (x *dokoStartSessionServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// Doko_ServiceDesc is the grpc.ServiceDesc for Doko service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Doko_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Doko",
	HandlerType: (*DokoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Doko_Register_Handler,
		},
		{
			MethodName: "CheckLogin",
			Handler:    _Doko_CheckLogin_Handler,
		},
		{
			MethodName: "CreateTable",
			Handler:    _Doko_CreateTable_Handler,
		},
		{
			MethodName: "StartTable",
			Handler:    _Doko_StartTable_Handler,
		},
		{
			MethodName: "JoinTable",
			Handler:    _Doko_JoinTable_Handler,
		},
		{
			MethodName: "PlayCard",
			Handler:    _Doko_PlayCard_Handler,
		},
		{
			MethodName: "PlaceBid",
			Handler:    _Doko_PlaceBid_Handler,
		},
		{
			MethodName: "Declare",
			Handler:    _Doko_Declare_Handler,
		},
		{
			MethodName: "StartNextMatch",
			Handler:    _Doko_StartNextMatch_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartSession",
			Handler:       _Doko_StartSession_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/karlchen.proto",
}