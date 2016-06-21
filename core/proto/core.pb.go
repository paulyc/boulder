// Code generated by protoc-gen-go.
// source: core/proto/core.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	core/proto/core.proto

It has these top-level messages:
	Challenge
	ValidationRecord
	ProblemDetails
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto1.ProtoPackageIsVersion1

type Challenge struct {
	Id                *int64              `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type              *string             `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Status            *string             `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	Uri               *string             `protobuf:"bytes,9,opt,name=uri" json:"uri,omitempty"`
	Token             *string             `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	KeyAuthorization  *string             `protobuf:"bytes,5,opt,name=keyAuthorization" json:"keyAuthorization,omitempty"`
	Validationrecords []*ValidationRecord `protobuf:"bytes,10,rep,name=validationrecords" json:"validationrecords,omitempty"`
	Error             *ProblemDetails     `protobuf:"bytes,7,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized  []byte              `json:"-"`
}

func (m *Challenge) Reset()                    { *m = Challenge{} }
func (m *Challenge) String() string            { return proto1.CompactTextString(m) }
func (*Challenge) ProtoMessage()               {}
func (*Challenge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Challenge) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Challenge) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Challenge) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *Challenge) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Challenge) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Challenge) GetKeyAuthorization() string {
	if m != nil && m.KeyAuthorization != nil {
		return *m.KeyAuthorization
	}
	return ""
}

func (m *Challenge) GetValidationrecords() []*ValidationRecord {
	if m != nil {
		return m.Validationrecords
	}
	return nil
}

func (m *Challenge) GetError() *ProblemDetails {
	if m != nil {
		return m.Error
	}
	return nil
}

type ValidationRecord struct {
	Hostname          *string  `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Port              *string  `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	AddressesResolved [][]byte `protobuf:"bytes,3,rep,name=addressesResolved" json:"addressesResolved,omitempty"`
	AddressUsed       []byte   `protobuf:"bytes,4,opt,name=addressUsed" json:"addressUsed,omitempty"`
	Authorities       []string `protobuf:"bytes,5,rep,name=authorities" json:"authorities,omitempty"`
	Url               *string  `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *ValidationRecord) Reset()                    { *m = ValidationRecord{} }
func (m *ValidationRecord) String() string            { return proto1.CompactTextString(m) }
func (*ValidationRecord) ProtoMessage()               {}
func (*ValidationRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ValidationRecord) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *ValidationRecord) GetPort() string {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return ""
}

func (m *ValidationRecord) GetAddressesResolved() [][]byte {
	if m != nil {
		return m.AddressesResolved
	}
	return nil
}

func (m *ValidationRecord) GetAddressUsed() []byte {
	if m != nil {
		return m.AddressUsed
	}
	return nil
}

func (m *ValidationRecord) GetAuthorities() []string {
	if m != nil {
		return m.Authorities
	}
	return nil
}

func (m *ValidationRecord) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

