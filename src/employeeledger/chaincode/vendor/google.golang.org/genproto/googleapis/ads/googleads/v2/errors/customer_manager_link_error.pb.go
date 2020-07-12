// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/customer_manager_link_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible CustomerManagerLink errors.
type CustomerManagerLinkErrorEnum_CustomerManagerLinkError int32

const (
	// Enum unspecified.
	CustomerManagerLinkErrorEnum_UNSPECIFIED CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 0
	// The received error code is not known in this version.
	CustomerManagerLinkErrorEnum_UNKNOWN CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 1
	// No pending invitation.
	CustomerManagerLinkErrorEnum_NO_PENDING_INVITE CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 2
	// Attempt to operate on the same client more than once in the same call.
	CustomerManagerLinkErrorEnum_SAME_CLIENT_MORE_THAN_ONCE_PER_CALL CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 3
	// Manager account has the maximum number of linked accounts.
	CustomerManagerLinkErrorEnum_MANAGER_HAS_MAX_NUMBER_OF_LINKED_ACCOUNTS CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 4
	// If no active user on account it cannot be unlinked from its manager.
	CustomerManagerLinkErrorEnum_CANNOT_UNLINK_ACCOUNT_WITHOUT_ACTIVE_USER CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 5
	// Account should have at least one active owner on it before being
	// unlinked.
	CustomerManagerLinkErrorEnum_CANNOT_REMOVE_LAST_CLIENT_ACCOUNT_OWNER CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 6
	// Only account owners may change their permission role.
	CustomerManagerLinkErrorEnum_CANNOT_CHANGE_ROLE_BY_NON_ACCOUNT_OWNER CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 7
	// When a client's link to its manager is not active, the link role cannot
	// be changed.
	CustomerManagerLinkErrorEnum_CANNOT_CHANGE_ROLE_FOR_NON_ACTIVE_LINK_ACCOUNT CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 8
	// Attempt to link a child to a parent that contains or will contain
	// duplicate children.
	CustomerManagerLinkErrorEnum_DUPLICATE_CHILD_FOUND CustomerManagerLinkErrorEnum_CustomerManagerLinkError = 9
)

var CustomerManagerLinkErrorEnum_CustomerManagerLinkError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "NO_PENDING_INVITE",
	3: "SAME_CLIENT_MORE_THAN_ONCE_PER_CALL",
	4: "MANAGER_HAS_MAX_NUMBER_OF_LINKED_ACCOUNTS",
	5: "CANNOT_UNLINK_ACCOUNT_WITHOUT_ACTIVE_USER",
	6: "CANNOT_REMOVE_LAST_CLIENT_ACCOUNT_OWNER",
	7: "CANNOT_CHANGE_ROLE_BY_NON_ACCOUNT_OWNER",
	8: "CANNOT_CHANGE_ROLE_FOR_NON_ACTIVE_LINK_ACCOUNT",
	9: "DUPLICATE_CHILD_FOUND",
}

var CustomerManagerLinkErrorEnum_CustomerManagerLinkError_value = map[string]int32{
	"UNSPECIFIED":                         0,
	"UNKNOWN":                             1,
	"NO_PENDING_INVITE":                   2,
	"SAME_CLIENT_MORE_THAN_ONCE_PER_CALL": 3,
	"MANAGER_HAS_MAX_NUMBER_OF_LINKED_ACCOUNTS":      4,
	"CANNOT_UNLINK_ACCOUNT_WITHOUT_ACTIVE_USER":      5,
	"CANNOT_REMOVE_LAST_CLIENT_ACCOUNT_OWNER":        6,
	"CANNOT_CHANGE_ROLE_BY_NON_ACCOUNT_OWNER":        7,
	"CANNOT_CHANGE_ROLE_FOR_NON_ACTIVE_LINK_ACCOUNT": 8,
	"DUPLICATE_CHILD_FOUND":                          9,
}

func (x CustomerManagerLinkErrorEnum_CustomerManagerLinkError) String() string {
	return proto.EnumName(CustomerManagerLinkErrorEnum_CustomerManagerLinkError_name, int32(x))
}

