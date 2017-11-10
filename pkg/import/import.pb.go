// Code generated by protoc-gen-gogo.
// source: import.proto
// DO NOT EDIT!

/*
	Package importpb is a generated protocol buffer package.

	It is generated from these files:
		import.proto

	It has these top-level messages:
		Mutation
		WriteBatch
		WriteRequest
		WriteResponse
		FlushRequest
		FlushResponse
*/
package importpb

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

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

type Mutation_OP int32

const (
	Mutation_Put Mutation_OP = 0
)

var Mutation_OP_name = map[int32]string{
	0: "Put",
}
var Mutation_OP_value = map[string]int32{
	"Put": 0,
}

func (x Mutation_OP) String() string {
	return proto.EnumName(Mutation_OP_name, int32(x))
}
func (Mutation_OP) EnumDescriptor() ([]byte, []int) { return fileDescriptorImport, []int{0, 0} }

type Mutation struct {
	Op    Mutation_OP `protobuf:"varint,1,opt,name=op,proto3,enum=importpb.Mutation_OP" json:"op,omitempty"`
	Key   []byte      `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte      `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Mutation) Reset()                    { *m = Mutation{} }
func (m *Mutation) String() string            { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()               {}
func (*Mutation) Descriptor() ([]byte, []int) { return fileDescriptorImport, []int{0} }

func (m *Mutation) GetOp() Mutation_OP {
	if m != nil {
		return m.Op
	}
	return Mutation_Put
}

func (m *Mutation) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Mutation) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type WriteBatch struct {
	CommitTs uint64      `protobuf:"varint,1,opt,name=commit_ts,json=commitTs,proto3" json:"commit_ts,omitempty"`
	Mutation []*Mutation `protobuf:"bytes,2,rep,name=mutation" json:"mutation,omitempty"`
}

func (m *WriteBatch) Reset()                    { *m = WriteBatch{} }
func (m *WriteBatch) String() string            { return proto.CompactTextString(m) }
func (*WriteBatch) ProtoMessage()               {}
func (*WriteBatch) Descriptor() ([]byte, []int) { return fileDescriptorImport, []int{1} }

func (m *WriteBatch) GetCommitTs() uint64 {
	if m != nil {
		return m.CommitTs
	}
	return 0
}

func (m *WriteBatch) GetMutation() []*Mutation {
	if m != nil {
		return m.Mutation
	}
	return nil
}

type WriteRequest struct {
	Uuid  []byte      `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Batch *WriteBatch `protobuf:"bytes,2,opt,name=batch" json:"batch,omitempty"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptorImport, []int{2} }

func (m *WriteRequest) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *WriteRequest) GetBatch() *WriteBatch {
	if m != nil {
		return m.Batch
	}
	return nil
}

type WriteResponse struct {
}