type ProblemDetails struct {
	ProblemType      *string `protobuf:"bytes,1,opt,name=problemType" json:"problemType,omitempty"`
	Detail           *string `protobuf:"bytes,2,opt,name=detail" json:"detail,omitempty"`
	HttpStatus       *int32  `protobuf:"varint,3,opt,name=httpStatus" json:"httpStatus,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProblemDetails) Reset()                    { *m = ProblemDetails{} }
func (m *ProblemDetails) String() string            { return proto1.CompactTextString(m) }
func (*ProblemDetails) ProtoMessage()               {}
func (*ProblemDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProblemDetails) GetProblemType() string {
	if m != nil && m.ProblemType != nil {
		return *m.ProblemType
	}
	return ""
}

func (m *ProblemDetails) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

func (m *ProblemDetails) GetHttpStatus() int32 {
	if m != nil && m.HttpStatus != nil {
		return *m.HttpStatus
	}
	return 0
}

func init() {
	proto1.RegisterType((*Challenge)(nil), "core.Challenge")
	proto1.RegisterType((*ValidationRecord)(nil), "core.ValidationRecord")
	proto1.RegisterType((*ProblemDetails)(nil), "core.ProblemDetails")
}

var fileDescriptor0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x51, 0xcb, 0x4e, 0x02, 0x41,
	0x10, 0xcc, 0xb2, 0x2c, 0xb8, 0xbd, 0x48, 0x60, 0x7d, 0x64, 0xbc, 0x11, 0xbc, 0x70, 0x82, 0xc8,
	0x1f, 0xf8, 0xb8, 0x78, 0x33, 0xf8, 0x38, 0x78, 0x1b, 0x9d, 0x8e, 0x3b, 0x61, 0xd8, 0xd9, 0xcc,
	0x34, 0x24, 0x78, 0xf6, 0xcf, 0xfc, 0x31, 0xe7, 0xb1, 0x98, 0x18, 0x6f, 0x5d, 0x55, 0xdd, 0x93,
	0xaa, 0x1a, 0x38, 0x7b, 0xd7, 0x06, 0x17, 0x8d, 0xd1, 0xa4, 0x17, 0x7e, 0x9c, 0x87, 0xb1, 0xec,
	0xfa, 0x79, 0xfa, 0x9d, 0x40, 0x7e, 0x5b, 0x71, 0xa5, 0xb0, 0xfe, 0xc0, 0x12, 0xa0, 0x23, 0x05,
	0x4b, 0x26, 0xc9, 0x2c, 0x2d, 0x07, 0xd0, 0xa5, 0x7d, 0x83, 0xac, 0xe3, 0x50, 0x5e, 0x0e, 0xa1,
	0x67, 0x89, 0xd3, 0xd6, 0xb2, 0x5e, 0xc0, 0x05, 0xa4, 0x5b, 0x23, 0x59, 0x1e, 0xc0, 0x31, 0x64,
	0xa4, 0xd7, 0x58, 0xb3, 0x34, 0x40, 0x06, 0xa3, 0x35, 0xee, 0xaf, 0xb7, 0x54, 0x69, 0x23, 0x3f,
	0x39, 0x49, 0x5d, 0xb3, 0x2c, 0x28, 0x57, 0x30, 0xde, 0x71, 0x25, 0x45, 0xe0, 0x0c, 0x3a, 0x07,
	0xc2, 0x32, 0x98, 0xa4, 0xb3, 0x62, 0x79, 0x3e, 0x0f, 0xde, 0x5e, 0x7e, 0xe5, 0x55, 0x90, 0xcb,
	0x4b, 0xc8, 0xd0, 0x18, 0x6d, 0x58, 0xdf, 0xbd, 0x50, 0x2c, 0x4f, 0xe3, 0xda, 0x83, 0xd1, 0x6f,
	0x0a, 0x37, 0x77, 0x48, 0x5c, 0x2a, 0x3b, 0xfd, 0x4a, 0x60, 0xf4, 0xef, 0x72, 0x04, 0x47, 0x95,
	0xb6, 0x54, 0xf3, 0x0d, 0x86, 0x48, 0xb9, 0x8f, 0xd4, 0x68, 0x43, 0x6d, 0xa4, 0x0b, 0x18, 0x73,
	0x21, 0x0c, 0x5a, 0x8b, 0x76, 0x85, 0x56, 0xab, 0x1d, 0x0a, 0x97, 0x20, 0x9d, 0x0d, 0xca, 0x13,
	0x28, 0x5a, 0xe9, 0xd9, 0x3a, 0xb2, 0xeb, 0xf6, 0x23, 0x19, 0x33, 0x91, 0x44, 0xeb, 0x12, 0xa5,
	0x87, 0x1e, 0x54, 0x2c, 0x65, 0x7a, 0x0f, 0xc3, 0xbf, 0xc6, 0xfc, 0x4d, 0x13, 0x99, 0x27, 0xdf,
	0x65, 0x72, 0xe8, 0x52, 0x04, 0xbd, 0x35, 0xe2, 0x6a, 0xaf, 0x88, 0x9a, 0xc7, 0xd8, 0xaf, 0xef,
	0x30, 0xbb, 0xe9, 0xbf, 0x66, 0xe1, 0x9b, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xa6, 0xad,
	0x9d, 0xbe, 0x01, 0x00, 0x00,
}
