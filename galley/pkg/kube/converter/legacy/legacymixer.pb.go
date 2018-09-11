// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: legacymixer.proto

/*
Package legacy is a generated protocol buffer package.

TODO: Temporarily placing this file in the istio repo. Eventually this should go into istio.io/api.

It is generated from these files:
	legacymixer.proto

It has these top-level messages:
	LegacyMixerResource
*/
package legacy

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// LegacyMixerResource is used to multiplex old-style one-per-kind Mixer instances and templates through
// the MCP protocol.
type LegacyMixerResource struct {
	// The original name of the resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The original kind of the resource.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// The original contents of the resource.
	Contents *google_protobuf.Struct `protobuf:"bytes,3,opt,name=contents" json:"contents,omitempty"`
}

func (m *LegacyMixerResource) Reset()                    { *m = LegacyMixerResource{} }
func (m *LegacyMixerResource) String() string            { return proto.CompactTextString(m) }
func (*LegacyMixerResource) ProtoMessage()               {}
func (*LegacyMixerResource) Descriptor() ([]byte, []int) { return fileDescriptorLegacymixer, []int{0} }

func (m *LegacyMixerResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LegacyMixerResource) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *LegacyMixerResource) GetContents() *google_protobuf.Struct {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterType((*LegacyMixerResource)(nil), "istio.mcp.v1alpha1.extensions.LegacyMixerResource")
}

func init() { proto.RegisterFile("legacymixer.proto", fileDescriptorLegacymixer) }

var fileDescriptorLegacymixer = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x31, 0x4f, 0x04, 0x21,
	0x10, 0x85, 0xb3, 0x6a, 0x8c, 0x62, 0x25, 0x16, 0x6e, 0x8c, 0x26, 0x17, 0xab, 0xab, 0x98, 0xac,
	0xf7, 0x0f, 0xac, 0xb5, 0x59, 0x3b, 0x3b, 0x16, 0x47, 0x24, 0xcb, 0x32, 0x04, 0x86, 0xcb, 0xed,
	0xbf, 0x37, 0x0b, 0xd1, 0xee, 0xe5, 0xe5, 0xf1, 0x7d, 0x8c, 0xb8, 0xf5, 0x68, 0xb5, 0x59, 0x17,
	0x77, 0xc2, 0xa4, 0x62, 0x22, 0x26, 0xf9, 0xe4, 0x32, 0x3b, 0x52, 0x8b, 0x89, 0xea, 0x38, 0x68,
	0x1f, 0x7f, 0xf4, 0xa0, 0xf0, 0xc4, 0x18, 0xb2, 0xa3, 0x90, 0x1f, 0x1e, 0x2d, 0x91, 0xf5, 0x08,
	0x75, 0x3c, 0x95, 0x6f, 0xc8, 0x9c, 0x8a, 0xe1, 0xf6, 0xf8, 0x39, 0x89, 0xbb, 0xb7, 0x4a, 0x7c,
	0xdf, 0x88, 0x23, 0x66, 0x2a, 0xc9, 0xa0, 0x94, 0xe2, 0x22, 0xe8, 0x05, 0xfb, 0x6e, 0xd7, 0xed,
	0xaf, 0xc7, 0x9a, 0xb7, 0x6e, 0x76, 0xe1, 0xab, 0x3f, 0x6b, 0xdd, 0x96, 0xe5, 0x41, 0x5c, 0x19,
	0x0a, 0x8c, 0x81, 0x73, 0x7f, 0xbe, 0xeb, 0xf6, 0x37, 0x2f, 0xf7, 0xaa, 0xf9, 0xd4, 0x9f, 0x4f,
	0x7d, 0x54, 0xdf, 0xf8, 0x3f, 0x7c, 0x1d, 0x3e, 0xa1, 0x7d, 0xd9, 0x51, 0x0b, 0x60, 0xb5, 0xf7,
	0xb8, 0x42, 0x9c, 0x2d, 0xcc, 0x65, 0x42, 0x30, 0x14, 0x8e, 0x98, 0x18, 0x13, 0xb4, 0x73, 0xa7,
	0xcb, 0x4a, 0x3b, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x08, 0xe2, 0x38, 0x61, 0xff, 0x00, 0x00,
	0x00,
}