func (CustomerManagerLinkErrorEnum_CustomerManagerLinkError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96f9efe44c7402db, []int{0, 0}
}

// Container for enum describing possible CustomerManagerLink errors.
type CustomerManagerLinkErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerManagerLinkErrorEnum) Reset()         { *m = CustomerManagerLinkErrorEnum{} }
func (m *CustomerManagerLinkErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CustomerManagerLinkErrorEnum) ProtoMessage()    {}
func (*CustomerManagerLinkErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f9efe44c7402db, []int{0}
}

func (m *CustomerManagerLinkErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerManagerLinkErrorEnum.Unmarshal(m, b)
}
func (m *CustomerManagerLinkErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerManagerLinkErrorEnum.Marshal(b, m, deterministic)
}
func (m *CustomerManagerLinkErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerManagerLinkErrorEnum.Merge(m, src)
}
func (m *CustomerManagerLinkErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerManagerLinkErrorEnum.Size(m)
}
func (m *CustomerManagerLinkErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerManagerLinkErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerManagerLinkErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.CustomerManagerLinkErrorEnum_CustomerManagerLinkError", CustomerManagerLinkErrorEnum_CustomerManagerLinkError_name, CustomerManagerLinkErrorEnum_CustomerManagerLinkError_value)
	proto.RegisterType((*CustomerManagerLinkErrorEnum)(nil), "google.ads.googleads.v2.errors.CustomerManagerLinkErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/customer_manager_link_error.proto", fileDescriptor_96f9efe44c7402db)
}

