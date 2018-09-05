// Code generated by protoc-gen-go.
// source: actionplugin.proto
// DO NOT EDIT!

/*
Package actionplugin is a generated protocol buffer package.

It is generated from these files:
	actionplugin.proto

It has these top-level messages:
	ActionPluginManifest
	ActionQuery
	ActionResult
*/
package actionplugin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ActionPluginManifest struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version     string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Author      string `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
}

func (m *ActionPluginManifest) Reset()                    { *m = ActionPluginManifest{} }
func (m *ActionPluginManifest) String() string            { return proto.CompactTextString(m) }
func (*ActionPluginManifest) ProtoMessage()               {}
func (*ActionPluginManifest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActionPluginManifest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActionPluginManifest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ActionPluginManifest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ActionPluginManifest) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type ActionQuery struct {
	Options map[string]string `protobuf:"bytes,1,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ActionQuery) Reset()                    { *m = ActionQuery{} }
func (m *ActionQuery) String() string            { return proto.CompactTextString(m) }
func (*ActionQuery) ProtoMessage()               {}
func (*ActionQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ActionQuery) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

type ActionResult struct {
	Status  string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details" json:"details,omitempty"`
}

func (m *ActionResult) Reset()                    { *m = ActionResult{} }
func (m *ActionResult) String() string            { return proto.CompactTextString(m) }
func (*ActionResult) ProtoMessage()               {}
func (*ActionResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ActionResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ActionResult) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func init() {
	proto.RegisterType((*ActionPluginManifest)(nil), "actionplugin.ActionPluginManifest")
	proto.RegisterType((*ActionQuery)(nil), "actionplugin.ActionQuery")
	proto.RegisterType((*ActionResult)(nil), "actionplugin.ActionResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ActionPlugin service

type ActionPluginClient interface {
	Manifest(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ActionPluginManifest, error)
	Run(ctx context.Context, in *ActionQuery, opts ...grpc.CallOption) (*ActionResult, error)
}

type actionPluginClient struct {
	cc *grpc.ClientConn
}

func NewActionPluginClient(cc *grpc.ClientConn) ActionPluginClient {
	return &actionPluginClient{cc}
}

func (c *actionPluginClient) Manifest(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ActionPluginManifest, error) {
	out := new(ActionPluginManifest)
	err := grpc.Invoke(ctx, "/actionplugin.ActionPlugin/Manifest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionPluginClient) Run(ctx context.Context, in *ActionQuery, opts ...grpc.CallOption) (*ActionResult, error) {
	out := new(ActionResult)
	err := grpc.Invoke(ctx, "/actionplugin.ActionPlugin/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ActionPlugin service

type ActionPluginServer interface {
	Manifest(context.Context, *google_protobuf.Empty) (*ActionPluginManifest, error)
	Run(context.Context, *ActionQuery) (*ActionResult, error)
}

func RegisterActionPluginServer(s *grpc.Server, srv ActionPluginServer) {
	s.RegisterService(&_ActionPlugin_serviceDesc, srv)
}

func _ActionPlugin_Manifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionPluginServer).Manifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actionplugin.ActionPlugin/Manifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionPluginServer).Manifest(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionPlugin_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionPluginServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actionplugin.ActionPlugin/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionPluginServer).Run(ctx, req.(*ActionQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActionPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "actionplugin.ActionPlugin",
	HandlerType: (*ActionPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Manifest",
			Handler:    _ActionPlugin_Manifest_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _ActionPlugin_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "actionplugin.proto",
}

func init() { proto.RegisterFile("actionplugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0xa5, 0x94, 0x1f, 0xfc, 0x1c, 0x38, 0xe8, 0x86, 0x90, 0x5a, 0x2f, 0xa4, 0x07, 0xe5, 0xb4,
	0x24, 0x78, 0x31, 0x1c, 0x0c, 0x92, 0x90, 0x78, 0xd0, 0x88, 0x3d, 0x7a, 0x2b, 0xed, 0x52, 0x1a,
	0xca, 0x6e, 0xb3, 0x7f, 0x48, 0x7a, 0xf1, 0x0b, 0xf8, 0x05, 0xfc, 0xb8, 0xa6, 0xbb, 0x5b, 0x53,
	0x12, 0xbc, 0xed, 0x9b, 0xf7, 0x66, 0x3a, 0xef, 0x4d, 0x01, 0x45, 0xb1, 0xcc, 0x18, 0x2d, 0x72,
	0x95, 0x66, 0x14, 0x17, 0x9c, 0x49, 0x86, 0x06, 0xcd, 0x9a, 0x7f, 0x93, 0x32, 0x96, 0xe6, 0x64,
	0xaa, 0xb9, 0x8d, 0xda, 0x4e, 0xc9, 0xa1, 0x90, 0xa5, 0x91, 0x06, 0x9f, 0x30, 0x7c, 0xd2, 0xe2,
	0xb5, 0x16, 0xbf, 0x46, 0x34, 0xdb, 0x12, 0x21, 0x11, 0x82, 0x0e, 0x8d, 0x0e, 0xc4, 0x73, 0xc6,
	0xce, 0xe4, 0x22, 0xd4, 0x6f, 0xe4, 0x41, 0xef, 0x48, 0xb8, 0xc8, 0x18, 0xf5, 0xda, 0xba, 0x5c,
	0x43, 0x34, 0x86, 0x7e, 0x42, 0x44, 0xcc, 0xb3, 0xa2, 0x1a, 0xe5, 0xb9, 0x9a, 0x6d, 0x96, 0xd0,
	0x08, 0xba, 0x91, 0x92, 0x3b, 0xc6, 0xbd, 0x8e, 0x26, 0x2d, 0x0a, 0xbe, 0x1c, 0xe8, 0x9b, 0x05,
	0xde, 0x15, 0xe1, 0x25, 0x5a, 0x40, 0x8f, 0xe9, 0x0e, 0xe1, 0x39, 0x63, 0x77, 0xd2, 0x9f, 0xdd,
	0xe2, 0x13, 0x83, 0x0d, 0x2d, 0x7e, 0x33, 0xc2, 0x15, 0x95, 0xbc, 0x0c, 0xeb, 0x36, 0x7f, 0x0e,
	0x83, 0x26, 0x81, 0x2e, 0xc1, 0xdd, 0x93, 0xd2, 0x1a, 0xa9, 0x9e, 0x68, 0x08, 0xff, 0x8e, 0x51,
	0xae, 0x88, 0x75, 0x61, 0xc0, 0xbc, 0xfd, 0xe0, 0x04, 0x0b, 0x18, 0x98, 0x0f, 0x84, 0x44, 0xa8,
	0x5c, 0x56, 0x5b, 0x0b, 0x19, 0x49, 0x25, 0x6c, 0xbb, 0x45, 0x55, 0x12, 0x09, 0x91, 0x51, 0x96,
	0x8b, 0x3a, 0x09, 0x0b, 0x67, 0xdf, 0x4e, 0x3d, 0xc2, 0x04, 0x8a, 0x9e, 0xe1, 0xff, 0x6f, 0xa8,
	0x23, 0x6c, 0x4e, 0x81, 0xeb, 0x53, 0xe0, 0x55, 0x75, 0x0a, 0x3f, 0x38, 0xe7, 0xf1, 0xf4, 0x20,
	0x41, 0x0b, 0x3d, 0x82, 0x1b, 0x2a, 0x8a, 0xae, 0xff, 0x0c, 0xc4, 0xf7, 0xcf, 0x51, 0xc6, 0x4a,
	0xd0, 0x5a, 0xbe, 0xc0, 0x5d, 0xcc, 0x0e, 0x98, 0x1d, 0x77, 0x38, 0x4e, 0x04, 0x16, 0xc9, 0x1e,
	0xa7, 0xbc, 0x88, 0xad, 0xb8, 0xd9, 0xb9, 0xbc, 0x6a, 0xae, 0xb0, 0xae, 0x96, 0x5d, 0x3b, 0x1f,
	0x27, 0x7f, 0xd5, 0xa6, 0xab, 0x3d, 0xdc, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x91, 0xca,
	0x58, 0x80, 0x02, 0x00, 0x00,
}