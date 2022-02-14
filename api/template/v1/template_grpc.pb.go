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

// TemplateClient is the client API for Template service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateClient interface {
	//Operation template entity
	CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error)
	UpdateTemplate(ctx context.Context, in *UpdateTemplateRequest, opts ...grpc.CallOption) (*UpdateTemplateResponse, error)
	DeleteTemplate(ctx context.Context, in *DeleteTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error)
	//operation template attribute
	AddTemplateAttribute(ctx context.Context, in *AddTemplateAttributeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateTemplateAttribute(ctx context.Context, in *UpdateTemplateAttributeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTemplateAttribute(ctx context.Context, in *DeleteTemplateAttributeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTemplateAttribute(ctx context.Context, in *GetTemplateAttributeRequest, opts ...grpc.CallOption) (*GetTemplateAttributeResponse, error)
	ListTemplateAttribute(ctx context.Context, in *ListTemplateAttributeRequest, opts ...grpc.CallOption) (*ListTemplateAttributeResponse, error)
	//operation template telemetry
	AddTemplateTelemetry(ctx context.Context, in *AddTemplateTelemetryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateTemplateTelemetry(ctx context.Context, in *UpdateTemplateTelemetryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTemplateTelemetry(ctx context.Context, in *DeleteTemplateTelemetryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTemplateTelemetry(ctx context.Context, in *GetTemplateTelemetryRequest, opts ...grpc.CallOption) (*GetTemplateTelemetryResponse, error)
	ListTemplateTelemetry(ctx context.Context, in *ListTemplateTelemetryRequest, opts ...grpc.CallOption) (*ListTemplateTelemetryResponse, error)
	AddTemplateTelemetryExt(ctx context.Context, in *AddTemplateTelemetryExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateTemplateTelemetryExt(ctx context.Context, in *UpdateTemplateTelemetryExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTemplateTelemetryExt(ctx context.Context, in *DeleteTemplateTelemetryExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//operation template command
	AddTemplateCommand(ctx context.Context, in *AddTemplateCommandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateTemplateCommand(ctx context.Context, in *UpdateTemplateCommandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTemplateCommand(ctx context.Context, in *DeleteTemplateCommandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTemplateCommand(ctx context.Context, in *GetTemplateCommandRequest, opts ...grpc.CallOption) (*GetTemplateCommandResponse, error)
	ListTemplateCommand(ctx context.Context, in *ListTemplateCommandRequest, opts ...grpc.CallOption) (*ListTemplateCommandResponse, error)
}

type templateClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateClient(cc grpc.ClientConnInterface) TemplateClient {
	return &templateClient{cc}
}

func (c *templateClient) CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error) {
	out := new(CreateTemplateResponse)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) UpdateTemplate(ctx context.Context, in *UpdateTemplateRequest, opts ...grpc.CallOption) (*UpdateTemplateResponse, error) {
	out := new(UpdateTemplateResponse)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/UpdateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) DeleteTemplate(ctx context.Context, in *DeleteTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/DeleteTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error) {
	out := new(GetTemplateResponse)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) AddTemplateAttribute(ctx context.Context, in *AddTemplateAttributeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/AddTemplateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) UpdateTemplateAttribute(ctx context.Context, in *UpdateTemplateAttributeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/UpdateTemplateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) DeleteTemplateAttribute(ctx context.Context, in *DeleteTemplateAttributeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/DeleteTemplateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) GetTemplateAttribute(ctx context.Context, in *GetTemplateAttributeRequest, opts ...grpc.CallOption) (*GetTemplateAttributeResponse, error) {
	out := new(GetTemplateAttributeResponse)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/GetTemplateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) ListTemplateAttribute(ctx context.Context, in *ListTemplateAttributeRequest, opts ...grpc.CallOption) (*ListTemplateAttributeResponse, error) {
	out := new(ListTemplateAttributeResponse)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/ListTemplateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) AddTemplateTelemetry(ctx context.Context, in *AddTemplateTelemetryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/AddTemplateTelemetry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) UpdateTemplateTelemetry(ctx context.Context, in *UpdateTemplateTelemetryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/UpdateTemplateTelemetry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) DeleteTemplateTelemetry(ctx context.Context, in *DeleteTemplateTelemetryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/DeleteTemplateTelemetry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) GetTemplateTelemetry(ctx context.Context, in *GetTemplateTelemetryRequest, opts ...grpc.CallOption) (*GetTemplateTelemetryResponse, error) {
	out := new(GetTemplateTelemetryResponse)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/GetTemplateTelemetry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) ListTemplateTelemetry(ctx context.Context, in *ListTemplateTelemetryRequest, opts ...grpc.CallOption) (*ListTemplateTelemetryResponse, error) {
	out := new(ListTemplateTelemetryResponse)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/ListTemplateTelemetry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) AddTemplateTelemetryExt(ctx context.Context, in *AddTemplateTelemetryExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/AddTemplateTelemetryExt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) UpdateTemplateTelemetryExt(ctx context.Context, in *UpdateTemplateTelemetryExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/UpdateTemplateTelemetryExt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) DeleteTemplateTelemetryExt(ctx context.Context, in *DeleteTemplateTelemetryExtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/DeleteTemplateTelemetryExt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) AddTemplateCommand(ctx context.Context, in *AddTemplateCommandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/AddTemplateCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) UpdateTemplateCommand(ctx context.Context, in *UpdateTemplateCommandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/UpdateTemplateCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) DeleteTemplateCommand(ctx context.Context, in *DeleteTemplateCommandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/DeleteTemplateCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) GetTemplateCommand(ctx context.Context, in *GetTemplateCommandRequest, opts ...grpc.CallOption) (*GetTemplateCommandResponse, error) {
	out := new(GetTemplateCommandResponse)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/GetTemplateCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) ListTemplateCommand(ctx context.Context, in *ListTemplateCommandRequest, opts ...grpc.CallOption) (*ListTemplateCommandResponse, error) {
	out := new(ListTemplateCommandResponse)
	err := c.cc.Invoke(ctx, "/api.template.v1.Template/ListTemplateCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServer is the server API for Template service.
// All implementations must embed UnimplementedTemplateServer
// for forward compatibility
type TemplateServer interface {
	//Operation template entity
	CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error)
	UpdateTemplate(context.Context, *UpdateTemplateRequest) (*UpdateTemplateResponse, error)
	DeleteTemplate(context.Context, *DeleteTemplateRequest) (*emptypb.Empty, error)
	GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error)
	//operation template attribute
	AddTemplateAttribute(context.Context, *AddTemplateAttributeRequest) (*emptypb.Empty, error)
	UpdateTemplateAttribute(context.Context, *UpdateTemplateAttributeRequest) (*emptypb.Empty, error)
	DeleteTemplateAttribute(context.Context, *DeleteTemplateAttributeRequest) (*emptypb.Empty, error)
	GetTemplateAttribute(context.Context, *GetTemplateAttributeRequest) (*GetTemplateAttributeResponse, error)
	ListTemplateAttribute(context.Context, *ListTemplateAttributeRequest) (*ListTemplateAttributeResponse, error)
	//operation template telemetry
	AddTemplateTelemetry(context.Context, *AddTemplateTelemetryRequest) (*emptypb.Empty, error)
	UpdateTemplateTelemetry(context.Context, *UpdateTemplateTelemetryRequest) (*emptypb.Empty, error)
	DeleteTemplateTelemetry(context.Context, *DeleteTemplateTelemetryRequest) (*emptypb.Empty, error)
	GetTemplateTelemetry(context.Context, *GetTemplateTelemetryRequest) (*GetTemplateTelemetryResponse, error)
	ListTemplateTelemetry(context.Context, *ListTemplateTelemetryRequest) (*ListTemplateTelemetryResponse, error)
	AddTemplateTelemetryExt(context.Context, *AddTemplateTelemetryExtRequest) (*emptypb.Empty, error)
	UpdateTemplateTelemetryExt(context.Context, *UpdateTemplateTelemetryExtRequest) (*emptypb.Empty, error)
	DeleteTemplateTelemetryExt(context.Context, *DeleteTemplateTelemetryExtRequest) (*emptypb.Empty, error)
	//operation template command
	AddTemplateCommand(context.Context, *AddTemplateCommandRequest) (*emptypb.Empty, error)
	UpdateTemplateCommand(context.Context, *UpdateTemplateCommandRequest) (*emptypb.Empty, error)
	DeleteTemplateCommand(context.Context, *DeleteTemplateCommandRequest) (*emptypb.Empty, error)
	GetTemplateCommand(context.Context, *GetTemplateCommandRequest) (*GetTemplateCommandResponse, error)
	ListTemplateCommand(context.Context, *ListTemplateCommandRequest) (*ListTemplateCommandResponse, error)
	mustEmbedUnimplementedTemplateServer()
}

// UnimplementedTemplateServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateServer struct {
}

func (UnimplementedTemplateServer) CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}
func (UnimplementedTemplateServer) UpdateTemplate(context.Context, *UpdateTemplateRequest) (*UpdateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplate not implemented")
}
func (UnimplementedTemplateServer) DeleteTemplate(context.Context, *DeleteTemplateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplate not implemented")
}
func (UnimplementedTemplateServer) GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}
func (UnimplementedTemplateServer) AddTemplateAttribute(context.Context, *AddTemplateAttributeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTemplateAttribute not implemented")
}
func (UnimplementedTemplateServer) UpdateTemplateAttribute(context.Context, *UpdateTemplateAttributeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplateAttribute not implemented")
}
func (UnimplementedTemplateServer) DeleteTemplateAttribute(context.Context, *DeleteTemplateAttributeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplateAttribute not implemented")
}
func (UnimplementedTemplateServer) GetTemplateAttribute(context.Context, *GetTemplateAttributeRequest) (*GetTemplateAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplateAttribute not implemented")
}
func (UnimplementedTemplateServer) ListTemplateAttribute(context.Context, *ListTemplateAttributeRequest) (*ListTemplateAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTemplateAttribute not implemented")
}
func (UnimplementedTemplateServer) AddTemplateTelemetry(context.Context, *AddTemplateTelemetryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTemplateTelemetry not implemented")
}
func (UnimplementedTemplateServer) UpdateTemplateTelemetry(context.Context, *UpdateTemplateTelemetryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplateTelemetry not implemented")
}
func (UnimplementedTemplateServer) DeleteTemplateTelemetry(context.Context, *DeleteTemplateTelemetryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplateTelemetry not implemented")
}
func (UnimplementedTemplateServer) GetTemplateTelemetry(context.Context, *GetTemplateTelemetryRequest) (*GetTemplateTelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplateTelemetry not implemented")
}
func (UnimplementedTemplateServer) ListTemplateTelemetry(context.Context, *ListTemplateTelemetryRequest) (*ListTemplateTelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTemplateTelemetry not implemented")
}
func (UnimplementedTemplateServer) AddTemplateTelemetryExt(context.Context, *AddTemplateTelemetryExtRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTemplateTelemetryExt not implemented")
}
func (UnimplementedTemplateServer) UpdateTemplateTelemetryExt(context.Context, *UpdateTemplateTelemetryExtRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplateTelemetryExt not implemented")
}
func (UnimplementedTemplateServer) DeleteTemplateTelemetryExt(context.Context, *DeleteTemplateTelemetryExtRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplateTelemetryExt not implemented")
}
func (UnimplementedTemplateServer) AddTemplateCommand(context.Context, *AddTemplateCommandRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTemplateCommand not implemented")
}
func (UnimplementedTemplateServer) UpdateTemplateCommand(context.Context, *UpdateTemplateCommandRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplateCommand not implemented")
}
func (UnimplementedTemplateServer) DeleteTemplateCommand(context.Context, *DeleteTemplateCommandRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplateCommand not implemented")
}
func (UnimplementedTemplateServer) GetTemplateCommand(context.Context, *GetTemplateCommandRequest) (*GetTemplateCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplateCommand not implemented")
}
func (UnimplementedTemplateServer) ListTemplateCommand(context.Context, *ListTemplateCommandRequest) (*ListTemplateCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTemplateCommand not implemented")
}
func (UnimplementedTemplateServer) mustEmbedUnimplementedTemplateServer() {}

// UnsafeTemplateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServer will
// result in compilation errors.
type UnsafeTemplateServer interface {
	mustEmbedUnimplementedTemplateServer()
}

func RegisterTemplateServer(s grpc.ServiceRegistrar, srv TemplateServer) {
	s.RegisterService(&Template_ServiceDesc, srv)
}

func _Template_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).CreateTemplate(ctx, req.(*CreateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_UpdateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).UpdateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/UpdateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).UpdateTemplate(ctx, req.(*UpdateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_DeleteTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).DeleteTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/DeleteTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).DeleteTemplate(ctx, req.(*DeleteTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).GetTemplate(ctx, req.(*GetTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_AddTemplateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTemplateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).AddTemplateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/AddTemplateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).AddTemplateAttribute(ctx, req.(*AddTemplateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_UpdateTemplateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTemplateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).UpdateTemplateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/UpdateTemplateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).UpdateTemplateAttribute(ctx, req.(*UpdateTemplateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_DeleteTemplateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTemplateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).DeleteTemplateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/DeleteTemplateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).DeleteTemplateAttribute(ctx, req.(*DeleteTemplateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_GetTemplateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).GetTemplateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/GetTemplateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).GetTemplateAttribute(ctx, req.(*GetTemplateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_ListTemplateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTemplateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).ListTemplateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/ListTemplateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).ListTemplateAttribute(ctx, req.(*ListTemplateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_AddTemplateTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTemplateTelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).AddTemplateTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/AddTemplateTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).AddTemplateTelemetry(ctx, req.(*AddTemplateTelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_UpdateTemplateTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTemplateTelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).UpdateTemplateTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/UpdateTemplateTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).UpdateTemplateTelemetry(ctx, req.(*UpdateTemplateTelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_DeleteTemplateTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTemplateTelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).DeleteTemplateTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/DeleteTemplateTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).DeleteTemplateTelemetry(ctx, req.(*DeleteTemplateTelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_GetTemplateTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateTelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).GetTemplateTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/GetTemplateTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).GetTemplateTelemetry(ctx, req.(*GetTemplateTelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_ListTemplateTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTemplateTelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).ListTemplateTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/ListTemplateTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).ListTemplateTelemetry(ctx, req.(*ListTemplateTelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_AddTemplateTelemetryExt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTemplateTelemetryExtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).AddTemplateTelemetryExt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/AddTemplateTelemetryExt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).AddTemplateTelemetryExt(ctx, req.(*AddTemplateTelemetryExtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_UpdateTemplateTelemetryExt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTemplateTelemetryExtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).UpdateTemplateTelemetryExt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/UpdateTemplateTelemetryExt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).UpdateTemplateTelemetryExt(ctx, req.(*UpdateTemplateTelemetryExtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_DeleteTemplateTelemetryExt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTemplateTelemetryExtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).DeleteTemplateTelemetryExt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/DeleteTemplateTelemetryExt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).DeleteTemplateTelemetryExt(ctx, req.(*DeleteTemplateTelemetryExtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_AddTemplateCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTemplateCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).AddTemplateCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/AddTemplateCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).AddTemplateCommand(ctx, req.(*AddTemplateCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_UpdateTemplateCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTemplateCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).UpdateTemplateCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/UpdateTemplateCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).UpdateTemplateCommand(ctx, req.(*UpdateTemplateCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_DeleteTemplateCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTemplateCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).DeleteTemplateCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/DeleteTemplateCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).DeleteTemplateCommand(ctx, req.(*DeleteTemplateCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_GetTemplateCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).GetTemplateCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/GetTemplateCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).GetTemplateCommand(ctx, req.(*GetTemplateCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_ListTemplateCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTemplateCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).ListTemplateCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.template.v1.Template/ListTemplateCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).ListTemplateCommand(ctx, req.(*ListTemplateCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Template_ServiceDesc is the grpc.ServiceDesc for Template service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Template_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.template.v1.Template",
	HandlerType: (*TemplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTemplate",
			Handler:    _Template_CreateTemplate_Handler,
		},
		{
			MethodName: "UpdateTemplate",
			Handler:    _Template_UpdateTemplate_Handler,
		},
		{
			MethodName: "DeleteTemplate",
			Handler:    _Template_DeleteTemplate_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _Template_GetTemplate_Handler,
		},
		{
			MethodName: "AddTemplateAttribute",
			Handler:    _Template_AddTemplateAttribute_Handler,
		},
		{
			MethodName: "UpdateTemplateAttribute",
			Handler:    _Template_UpdateTemplateAttribute_Handler,
		},
		{
			MethodName: "DeleteTemplateAttribute",
			Handler:    _Template_DeleteTemplateAttribute_Handler,
		},
		{
			MethodName: "GetTemplateAttribute",
			Handler:    _Template_GetTemplateAttribute_Handler,
		},
		{
			MethodName: "ListTemplateAttribute",
			Handler:    _Template_ListTemplateAttribute_Handler,
		},
		{
			MethodName: "AddTemplateTelemetry",
			Handler:    _Template_AddTemplateTelemetry_Handler,
		},
		{
			MethodName: "UpdateTemplateTelemetry",
			Handler:    _Template_UpdateTemplateTelemetry_Handler,
		},
		{
			MethodName: "DeleteTemplateTelemetry",
			Handler:    _Template_DeleteTemplateTelemetry_Handler,
		},
		{
			MethodName: "GetTemplateTelemetry",
			Handler:    _Template_GetTemplateTelemetry_Handler,
		},
		{
			MethodName: "ListTemplateTelemetry",
			Handler:    _Template_ListTemplateTelemetry_Handler,
		},
		{
			MethodName: "AddTemplateTelemetryExt",
			Handler:    _Template_AddTemplateTelemetryExt_Handler,
		},
		{
			MethodName: "UpdateTemplateTelemetryExt",
			Handler:    _Template_UpdateTemplateTelemetryExt_Handler,
		},
		{
			MethodName: "DeleteTemplateTelemetryExt",
			Handler:    _Template_DeleteTemplateTelemetryExt_Handler,
		},
		{
			MethodName: "AddTemplateCommand",
			Handler:    _Template_AddTemplateCommand_Handler,
		},
		{
			MethodName: "UpdateTemplateCommand",
			Handler:    _Template_UpdateTemplateCommand_Handler,
		},
		{
			MethodName: "DeleteTemplateCommand",
			Handler:    _Template_DeleteTemplateCommand_Handler,
		},
		{
			MethodName: "GetTemplateCommand",
			Handler:    _Template_GetTemplateCommand_Handler,
		},
		{
			MethodName: "ListTemplateCommand",
			Handler:    _Template_ListTemplateCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/template/v1/template.proto",
}
