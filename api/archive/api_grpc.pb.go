// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package archive

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

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	GetFirst(ctx context.Context, in *GetFirstRequest, opts ...grpc.CallOption) (*GetFirstResponse, error)
	GetLast(ctx context.Context, in *GetLastRequest, opts ...grpc.CallOption) (*GetLastResponse, error)
	GetHeightForBlock(ctx context.Context, in *GetHeightForBlockRequest, opts ...grpc.CallOption) (*GetHeightForBlockResponse, error)
	GetCommit(ctx context.Context, in *GetCommitRequest, opts ...grpc.CallOption) (*GetCommitResponse, error)
	GetHeader(ctx context.Context, in *GetHeaderRequest, opts ...grpc.CallOption) (*GetHeaderResponse, error)
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
	GetRegisterValues(ctx context.Context, in *GetRegisterValuesRequest, opts ...grpc.CallOption) (*GetRegisterValuesResponse, error)
	GetCollection(ctx context.Context, in *GetCollectionRequest, opts ...grpc.CallOption) (*GetCollectionResponse, error)
	ListCollectionsForHeight(ctx context.Context, in *ListCollectionsForHeightRequest, opts ...grpc.CallOption) (*ListCollectionsForHeightResponse, error)
	GetGuarantee(ctx context.Context, in *GetGuaranteeRequest, opts ...grpc.CallOption) (*GetGuaranteeResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
	GetHeightForTransaction(ctx context.Context, in *GetHeightForTransactionRequest, opts ...grpc.CallOption) (*GetHeightForTransactionResponse, error)
	ListTransactionsForHeight(ctx context.Context, in *ListTransactionsForHeightRequest, opts ...grpc.CallOption) (*ListTransactionsForHeightResponse, error)
	GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error)
	GetSeal(ctx context.Context, in *GetSealRequest, opts ...grpc.CallOption) (*GetSealResponse, error)
	ListSealsForHeight(ctx context.Context, in *ListSealsForHeightRequest, opts ...grpc.CallOption) (*ListSealsForHeightResponse, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetFirst(ctx context.Context, in *GetFirstRequest, opts ...grpc.CallOption) (*GetFirstResponse, error) {
	out := new(GetFirstResponse)
	err := c.cc.Invoke(ctx, "/API/GetFirst", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetLast(ctx context.Context, in *GetLastRequest, opts ...grpc.CallOption) (*GetLastResponse, error) {
	out := new(GetLastResponse)
	err := c.cc.Invoke(ctx, "/API/GetLast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetHeightForBlock(ctx context.Context, in *GetHeightForBlockRequest, opts ...grpc.CallOption) (*GetHeightForBlockResponse, error) {
	out := new(GetHeightForBlockResponse)
	err := c.cc.Invoke(ctx, "/API/GetHeightForBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetCommit(ctx context.Context, in *GetCommitRequest, opts ...grpc.CallOption) (*GetCommitResponse, error) {
	out := new(GetCommitResponse)
	err := c.cc.Invoke(ctx, "/API/GetCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetHeader(ctx context.Context, in *GetHeaderRequest, opts ...grpc.CallOption) (*GetHeaderResponse, error) {
	out := new(GetHeaderResponse)
	err := c.cc.Invoke(ctx, "/API/GetHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, "/API/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetRegisterValues(ctx context.Context, in *GetRegisterValuesRequest, opts ...grpc.CallOption) (*GetRegisterValuesResponse, error) {
	out := new(GetRegisterValuesResponse)
	err := c.cc.Invoke(ctx, "/API/GetRegisterValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetCollection(ctx context.Context, in *GetCollectionRequest, opts ...grpc.CallOption) (*GetCollectionResponse, error) {
	out := new(GetCollectionResponse)
	err := c.cc.Invoke(ctx, "/API/GetCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListCollectionsForHeight(ctx context.Context, in *ListCollectionsForHeightRequest, opts ...grpc.CallOption) (*ListCollectionsForHeightResponse, error) {
	out := new(ListCollectionsForHeightResponse)
	err := c.cc.Invoke(ctx, "/API/ListCollectionsForHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetGuarantee(ctx context.Context, in *GetGuaranteeRequest, opts ...grpc.CallOption) (*GetGuaranteeResponse, error) {
	out := new(GetGuaranteeResponse)
	err := c.cc.Invoke(ctx, "/API/GetGuarantee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, "/API/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetHeightForTransaction(ctx context.Context, in *GetHeightForTransactionRequest, opts ...grpc.CallOption) (*GetHeightForTransactionResponse, error) {
	out := new(GetHeightForTransactionResponse)
	err := c.cc.Invoke(ctx, "/API/GetHeightForTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListTransactionsForHeight(ctx context.Context, in *ListTransactionsForHeightRequest, opts ...grpc.CallOption) (*ListTransactionsForHeightResponse, error) {
	out := new(ListTransactionsForHeightResponse)
	err := c.cc.Invoke(ctx, "/API/ListTransactionsForHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error) {
	out := new(GetResultResponse)
	err := c.cc.Invoke(ctx, "/API/GetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetSeal(ctx context.Context, in *GetSealRequest, opts ...grpc.CallOption) (*GetSealResponse, error) {
	out := new(GetSealResponse)
	err := c.cc.Invoke(ctx, "/API/GetSeal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListSealsForHeight(ctx context.Context, in *ListSealsForHeightRequest, opts ...grpc.CallOption) (*ListSealsForHeightResponse, error) {
	out := new(ListSealsForHeightResponse)
	err := c.cc.Invoke(ctx, "/API/ListSealsForHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations should embed UnimplementedAPIServer
// for forward compatibility
type APIServer interface {
	GetFirst(context.Context, *GetFirstRequest) (*GetFirstResponse, error)
	GetLast(context.Context, *GetLastRequest) (*GetLastResponse, error)
	GetHeightForBlock(context.Context, *GetHeightForBlockRequest) (*GetHeightForBlockResponse, error)
	GetCommit(context.Context, *GetCommitRequest) (*GetCommitResponse, error)
	GetHeader(context.Context, *GetHeaderRequest) (*GetHeaderResponse, error)
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
	GetRegisterValues(context.Context, *GetRegisterValuesRequest) (*GetRegisterValuesResponse, error)
	GetCollection(context.Context, *GetCollectionRequest) (*GetCollectionResponse, error)
	ListCollectionsForHeight(context.Context, *ListCollectionsForHeightRequest) (*ListCollectionsForHeightResponse, error)
	GetGuarantee(context.Context, *GetGuaranteeRequest) (*GetGuaranteeResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	GetHeightForTransaction(context.Context, *GetHeightForTransactionRequest) (*GetHeightForTransactionResponse, error)
	ListTransactionsForHeight(context.Context, *ListTransactionsForHeightRequest) (*ListTransactionsForHeightResponse, error)
	GetResult(context.Context, *GetResultRequest) (*GetResultResponse, error)
	GetSeal(context.Context, *GetSealRequest) (*GetSealResponse, error)
	ListSealsForHeight(context.Context, *ListSealsForHeightRequest) (*ListSealsForHeightResponse, error)
}

// UnimplementedAPIServer should be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (UnimplementedAPIServer) GetFirst(context.Context, *GetFirstRequest) (*GetFirstResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirst not implemented")
}
func (UnimplementedAPIServer) GetLast(context.Context, *GetLastRequest) (*GetLastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLast not implemented")
}
func (UnimplementedAPIServer) GetHeightForBlock(context.Context, *GetHeightForBlockRequest) (*GetHeightForBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeightForBlock not implemented")
}
func (UnimplementedAPIServer) GetCommit(context.Context, *GetCommitRequest) (*GetCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommit not implemented")
}
func (UnimplementedAPIServer) GetHeader(context.Context, *GetHeaderRequest) (*GetHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeader not implemented")
}
func (UnimplementedAPIServer) GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedAPIServer) GetRegisterValues(context.Context, *GetRegisterValuesRequest) (*GetRegisterValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisterValues not implemented")
}
func (UnimplementedAPIServer) GetCollection(context.Context, *GetCollectionRequest) (*GetCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollection not implemented")
}
func (UnimplementedAPIServer) ListCollectionsForHeight(context.Context, *ListCollectionsForHeightRequest) (*ListCollectionsForHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollectionsForHeight not implemented")
}
func (UnimplementedAPIServer) GetGuarantee(context.Context, *GetGuaranteeRequest) (*GetGuaranteeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuarantee not implemented")
}
func (UnimplementedAPIServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedAPIServer) GetHeightForTransaction(context.Context, *GetHeightForTransactionRequest) (*GetHeightForTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeightForTransaction not implemented")
}
func (UnimplementedAPIServer) ListTransactionsForHeight(context.Context, *ListTransactionsForHeightRequest) (*ListTransactionsForHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactionsForHeight not implemented")
}
func (UnimplementedAPIServer) GetResult(context.Context, *GetResultRequest) (*GetResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedAPIServer) GetSeal(context.Context, *GetSealRequest) (*GetSealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeal not implemented")
}
func (UnimplementedAPIServer) ListSealsForHeight(context.Context, *ListSealsForHeightRequest) (*ListSealsForHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSealsForHeight not implemented")
}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_GetFirst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFirstRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetFirst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetFirst",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetFirst(ctx, req.(*GetFirstRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetLast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetLast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetLast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetLast(ctx, req.(*GetLastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetHeightForBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeightForBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetHeightForBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetHeightForBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetHeightForBlock(ctx, req.(*GetHeightForBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetCommit(ctx, req.(*GetCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetHeader(ctx, req.(*GetHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetRegisterValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisterValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetRegisterValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetRegisterValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetRegisterValues(ctx, req.(*GetRegisterValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetCollection(ctx, req.(*GetCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListCollectionsForHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionsForHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListCollectionsForHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/ListCollectionsForHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListCollectionsForHeight(ctx, req.(*ListCollectionsForHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetGuarantee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGuaranteeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetGuarantee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetGuarantee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetGuarantee(ctx, req.(*GetGuaranteeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetHeightForTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeightForTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetHeightForTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetHeightForTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetHeightForTransaction(ctx, req.(*GetHeightForTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListTransactionsForHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsForHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListTransactionsForHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/ListTransactionsForHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListTransactionsForHeight(ctx, req.(*ListTransactionsForHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetResult(ctx, req.(*GetResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetSeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetSeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetSeal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetSeal(ctx, req.(*GetSealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListSealsForHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSealsForHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ListSealsForHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/ListSealsForHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ListSealsForHeight(ctx, req.(*ListSealsForHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFirst",
			Handler:    _API_GetFirst_Handler,
		},
		{
			MethodName: "GetLast",
			Handler:    _API_GetLast_Handler,
		},
		{
			MethodName: "GetHeightForBlock",
			Handler:    _API_GetHeightForBlock_Handler,
		},
		{
			MethodName: "GetCommit",
			Handler:    _API_GetCommit_Handler,
		},
		{
			MethodName: "GetHeader",
			Handler:    _API_GetHeader_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _API_GetEvents_Handler,
		},
		{
			MethodName: "GetRegisterValues",
			Handler:    _API_GetRegisterValues_Handler,
		},
		{
			MethodName: "GetCollection",
			Handler:    _API_GetCollection_Handler,
		},
		{
			MethodName: "ListCollectionsForHeight",
			Handler:    _API_ListCollectionsForHeight_Handler,
		},
		{
			MethodName: "GetGuarantee",
			Handler:    _API_GetGuarantee_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _API_GetTransaction_Handler,
		},
		{
			MethodName: "GetHeightForTransaction",
			Handler:    _API_GetHeightForTransaction_Handler,
		},
		{
			MethodName: "ListTransactionsForHeight",
			Handler:    _API_ListTransactionsForHeight_Handler,
		},
		{
			MethodName: "GetResult",
			Handler:    _API_GetResult_Handler,
		},
		{
			MethodName: "GetSeal",
			Handler:    _API_GetSeal_Handler,
		},
		{
			MethodName: "ListSealsForHeight",
			Handler:    _API_ListSealsForHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}