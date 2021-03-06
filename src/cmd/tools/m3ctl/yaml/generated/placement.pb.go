// Code generated by protoc-gen-go. DO NOT EDIT.
// source: placement.proto

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package yaml

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	admin "github.com/m3db/m3/src/query/generated/proto/admin"
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

type PlacementInitRequestYaml struct {
	Operation            string                      `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Request              *admin.PlacementInitRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *PlacementInitRequestYaml) Reset()         { *m = PlacementInitRequestYaml{} }
func (m *PlacementInitRequestYaml) String() string { return proto.CompactTextString(m) }
func (*PlacementInitRequestYaml) ProtoMessage()    {}
func (*PlacementInitRequestYaml) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae0216eeb0d08e49, []int{0}
}

func (m *PlacementInitRequestYaml) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlacementInitRequestYaml.Unmarshal(m, b)
}
func (m *PlacementInitRequestYaml) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlacementInitRequestYaml.Marshal(b, m, deterministic)
}
func (m *PlacementInitRequestYaml) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlacementInitRequestYaml.Merge(m, src)
}
func (m *PlacementInitRequestYaml) XXX_Size() int {
	return xxx_messageInfo_PlacementInitRequestYaml.Size(m)
}
func (m *PlacementInitRequestYaml) XXX_DiscardUnknown() {
	xxx_messageInfo_PlacementInitRequestYaml.DiscardUnknown(m)
}

var xxx_messageInfo_PlacementInitRequestYaml proto.InternalMessageInfo

func (m *PlacementInitRequestYaml) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *PlacementInitRequestYaml) GetRequest() *admin.PlacementInitRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type PlacementReplaceRequestYaml struct {
	Operation            string                         `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	Request              *admin.PlacementReplaceRequest `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *PlacementReplaceRequestYaml) Reset()         { *m = PlacementReplaceRequestYaml{} }
func (m *PlacementReplaceRequestYaml) String() string { return proto.CompactTextString(m) }
func (*PlacementReplaceRequestYaml) ProtoMessage()    {}
func (*PlacementReplaceRequestYaml) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae0216eeb0d08e49, []int{1}
}

func (m *PlacementReplaceRequestYaml) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlacementReplaceRequestYaml.Unmarshal(m, b)
}
func (m *PlacementReplaceRequestYaml) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlacementReplaceRequestYaml.Marshal(b, m, deterministic)
}
func (m *PlacementReplaceRequestYaml) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlacementReplaceRequestYaml.Merge(m, src)
}
func (m *PlacementReplaceRequestYaml) XXX_Size() int {
	return xxx_messageInfo_PlacementReplaceRequestYaml.Size(m)
}
func (m *PlacementReplaceRequestYaml) XXX_DiscardUnknown() {
	xxx_messageInfo_PlacementReplaceRequestYaml.DiscardUnknown(m)
}

var xxx_messageInfo_PlacementReplaceRequestYaml proto.InternalMessageInfo

func (m *PlacementReplaceRequestYaml) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *PlacementReplaceRequestYaml) GetRequest() *admin.PlacementReplaceRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto.RegisterType((*PlacementInitRequestYaml)(nil), "yaml.PlacementInitRequestYaml")
	proto.RegisterType((*PlacementReplaceRequestYaml)(nil), "yaml.PlacementReplaceRequestYaml")
}

func init() { proto.RegisterFile("placement.proto", fileDescriptor_ae0216eeb0d08e49) }

var fileDescriptor_ae0216eeb0d08e49 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc8, 0x49, 0x4c,
	0x4e, 0xcd, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xa9, 0x4c, 0xcc,
	0xcd, 0x91, 0x72, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x35,
	0x4e, 0x49, 0xd2, 0xcf, 0x35, 0xd6, 0x2f, 0x2e, 0x4a, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4,
	0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0x49, 0x4d, 0xd1, 0x07, 0xeb, 0xd1, 0x4f, 0x4c, 0xc9,
	0xcd, 0xcc, 0xd3, 0x47, 0x33, 0x49, 0x29, 0x9f, 0x4b, 0x22, 0x00, 0x26, 0xe4, 0x99, 0x97, 0x59,
	0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x12, 0x99, 0x98, 0x9b, 0x23, 0x24, 0xc3, 0xc5, 0x99,
	0x5f, 0x00, 0x32, 0x23, 0x33, 0x3f, 0x4f, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x21, 0x20,
	0x64, 0xca, 0xc5, 0x5e, 0x04, 0x51, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xad, 0x07,
	0xb6, 0x42, 0x0f, 0x9b, 0x79, 0x41, 0x30, 0xb5, 0x4a, 0xa5, 0x5c, 0xd2, 0x70, 0x05, 0x41, 0xa9,
	0x60, 0xe7, 0xe0, 0xb4, 0x93, 0x19, 0xdd, 0x4e, 0x0b, 0x84, 0x9d, 0x2c, 0x60, 0x3b, 0xe5, 0xd0,
	0xed, 0x44, 0x35, 0x12, 0x6e, 0x6d, 0x12, 0x1b, 0xd8, 0xbb, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x5e, 0x24, 0x27, 0x1e, 0x4b, 0x01, 0x00, 0x00,
}