func (m *WriteResponse) Reset()                    { *m = WriteResponse{} }
func (m *WriteResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()               {}
func (*WriteResponse) Descriptor() ([]byte, []int) { return fileDescriptorImport, []int{3} }

type FlushRequest struct {
	Uuid []byte `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (m *FlushRequest) Reset()                    { *m = FlushRequest{} }
func (m *FlushRequest) String() string            { return proto.CompactTextString(m) }
func (*FlushRequest) ProtoMessage()               {}
func (*FlushRequest) Descriptor() ([]byte, []int) { return fileDescriptorImport, []int{4} }

func (m *FlushRequest) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

type FlushResponse struct {
}

func (m *FlushResponse) Reset()                    { *m = FlushResponse{} }
func (m *FlushResponse) String() string            { return proto.CompactTextString(m) }
func (*FlushResponse) ProtoMessage()               {}
func (*FlushResponse) Descriptor() ([]byte, []int) { return fileDescriptorImport, []int{5} }

func init() {
	proto.RegisterType((*Mutation)(nil), "importpb.Mutation")
	proto.RegisterType((*WriteBatch)(nil), "importpb.WriteBatch")
	proto.RegisterType((*WriteRequest)(nil), "importpb.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "importpb.WriteResponse")
	proto.RegisterType((*FlushRequest)(nil), "importpb.FlushRequest")
	proto.RegisterType((*FlushResponse)(nil), "importpb.FlushResponse")
	proto.RegisterEnum("importpb.Mutation_OP", Mutation_OP_name, Mutation_OP_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Import service

type ImportClient interface {
	Write(ctx context.Context, opts ...grpc.CallOption) (Import_WriteClient, error)
	Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error)
}

type importClient struct {
	cc *grpc.ClientConn
}

func NewImportClient(cc *grpc.ClientConn) ImportClient {
	return &importClient{cc}
}

func (c *importClient) Write(ctx context.Context, opts ...grpc.CallOption) (Import_WriteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Import_serviceDesc.Streams[0], c.cc, "/importpb.Import/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &importWriteClient{stream}
	return x, nil
}

type Import_WriteClient interface {
	Send(*WriteRequest) error
	CloseAndRecv() (*WriteResponse, error)
	grpc.ClientStream
}

type importWriteClient struct {
	grpc.ClientStream
}

func (x *importWriteClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *importWriteClient) CloseAndRecv() (*WriteResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *importClient) Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error) {
	out := new(FlushResponse)
	err := grpc.Invoke(ctx, "/importpb.Import/Flush", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Import service

type ImportServer interface {
	Write(Import_WriteServer) error
	Flush(context.Context, *FlushRequest) (*FlushResponse, error)
}

func RegisterImportServer(s *grpc.Server, srv ImportServer) {
	s.RegisterService(&_Import_serviceDesc, srv)
}

func _Import_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImportServer).Write(&importWriteServer{stream})
}

type Import_WriteServer interface {
	SendAndClose(*WriteResponse) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type importWriteServer struct {
	grpc.ServerStream
}

func (x *importWriteServer) SendAndClose(m *WriteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *importWriteServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Import_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/importpb.Import/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportServer).Flush(ctx, req.(*FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Import_serviceDesc = grpc.ServiceDesc{
	ServiceName: "importpb.Import",
	HandlerType: (*ImportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Flush",
			Handler:    _Import_Flush_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Write",
			Handler:       _Import_Write_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "import.proto",
}

func (m *Mutation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mutation) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Op != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintImport(dAtA, i, uint64(m.Op))
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintImport(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintImport(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *WriteBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteBatch) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CommitTs != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintImport(dAtA, i, uint64(m.CommitTs))
	}
	if len(m.Mutation) > 0 {
		for _, msg := range m.Mutation {
			dAtA[i] = 0x12
			i++
			i = encodeVarintImport(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *WriteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uuid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintImport(dAtA, i, uint64(len(m.Uuid)))
		i += copy(dAtA[i:], m.Uuid)
	}
	if m.Batch != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintImport(dAtA, i, uint64(m.Batch.Size()))
		n1, err := m.Batch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *WriteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *FlushRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FlushRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uuid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintImport(dAtA, i, uint64(len(m.Uuid)))
		i += copy(dAtA[i:], m.Uuid)
	}
	return i, nil
}

func (m *FlushResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FlushResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64Import(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Import(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintImport(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Mutation) Size() (n int) {
	var l int
	_ = l
	if m.Op != 0 {
		n += 1 + sovImport(uint64(m.Op))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovImport(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovImport(uint64(l))
	}
	return n
}

func (m *WriteBatch) Size() (n int) {
	var l int
	_ = l
	if m.CommitTs != 0 {
		n += 1 + sovImport(uint64(m.CommitTs))
	}
	if len(m.Mutation) > 0 {
		for _, e := range m.Mutation {
			l = e.Size()
			n += 1 + l + sovImport(uint64(l))
		}
	}
	return n
}

func (m *WriteRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovImport(uint64(l))
	}
	if m.Batch != nil {
		l = m.Batch.Size()
		n += 1 + l + sovImport(uint64(l))
	}
	return n
}

func (m *WriteResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *FlushRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovImport(uint64(l))
	}
	return n
}

func (m *FlushResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovImport(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozImport(x uint64) (n int) {
	return sovImport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mutation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Mutation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mutation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= (Mutation_OP(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthImport
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthImport
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImport
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
func (m *WriteBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WriteBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitTs", wireType)
			}
			m.CommitTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitTs |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mutation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthImport
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mutation = append(m.Mutation, &Mutation{})
			if err := m.Mutation[len(m.Mutation)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImport
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
func (m *WriteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WriteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthImport
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = append(m.Uuid[:0], dAtA[iNdEx:postIndex]...)
			if m.Uuid == nil {
				m.Uuid = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthImport
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Batch == nil {
				m.Batch = &WriteBatch{}
			}
			if err := m.Batch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImport
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
func (m *WriteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WriteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipImport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImport
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
func (m *FlushRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FlushRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FlushRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthImport
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = append(m.Uuid[:0], dAtA[iNdEx:postIndex]...)
			if m.Uuid == nil {
				m.Uuid = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipImport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImport
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
func (m *FlushResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FlushResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FlushResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipImport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImport
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
func skipImport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImport
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
					return 0, ErrIntOverflowImport
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
					return 0, ErrIntOverflowImport
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthImport
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowImport
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
				next, err := skipImport(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthImport = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImport   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("import.proto", fileDescriptorImport) }

var fileDescriptorImport = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xb3, 0x49, 0x53, 0xe3, 0x34, 0xd5, 0x32, 0x54, 0x0d, 0x0a, 0xa1, 0x2c, 0x08, 0xc1,
	0x43, 0x0e, 0xf1, 0x26, 0x9e, 0x7a, 0x10, 0x3c, 0x68, 0xcb, 0x22, 0x88, 0x27, 0x69, 0xeb, 0x42,
	0x83, 0x4d, 0x77, 0xed, 0xee, 0x0a, 0x5e, 0x7d, 0x0a, 0x1f, 0xc9, 0xa3, 0x8f, 0x20, 0xf5, 0x45,
	0xa4, 0x9b, 0xd4, 0x56, 0x0b, 0xde, 0x66, 0x66, 0xff, 0xff, 0xff, 0xfe, 0xc3, 0x42, 0x98, 0x17,
	0x52, 0xcc, 0x74, 0x2a, 0x67, 0x42, 0x0b, 0x0c, 0xca, 0x4d, 0x0e, 0xe9, 0x18, 0x82, 0x2b, 0xa3,
	0x07, 0x3a, 0x17, 0x53, 0x3c, 0x06, 0x57, 0xc8, 0x88, 0x74, 0x48, 0xb2, 0x93, 0xed, 0xa5, 0x4b,
	0x49, 0xba, 0x7c, 0x4f, 0x7b, 0x7d, 0xe6, 0x0a, 0x89, 0x2d, 0xf0, 0x1e, 0xf9, 0x4b, 0xe4, 0x76,
	0x48, 0x12, 0xb2, 0xc5, 0x88, 0x6d, 0xf0, 0x9f, 0x07, 0x13, 0xc3, 0x23, 0xcf, 0xde, 0xca, 0x85,
	0x36, 0xc1, 0xed, 0xf5, 0x71, 0x0b, 0xbc, 0xbe, 0xd1, 0x2d, 0x87, 0xde, 0x01, 0xdc, 0xce, 0x72,
	0xcd, 0xbb, 0x03, 0x3d, 0x1a, 0xe3, 0x11, 0x6c, 0x8f, 0x44, 0x51, 0xe4, 0xfa, 0x5e, 0x2b, 0x8b,
	0xac, 0xb1, 0xa0, 0x3c, 0xdc, 0x28, 0x4c, 0x21, 0x28, 0x2a, 0x68, 0xe4, 0x76, 0xbc, 0xa4, 0x91,
	0xe1, 0x66, 0x1d, 0xf6, 0xa3, 0xa1, 0xd7, 0x10, 0xda, 0x68, 0xc6, 0x9f, 0x0c, 0x57, 0x1a, 0x11,
	0x6a, 0xc6, 0xe4, 0x0f, 0x36, 0x37, 0x64, 0x76, 0xc6, 0x13, 0xf0, 0x87, 0x0b, 0xb2, 0xed, 0xdd,
	0xc8, 0xda, 0xab, 0xc0, 0x55, 0x2b, 0x56, 0x4a, 0xe8, 0x2e, 0x34, 0xab, 0x3c, 0x25, 0xc5, 0x54,
	0x71, 0x4a, 0x21, 0xbc, 0x98, 0x18, 0x35, 0xfe, 0x07, 0xb0, 0x30, 0x55, 0x9a, 0xd2, 0x94, 0xbd,
	0x12, 0xa8, 0x5f, 0x5a, 0x08, 0x9e, 0x83, 0x6f, 0x03, 0x71, 0xff, 0x0f, 0xb6, 0x0a, 0x3c, 0x3c,
	0xd8, 0xb8, 0x57, 0x64, 0x27, 0x21, 0x78, 0x06, 0xbe, 0x4d, 0x5e, 0x77, 0xaf, 0xd7, 0x59, 0x77,
	0xff, 0xaa, 0x40, 0x9d, 0x6e, 0xeb, 0x7d, 0x1e, 0x93, 0x8f, 0x79, 0x4c, 0x3e, 0xe7, 0x31, 0x79,
	0xfb, 0x8a, 0x9d, 0x61, 0xdd, 0x7e, 0x81, 0xd3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x81,
	0x3d, 0x90, 0x12, 0x02, 0x00, 0x00,
}