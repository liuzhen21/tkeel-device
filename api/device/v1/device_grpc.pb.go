// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DeviceClient is the client API for Device service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceClient interface {
	CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceResponse, error)
	UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*UpdateDeviceResponse, error)
	DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceResponse, error)
	GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error)
	SearchEntity(ctx context.Context, in *ListDeviceRequest, opts ...grpc.CallOption) (*ListDeviceResponse, error)
	AddDeviceExt(ctx context.Context, in *AddDeviceExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteDeviceExt(ctx context.Context, in *DeleteDeviceExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateDeviceExt(ctx context.Context, in *UpdateDeviceExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateDeviceDataRelation(ctx context.Context, in *CreateDeviceDataRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateDeviceDataRelation(ctx context.Context, in *UpdateDeviceDataRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteDeviceDataRelation(ctx context.Context, in *DeleteDeviceDataRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListDeviceDataRelation(ctx context.Context, in *ListDeviceDataRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetDeviceRaw(ctx context.Context, in *SetDeviceRawRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetDeviceAttribte(ctx context.Context, in *SetDeviceAttributeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetDeviceCommand(ctx context.Context, in *SetDeviceCommandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type deviceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceClient(cc grpc.ClientConnInterface) DeviceClient {
	return &deviceClient{cc}
}

func (c *deviceClient) CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceResponse, error) {
	out := new(CreateDeviceResponse)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/CreateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*UpdateDeviceResponse, error) {
	out := new(UpdateDeviceResponse)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/UpdateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceResponse, error) {
	out := new(DeleteDeviceResponse)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error) {
	out := new(GetDeviceResponse)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/GetDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) SearchEntity(ctx context.Context, in *ListDeviceRequest, opts ...grpc.CallOption) (*ListDeviceResponse, error) {
	out := new(ListDeviceResponse)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/SearchEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) AddDeviceExt(ctx context.Context, in *AddDeviceExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/AddDeviceExt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) DeleteDeviceExt(ctx context.Context, in *DeleteDeviceExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/DeleteDeviceExt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) UpdateDeviceExt(ctx context.Context, in *UpdateDeviceExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/UpdateDeviceExt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) CreateDeviceDataRelation(ctx context.Context, in *CreateDeviceDataRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/CreateDeviceDataRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) UpdateDeviceDataRelation(ctx context.Context, in *UpdateDeviceDataRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/UpdateDeviceDataRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) DeleteDeviceDataRelation(ctx context.Context, in *DeleteDeviceDataRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/DeleteDeviceDataRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) ListDeviceDataRelation(ctx context.Context, in *ListDeviceDataRelationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/ListDeviceDataRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) SetDeviceRaw(ctx context.Context, in *SetDeviceRawRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/SetDeviceRaw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) SetDeviceAttribte(ctx context.Context, in *SetDeviceAttributeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/SetDeviceAttribte", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) SetDeviceCommand(ctx context.Context, in *SetDeviceCommandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.device.v1.Device/SetDeviceCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServer is the server API for Device service.
// All implementations must embed UnimplementedDeviceServer
// for forward compatibility
type DeviceServer interface {
	CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceResponse, error)
	UpdateDevice(context.Context, *UpdateDeviceRequest) (*UpdateDeviceResponse, error)
	DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error)
	GetDevice(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error)
	SearchEntity(context.Context, *ListDeviceRequest) (*ListDeviceResponse, error)
	AddDeviceExt(context.Context, *AddDeviceExtRequest) (*emptypb.Empty, error)
	DeleteDeviceExt(context.Context, *DeleteDeviceExtRequest) (*emptypb.Empty, error)
	UpdateDeviceExt(context.Context, *UpdateDeviceExtRequest) (*emptypb.Empty, error)
	CreateDeviceDataRelation(context.Context, *CreateDeviceDataRelationRequest) (*emptypb.Empty, error)
	UpdateDeviceDataRelation(context.Context, *UpdateDeviceDataRelationRequest) (*emptypb.Empty, error)
	DeleteDeviceDataRelation(context.Context, *DeleteDeviceDataRelationRequest) (*emptypb.Empty, error)
	ListDeviceDataRelation(context.Context, *ListDeviceDataRelationRequest) (*emptypb.Empty, error)
	SetDeviceRaw(context.Context, *SetDeviceRawRequest) (*emptypb.Empty, error)
	SetDeviceAttribte(context.Context, *SetDeviceAttributeRequest) (*emptypb.Empty, error)
	SetDeviceCommand(context.Context, *SetDeviceCommandRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDeviceServer()
}

// UnimplementedDeviceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceServer struct {
}

func (UnimplementedDeviceServer) CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedDeviceServer) UpdateDevice(context.Context, *UpdateDeviceRequest) (*UpdateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevice not implemented")
}
func (UnimplementedDeviceServer) DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (UnimplementedDeviceServer) GetDevice(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevice not implemented")
}
func (UnimplementedDeviceServer) SearchEntity(context.Context, *ListDeviceRequest) (*ListDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchEntity not implemented")
}
func (UnimplementedDeviceServer) AddDeviceExt(context.Context, *AddDeviceExtRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDeviceExt not implemented")
}
func (UnimplementedDeviceServer) DeleteDeviceExt(context.Context, *DeleteDeviceExtRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeviceExt not implemented")
}
func (UnimplementedDeviceServer) UpdateDeviceExt(context.Context, *UpdateDeviceExtRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeviceExt not implemented")
}
func (UnimplementedDeviceServer) CreateDeviceDataRelation(context.Context, *CreateDeviceDataRelationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceDataRelation not implemented")
}
func (UnimplementedDeviceServer) UpdateDeviceDataRelation(context.Context, *UpdateDeviceDataRelationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeviceDataRelation not implemented")
}
func (UnimplementedDeviceServer) DeleteDeviceDataRelation(context.Context, *DeleteDeviceDataRelationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeviceDataRelation not implemented")
}
func (UnimplementedDeviceServer) ListDeviceDataRelation(context.Context, *ListDeviceDataRelationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeviceDataRelation not implemented")
}
func (UnimplementedDeviceServer) SetDeviceRaw(context.Context, *SetDeviceRawRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeviceRaw not implemented")
}
func (UnimplementedDeviceServer) SetDeviceAttribte(context.Context, *SetDeviceAttributeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeviceAttribte not implemented")
}
func (UnimplementedDeviceServer) SetDeviceCommand(context.Context, *SetDeviceCommandRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeviceCommand not implemented")
}
func (UnimplementedDeviceServer) mustEmbedUnimplementedDeviceServer() {}

// UnsafeDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServer will
// result in compilation errors.
type UnsafeDeviceServer interface {
	mustEmbedUnimplementedDeviceServer()
}

func RegisterDeviceServer(s grpc.ServiceRegistrar, srv DeviceServer) {
	s.RegisterService(&Device_ServiceDesc, srv)
}

func _Device_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).CreateDevice(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_UpdateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).UpdateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/UpdateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).UpdateDevice(ctx, req.(*UpdateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).DeleteDevice(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/GetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).GetDevice(ctx, req.(*GetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_SearchEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).SearchEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/SearchEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).SearchEntity(ctx, req.(*ListDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_AddDeviceExt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceExtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).AddDeviceExt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/AddDeviceExt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).AddDeviceExt(ctx, req.(*AddDeviceExtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_DeleteDeviceExt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceExtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).DeleteDeviceExt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/DeleteDeviceExt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).DeleteDeviceExt(ctx, req.(*DeleteDeviceExtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_UpdateDeviceExt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceExtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).UpdateDeviceExt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/UpdateDeviceExt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).UpdateDeviceExt(ctx, req.(*UpdateDeviceExtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_CreateDeviceDataRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceDataRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).CreateDeviceDataRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/CreateDeviceDataRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).CreateDeviceDataRelation(ctx, req.(*CreateDeviceDataRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_UpdateDeviceDataRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceDataRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).UpdateDeviceDataRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/UpdateDeviceDataRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).UpdateDeviceDataRelation(ctx, req.(*UpdateDeviceDataRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_DeleteDeviceDataRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceDataRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).DeleteDeviceDataRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/DeleteDeviceDataRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).DeleteDeviceDataRelation(ctx, req.(*DeleteDeviceDataRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_ListDeviceDataRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceDataRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).ListDeviceDataRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/ListDeviceDataRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).ListDeviceDataRelation(ctx, req.(*ListDeviceDataRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_SetDeviceRaw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceRawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).SetDeviceRaw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/SetDeviceRaw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).SetDeviceRaw(ctx, req.(*SetDeviceRawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_SetDeviceAttribte_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).SetDeviceAttribte(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/SetDeviceAttribte",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).SetDeviceAttribte(ctx, req.(*SetDeviceAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_SetDeviceCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).SetDeviceCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.device.v1.Device/SetDeviceCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).SetDeviceCommand(ctx, req.(*SetDeviceCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Device_ServiceDesc is the grpc.ServiceDesc for Device service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Device_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.device.v1.Device",
	HandlerType: (*DeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDevice",
			Handler:    _Device_CreateDevice_Handler,
		},
		{
			MethodName: "UpdateDevice",
			Handler:    _Device_UpdateDevice_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _Device_DeleteDevice_Handler,
		},
		{
			MethodName: "GetDevice",
			Handler:    _Device_GetDevice_Handler,
		},
		{
			MethodName: "SearchEntity",
			Handler:    _Device_SearchEntity_Handler,
		},
		{
			MethodName: "AddDeviceExt",
			Handler:    _Device_AddDeviceExt_Handler,
		},
		{
			MethodName: "DeleteDeviceExt",
			Handler:    _Device_DeleteDeviceExt_Handler,
		},
		{
			MethodName: "UpdateDeviceExt",
			Handler:    _Device_UpdateDeviceExt_Handler,
		},
		{
			MethodName: "CreateDeviceDataRelation",
			Handler:    _Device_CreateDeviceDataRelation_Handler,
		},
		{
			MethodName: "UpdateDeviceDataRelation",
			Handler:    _Device_UpdateDeviceDataRelation_Handler,
		},
		{
			MethodName: "DeleteDeviceDataRelation",
			Handler:    _Device_DeleteDeviceDataRelation_Handler,
		},
		{
			MethodName: "ListDeviceDataRelation",
			Handler:    _Device_ListDeviceDataRelation_Handler,
		},
		{
			MethodName: "SetDeviceRaw",
			Handler:    _Device_SetDeviceRaw_Handler,
		},
		{
			MethodName: "SetDeviceAttribte",
			Handler:    _Device_SetDeviceAttribte_Handler,
		},
		{
			MethodName: "SetDeviceCommand",
			Handler:    _Device_SetDeviceCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/device/v1/device.proto",
}