var fileDescriptor_96f9efe44c7402db = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0x02, 0x2d, 0x6c, 0x0f, 0x18, 0x4b, 0x95, 0x00, 0x95, 0x1e, 0xc2, 0xa1, 0x42,
	0x08, 0x5b, 0x32, 0x37, 0x73, 0x61, 0xb3, 0x9e, 0x24, 0x56, 0xed, 0x59, 0xcb, 0xff, 0x02, 0x28,
	0xd2, 0xc8, 0x34, 0x91, 0x15, 0x35, 0xb1, 0x23, 0x3b, 0xed, 0xf3, 0x20, 0x8e, 0x3c, 0x0a, 0x8f,
	0xc2, 0x13, 0x70, 0x03, 0x39, 0x1b, 0x47, 0x15, 0x22, 0x3d, 0xed, 0x68, 0xe7, 0x37, 0xdf, 0xf7,
	0x49, 0x33, 0xec, 0x63, 0x5e, 0x96, 0xf9, 0x72, 0x6e, 0x66, 0xb3, 0xda, 0x54, 0x65, 0x53, 0xdd,
	0x5a, 0xe6, 0xbc, 0xaa, 0xca, 0xaa, 0x36, 0xaf, 0x6e, 0xea, 0x4d, 0xb9, 0x9a, 0x57, 0xb4, 0xca,
	0x8a, 0x2c, 0x9f, 0x57, 0xb4, 0x5c, 0x14, 0xd7, 0xb4, 0x6d, 0x1a, 0xeb, 0xaa, 0xdc, 0x94, 0xfa,
	0xb9, 0x1a, 0x33, 0xb2, 0x59, 0x6d, 0xec, 0x15, 0x8c, 0x5b, 0xcb, 0x50, 0x0a, 0x2f, 0xcf, 0x5a,
	0x87, 0xf5, 0xc2, 0xcc, 0x8a, 0xa2, 0xdc, 0x64, 0x9b, 0x45, 0x59, 0xd4, 0x6a, 0xba, 0xff, 0xad,
	0xc7, 0xce, 0xc4, 0xce, 0xc3, 0x57, 0x16, 0xde, 0xa2, 0xb8, 0x86, 0x66, 0x16, 0x8a, 0x9b, 0x55,
	0xff, 0x4f, 0x97, 0x3d, 0x3f, 0x04, 0xe8, 0x4f, 0xd9, 0x49, 0x82, 0x51, 0x00, 0xc2, 0x1d, 0xba,
	0xe0, 0x68, 0x0f, 0xf4, 0x13, 0x76, 0x9c, 0xe0, 0x25, 0xca, 0x09, 0x6a, 0x1d, 0xfd, 0x94, 0x3d,
	0x43, 0x49, 0x01, 0xa0, 0xe3, 0xe2, 0x88, 0x5c, 0x4c, 0xdd, 0x18, 0xb4, 0xae, 0x7e, 0xc1, 0x5e,
	0x47, 0xdc, 0x07, 0x12, 0x9e, 0x0b, 0x18, 0x93, 0x2f, 0x43, 0xa0, 0x78, 0xcc, 0x91, 0x24, 0x0a,
	0xa0, 0x00, 0x42, 0x12, 0xdc, 0xf3, 0xb4, 0x9e, 0xfe, 0x8e, 0xbd, 0xf1, 0x39, 0xf2, 0x11, 0x84,
	0x34, 0xe6, 0x11, 0xf9, 0xfc, 0x13, 0x61, 0xe2, 0x0f, 0x20, 0x24, 0x39, 0x24, 0xcf, 0xc5, 0x4b,
	0x70, 0x88, 0x0b, 0x21, 0x13, 0x8c, 0x23, 0xed, 0x61, 0x83, 0x0b, 0x8e, 0x28, 0x63, 0x4a, 0xb0,
	0xe9, 0xb6, 0x3d, 0x9a, 0xb8, 0xf1, 0x58, 0x26, 0x31, 0x71, 0x11, 0xbb, 0x29, 0x50, 0x12, 0x41,
	0xa8, 0x3d, 0xd2, 0xdf, 0xb2, 0x8b, 0x1d, 0x1e, 0x82, 0x2f, 0x53, 0x20, 0x8f, 0x47, 0x71, 0x1b,
	0xaa, 0x1d, 0x95, 0x13, 0x84, 0x50, 0x3b, 0xba, 0x03, 0x8b, 0x31, 0xc7, 0x11, 0x50, 0x28, 0x3d,
	0xa0, 0xc1, 0x67, 0x42, 0x89, 0xff, 0xc0, 0xc7, 0xba, 0xc5, 0x8c, 0xff, 0xc0, 0x43, 0x19, 0xee,
	0xe8, 0x6d, 0x8a, 0xbb, 0x09, 0xb5, 0xc7, 0xfa, 0x0b, 0x76, 0xea, 0x24, 0x81, 0xe7, 0x0a, 0x1e,
	0x03, 0x89, 0xb1, 0xeb, 0x39, 0x34, 0x94, 0x09, 0x3a, 0xda, 0x93, 0xc1, 0xef, 0x0e, 0xeb, 0x5f,
	0x95, 0x2b, 0xe3, 0xfe, 0x3d, 0x0f, 0x5e, 0x1d, 0xda, 0x52, 0xd0, 0x2c, 0x3a, 0xe8, 0x7c, 0x71,
	0x76, 0x02, 0x79, 0xb9, 0xcc, 0x8a, 0xdc, 0x28, 0xab, 0xdc, 0xcc, 0xe7, 0xc5, 0xf6, 0x0c, 0xda,
	0xd3, 0x5b, 0x2f, 0xea, 0x43, 0x97, 0xf8, 0x41, 0x3d, 0xdf, 0xbb, 0xbd, 0x11, 0xe7, 0x3f, 0xba,
	0xe7, 0x23, 0x25, 0xc6, 0x67, 0xb5, 0xa1, 0xca, 0xa6, 0x4a, 0x2d, 0x63, 0x6b, 0x59, 0xff, 0x6c,
	0x81, 0x29, 0x9f, 0xd5, 0xd3, 0x3d, 0x30, 0x4d, 0xad, 0xa9, 0x02, 0x7e, 0x75, 0xfb, 0xea, 0xd7,
	0xb6, 0xf9, 0xac, 0xb6, 0xed, 0x3d, 0x62, 0xdb, 0xa9, 0x65, 0xdb, 0x0a, 0xfa, 0x7a, 0xb4, 0x4d,
	0xf7, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x31, 0xc4, 0x95, 0x26, 0x03, 0x00, 0x00,
}
