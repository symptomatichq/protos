// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: authentication/authentication.proto

package authentication

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AuthenticateUserRequest struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *AuthenticateUserRequest) Reset()         { *m = AuthenticateUserRequest{} }
func (m *AuthenticateUserRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateUserRequest) ProtoMessage()    {}
func (*AuthenticateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_067dc6fa29e56c89, []int{0}
}
func (m *AuthenticateUserRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticateUserRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthenticateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateUserRequest.Merge(m, src)
}
func (m *AuthenticateUserRequest) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateUserRequest proto.InternalMessageInfo

func (m *AuthenticateUserRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type AuthenticateUserResponse struct {
	ID            string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Authenticated bool       `protobuf:"varint,2,opt,name=authenticated,proto3" json:"authenticated,omitempty"`
	LastLogin     *time.Time `protobuf:"bytes,3,opt,name=last_login,json=lastLogin,proto3,stdtime" json:"last_login,omitempty"`
}

func (m *AuthenticateUserResponse) Reset()         { *m = AuthenticateUserResponse{} }
func (m *AuthenticateUserResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateUserResponse) ProtoMessage()    {}
func (*AuthenticateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_067dc6fa29e56c89, []int{1}
}
func (m *AuthenticateUserResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticateUserResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthenticateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateUserResponse.Merge(m, src)
}
func (m *AuthenticateUserResponse) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateUserResponse proto.InternalMessageInfo

func (m *AuthenticateUserResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AuthenticateUserResponse) GetAuthenticated() bool {
	if m != nil {
		return m.Authenticated
	}
	return false
}

func (m *AuthenticateUserResponse) GetLastLogin() *time.Time {
	if m != nil {
		return m.LastLogin
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthenticateUserRequest)(nil), "authentication.AuthenticateUserRequest")
	proto.RegisterType((*AuthenticateUserResponse)(nil), "authentication.AuthenticateUserResponse")
}

func init() {
	proto.RegisterFile("authentication/authentication.proto", fileDescriptor_067dc6fa29e56c89)
}

var fileDescriptor_067dc6fa29e56c89 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2c, 0x2d, 0xc9,
	0x48, 0xcd, 0x2b, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x47, 0xe5, 0xea, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa1, 0x8a, 0x4a, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x95, 0x25, 0x95, 0xa6, 0x81, 0x79,
	0x60, 0x0e, 0x98, 0x05, 0xd1, 0x2e, 0x25, 0x9f, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0x8a, 0x50, 0x55,
	0x92, 0x99, 0x9b, 0x5a, 0x5c, 0x92, 0x98, 0x5b, 0x00, 0x51, 0xa0, 0x64, 0xc8, 0x25, 0xee, 0x88,
	0xb0, 0x21, 0x35, 0xb4, 0x38, 0xb5, 0x28, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8c,
	0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x89, 0xed, 0xd1, 0x3d, 0x79, 0x26,
	0x4f, 0x97, 0x20, 0xa6, 0xcc, 0x14, 0xa5, 0x99, 0x8c, 0x5c, 0x12, 0x98, 0x7a, 0x8a, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0x71, 0x69, 0x12, 0x52, 0xe1, 0xe2, 0x45, 0xf2, 0x49, 0x6a, 0x8a, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0x47, 0x10, 0xaa, 0xa0, 0x90, 0x3d, 0x17, 0x57, 0x4e, 0x62, 0x71, 0x49, 0x7c,
	0x4e, 0x7e, 0x7a, 0x66, 0x9e, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x94, 0x1e, 0xc4, 0x0f,
	0x7a, 0x30, 0x3f, 0xe8, 0x85, 0xc0, 0xfc, 0xe0, 0xc4, 0x32, 0xe1, 0xbe, 0x3c, 0x63, 0x10, 0x27,
	0x48, 0x8f, 0x0f, 0x48, 0x8b, 0x51, 0x25, 0x17, 0x9f, 0x23, 0x4a, 0x80, 0x09, 0xa5, 0x73, 0x09,
	0xa0, 0x3b, 0x56, 0x48, 0x5d, 0x0f, 0x2d, 0xac, 0x71, 0x04, 0x81, 0x94, 0x06, 0x61, 0x85, 0x10,
	0x7f, 0x2b, 0x31, 0x38, 0x69, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x14,
	0x5a, 0x1c, 0x26, 0xb1, 0x81, 0x7d, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x99, 0xce, 0x27,
	0xf7, 0x01, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationClient interface {
	AuthenticateUser(ctx context.Context, in *AuthenticateUserRequest, opts ...grpc.CallOption) (*AuthenticateUserResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) AuthenticateUser(ctx context.Context, in *AuthenticateUserRequest, opts ...grpc.CallOption) (*AuthenticateUserResponse, error) {
	out := new(AuthenticateUserResponse)
	err := c.cc.Invoke(ctx, "/authentication.Authentication/AuthenticateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
type AuthenticationServer interface {
	AuthenticateUser(context.Context, *AuthenticateUserRequest) (*AuthenticateUserResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_AuthenticateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).AuthenticateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/AuthenticateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).AuthenticateUser(ctx, req.(*AuthenticateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticateUser",
			Handler:    _Authentication_AuthenticateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication/authentication.proto",
}

func (m *AuthenticateUserRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticateUserRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuthentication(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	return i, nil
}

func (m *AuthenticateUserResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticateUserResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuthentication(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Authenticated {
		dAtA[i] = 0x10
		i++
		if m.Authenticated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.LastLogin != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAuthentication(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastLogin)))
		n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.LastLogin, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintAuthentication(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AuthenticateUserRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovAuthentication(uint64(l))
	}
	return n
}

func (m *AuthenticateUserResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovAuthentication(uint64(l))
	}
	if m.Authenticated {
		n += 2
	}
	if m.LastLogin != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastLogin)
		n += 1 + l + sovAuthentication(uint64(l))
	}
	return n
}

func sovAuthentication(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAuthentication(x uint64) (n int) {
	return sovAuthentication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthenticateUserRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthenticateUserRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateUserRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuthentication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthentication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuthentication
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthenticateUserResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthenticateUserResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateUserResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuthentication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthentication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authenticated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Authenticated = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastLogin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuthentication
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthentication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastLogin == nil {
				m.LastLogin = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.LastLogin, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuthentication
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAuthentication(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthentication
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuthentication
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuthentication
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAuthentication
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAuthentication
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuthentication
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAuthentication(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAuthentication
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAuthentication = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthentication   = fmt.Errorf("proto: integer overflow")
)