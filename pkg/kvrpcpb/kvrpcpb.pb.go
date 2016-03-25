// Code generated by protoc-gen-go.
// source: kvrpcpb.proto
// DO NOT EDIT!

/*
Package kvrpcpb is a generated protocol buffer package.

It is generated from these files:
	kvrpcpb.proto

It has these top-level messages:
	LockInfo
	ResultType
	KeyAddress
	CmdGetRequest
	CmdGetResponse
	CmdScanRequest
	Item
	CmdScanResponse
	Mutation
	CmdPrewriteRequest
	CmdPrewriteResponse
	CmdCommitRequest
	CmdCommitResponse
	CmdCleanupRequest
	CmdCleanupResponse
	CmdRollbackThenGetRequest
	CmdRollbackThenGetResponse
	CmdCommitThenGetRequest
	CmdCommitThenGetResponse
	Request
	Response
*/
package kvrpcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import metapb "github.com/pingcap/kvproto/pkg/metapb"
import errorpb "github.com/pingcap/kvproto/pkg/errorpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type MessageType int32

const (
	MessageType_CmdGet      MessageType = 1
	MessageType_CmdScan     MessageType = 2
	MessageType_CmdPrewrite MessageType = 3
	MessageType_CmdCommit   MessageType = 4
	MessageType_CmdCleanup  MessageType = 5
	// Below types both use for Get failed. If Get failed, it may be locked.
	// So it tries to clean primary lock(CmdCleanup), and then server will return
	// either committed or rolled back. Finally, client will commit/rollback
	// primary lock and THEN Get again.
	MessageType_CmdRollbackThenGet MessageType = 6
	MessageType_CmdCommitThenGet   MessageType = 7
)

var MessageType_name = map[int32]string{
	1: "CmdGet",
	2: "CmdScan",
	3: "CmdPrewrite",
	4: "CmdCommit",
	5: "CmdCleanup",
	6: "CmdRollbackThenGet",
	7: "CmdCommitThenGet",
}
var MessageType_value = map[string]int32{
	"CmdGet":             1,
	"CmdScan":            2,
	"CmdPrewrite":        3,
	"CmdCommit":          4,
	"CmdCleanup":         5,
	"CmdRollbackThenGet": 6,
	"CmdCommitThenGet":   7,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Operator int32

const (
	Operator_OpPut  Operator = 1
	Operator_OpDel  Operator = 2
	Operator_OpLock Operator = 3
)

var Operator_name = map[int32]string{
	1: "OpPut",
	2: "OpDel",
	3: "OpLock",
}
var Operator_value = map[string]int32{
	"OpPut":  1,
	"OpDel":  2,
	"OpLock": 3,
}

func (x Operator) Enum() *Operator {
	p := new(Operator)
	*p = x
	return p
}
func (x Operator) String() string {
	return proto.EnumName(Operator_name, int32(x))
}
func (x *Operator) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Operator_value, data, "Operator")
	if err != nil {
		return err
	}
	*x = Operator(value)
	return nil
}
func (Operator) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ResultType_Type int32

const (
	ResultType_Ok         ResultType_Type = 1
	ResultType_Retryable  ResultType_Type = 2
	ResultType_Locked     ResultType_Type = 3
	ResultType_Committed  ResultType_Type = 4
	ResultType_Rolledback ResultType_Type = 5
	ResultType_NotLeader  ResultType_Type = 6
	// Known result type add here.
	ResultType_Other ResultType_Type = 9
)

var ResultType_Type_name = map[int32]string{
	1: "Ok",
	2: "Retryable",
	3: "Locked",
	4: "Committed",
	5: "Rolledback",
	6: "NotLeader",
	9: "Other",
}
var ResultType_Type_value = map[string]int32{
	"Ok":         1,
	"Retryable":  2,
	"Locked":     3,
	"Committed":  4,
	"Rolledback": 5,
	"NotLeader":  6,
	"Other":      9,
}

func (x ResultType_Type) Enum() *ResultType_Type {
	p := new(ResultType_Type)
	*p = x
	return p
}
func (x ResultType_Type) String() string {
	return proto.EnumName(ResultType_Type_name, int32(x))
}
func (x *ResultType_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ResultType_Type_value, data, "ResultType_Type")
	if err != nil {
		return err
	}
	*x = ResultType_Type(value)
	return nil
}
func (ResultType_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type LockInfo struct {
	PrimaryLock      []byte  `protobuf:"bytes,1,opt,name=primary_lock" json:"primary_lock,omitempty"`
	LockVersion      *uint64 `protobuf:"varint,2,opt,name=lock_version" json:"lock_version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LockInfo) Reset()                    { *m = LockInfo{} }
func (m *LockInfo) String() string            { return proto.CompactTextString(m) }
func (*LockInfo) ProtoMessage()               {}
func (*LockInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LockInfo) GetPrimaryLock() []byte {
	if m != nil {
		return m.PrimaryLock
	}
	return nil
}

func (m *LockInfo) GetLockVersion() uint64 {
	if m != nil && m.LockVersion != nil {
		return *m.LockVersion
	}
	return 0
}

// ResultType use for many different result, so dispose each Type prudently.
type ResultType struct {
	Type *ResultType_Type `protobuf:"varint,1,opt,name=type,enum=kvrpcpb.ResultType_Type,def=9" json:"type,omitempty"`
	Msg  *string          `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	// If and only if Type == Locked
	LockInfo         *LockInfo               `protobuf:"bytes,3,opt,name=lock_info" json:"lock_info,omitempty"`
	LeaderInfo       *errorpb.NotLeaderError `protobuf:"bytes,4,opt,name=leader_info" json:"leader_info,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ResultType) Reset()                    { *m = ResultType{} }
func (m *ResultType) String() string            { return proto.CompactTextString(m) }
func (*ResultType) ProtoMessage()               {}
func (*ResultType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_ResultType_Type ResultType_Type = ResultType_Other

func (m *ResultType) GetType() ResultType_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_ResultType_Type
}

func (m *ResultType) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *ResultType) GetLockInfo() *LockInfo {
	if m != nil {
		return m.LockInfo
	}
	return nil
}

func (m *ResultType) GetLeaderInfo() *errorpb.NotLeaderError {
	if m != nil {
		return m.LeaderInfo
	}
	return nil
}

type KeyAddress struct {
	Key              []byte       `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	RegionId         *uint64      `protobuf:"varint,2,opt,name=region_id" json:"region_id,omitempty"`
	Peer             *metapb.Peer `protobuf:"bytes,3,opt,name=peer" json:"peer,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *KeyAddress) Reset()                    { *m = KeyAddress{} }
func (m *KeyAddress) String() string            { return proto.CompactTextString(m) }
func (*KeyAddress) ProtoMessage()               {}
func (*KeyAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KeyAddress) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyAddress) GetRegionId() uint64 {
	if m != nil && m.RegionId != nil {
		return *m.RegionId
	}
	return 0
}

func (m *KeyAddress) GetPeer() *metapb.Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

type CmdGetRequest struct {
	KeyAddress       *KeyAddress `protobuf:"bytes,1,opt,name=key_address" json:"key_address,omitempty"`
	Version          *uint64     `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdGetRequest) Reset()                    { *m = CmdGetRequest{} }
func (m *CmdGetRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdGetRequest) ProtoMessage()               {}
func (*CmdGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CmdGetRequest) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

func (m *CmdGetRequest) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type CmdGetResponse struct {
	ResType          *ResultType `protobuf:"bytes,1,opt,name=res_type" json:"res_type,omitempty"`
	Value            []byte      `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdGetResponse) Reset()                    { *m = CmdGetResponse{} }
func (m *CmdGetResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdGetResponse) ProtoMessage()               {}
func (*CmdGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CmdGetResponse) GetResType() *ResultType {
	if m != nil {
		return m.ResType
	}
	return nil
}

func (m *CmdGetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type CmdScanRequest struct {
	KeyAddress       *KeyAddress `protobuf:"bytes,1,opt,name=key_address" json:"key_address,omitempty"`
	Limit            *uint32     `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Version          *uint64     `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdScanRequest) Reset()                    { *m = CmdScanRequest{} }
func (m *CmdScanRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdScanRequest) ProtoMessage()               {}
func (*CmdScanRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CmdScanRequest) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

func (m *CmdScanRequest) GetLimit() uint32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *CmdScanRequest) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type Item struct {
	ResType          *ResultType `protobuf:"bytes,1,opt,name=res_type" json:"res_type,omitempty"`
	Key              []byte      `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value            []byte      `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Item) GetResType() *ResultType {
	if m != nil {
		return m.ResType
	}
	return nil
}

func (m *Item) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Item) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type CmdScanResponse struct {
	// ok if !ok then retry.
	Ok               *bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Results          []*Item `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CmdScanResponse) Reset()                    { *m = CmdScanResponse{} }
func (m *CmdScanResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdScanResponse) ProtoMessage()               {}
func (*CmdScanResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CmdScanResponse) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func (m *CmdScanResponse) GetResults() []*Item {
	if m != nil {
		return m.Results
	}
	return nil
}

type Mutation struct {
	Op               *Operator `protobuf:"varint,1,opt,name=op,enum=kvrpcpb.Operator" json:"op,omitempty"`
	Key              []byte    `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value            []byte    `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Mutation) Reset()                    { *m = Mutation{} }
func (m *Mutation) String() string            { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()               {}
func (*Mutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Mutation) GetOp() Operator {
	if m != nil && m.Op != nil {
		return *m.Op
	}
	return Operator_OpPut
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

type CmdPrewriteRequest struct {
	Mutations []*Mutation `protobuf:"bytes,1,rep,name=mutations" json:"mutations,omitempty"`
	// primary_lock_key
	PrimaryLock      []byte      `protobuf:"bytes,2,opt,name=primary_lock" json:"primary_lock,omitempty"`
	StartVersion     *uint64     `protobuf:"varint,3,opt,name=start_version" json:"start_version,omitempty"`
	KeyAddress       *KeyAddress `protobuf:"bytes,4,opt,name=key_address" json:"key_address,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdPrewriteRequest) Reset()                    { *m = CmdPrewriteRequest{} }
func (m *CmdPrewriteRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdPrewriteRequest) ProtoMessage()               {}
func (*CmdPrewriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CmdPrewriteRequest) GetMutations() []*Mutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

func (m *CmdPrewriteRequest) GetPrimaryLock() []byte {
	if m != nil {
		return m.PrimaryLock
	}
	return nil
}

func (m *CmdPrewriteRequest) GetStartVersion() uint64 {
	if m != nil && m.StartVersion != nil {
		return *m.StartVersion
	}
	return 0
}

func (m *CmdPrewriteRequest) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

type CmdPrewriteResponse struct {
	Ok *bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	// This Item doesn't contain value = 3
	Results          []*Item `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CmdPrewriteResponse) Reset()                    { *m = CmdPrewriteResponse{} }
func (m *CmdPrewriteResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdPrewriteResponse) ProtoMessage()               {}
func (*CmdPrewriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CmdPrewriteResponse) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func (m *CmdPrewriteResponse) GetResults() []*Item {
	if m != nil {
		return m.Results
	}
	return nil
}

type CmdCommitRequest struct {
	StartVersion     *uint64       `protobuf:"varint,1,opt,name=start_version" json:"start_version,omitempty"`
	KeysAddress      []*KeyAddress `protobuf:"bytes,2,rep,name=keys_address" json:"keys_address,omitempty"`
	CommitVersion    *uint64       `protobuf:"varint,3,opt,name=commit_version" json:"commit_version,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *CmdCommitRequest) Reset()                    { *m = CmdCommitRequest{} }
func (m *CmdCommitRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdCommitRequest) ProtoMessage()               {}
func (*CmdCommitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CmdCommitRequest) GetStartVersion() uint64 {
	if m != nil && m.StartVersion != nil {
		return *m.StartVersion
	}
	return 0
}

func (m *CmdCommitRequest) GetKeysAddress() []*KeyAddress {
	if m != nil {
		return m.KeysAddress
	}
	return nil
}

func (m *CmdCommitRequest) GetCommitVersion() uint64 {
	if m != nil && m.CommitVersion != nil {
		return *m.CommitVersion
	}
	return 0
}

type CmdCommitResponse struct {
	Ok               *bool  `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CmdCommitResponse) Reset()                    { *m = CmdCommitResponse{} }
func (m *CmdCommitResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdCommitResponse) ProtoMessage()               {}
func (*CmdCommitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CmdCommitResponse) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

type CmdCleanupRequest struct {
	KeyAddress       *KeyAddress `protobuf:"bytes,1,opt,name=key_address" json:"key_address,omitempty"`
	StartVersion     *uint64     `protobuf:"varint,2,opt,name=start_version" json:"start_version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdCleanupRequest) Reset()                    { *m = CmdCleanupRequest{} }
func (m *CmdCleanupRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdCleanupRequest) ProtoMessage()               {}
func (*CmdCleanupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *CmdCleanupRequest) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

func (m *CmdCleanupRequest) GetStartVersion() uint64 {
	if m != nil && m.StartVersion != nil {
		return *m.StartVersion
	}
	return 0
}

type CmdCleanupResponse struct {
	ResType          *ResultType `protobuf:"bytes,1,opt,name=res_type" json:"res_type,omitempty"`
	CommitVersion    *uint64     `protobuf:"varint,2,opt,name=commit_version" json:"commit_version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdCleanupResponse) Reset()                    { *m = CmdCleanupResponse{} }
func (m *CmdCleanupResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdCleanupResponse) ProtoMessage()               {}
func (*CmdCleanupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *CmdCleanupResponse) GetResType() *ResultType {
	if m != nil {
		return m.ResType
	}
	return nil
}

func (m *CmdCleanupResponse) GetCommitVersion() uint64 {
	if m != nil && m.CommitVersion != nil {
		return *m.CommitVersion
	}
	return 0
}

type CmdRollbackThenGetRequest struct {
	KeyAddress       *KeyAddress `protobuf:"bytes,1,opt,name=key_address" json:"key_address,omitempty"`
	LockVersion      *uint64     `protobuf:"varint,2,opt,name=lock_version" json:"lock_version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdRollbackThenGetRequest) Reset()                    { *m = CmdRollbackThenGetRequest{} }
func (m *CmdRollbackThenGetRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdRollbackThenGetRequest) ProtoMessage()               {}
func (*CmdRollbackThenGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *CmdRollbackThenGetRequest) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

func (m *CmdRollbackThenGetRequest) GetLockVersion() uint64 {
	if m != nil && m.LockVersion != nil {
		return *m.LockVersion
	}
	return 0
}

type CmdRollbackThenGetResponse struct {
	Ok               *bool  `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CmdRollbackThenGetResponse) Reset()                    { *m = CmdRollbackThenGetResponse{} }
func (m *CmdRollbackThenGetResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdRollbackThenGetResponse) ProtoMessage()               {}
func (*CmdRollbackThenGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *CmdRollbackThenGetResponse) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func (m *CmdRollbackThenGetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type CmdCommitThenGetRequest struct {
	KeyAddress       *KeyAddress `protobuf:"bytes,1,opt,name=key_address" json:"key_address,omitempty"`
	LockVersion      *uint64     `protobuf:"varint,2,opt,name=lock_version" json:"lock_version,omitempty"`
	CommitVersion    *uint64     `protobuf:"varint,3,opt,name=commit_version" json:"commit_version,omitempty"`
	GetVersion       *uint64     `protobuf:"varint,4,opt,name=get_version" json:"get_version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdCommitThenGetRequest) Reset()                    { *m = CmdCommitThenGetRequest{} }
func (m *CmdCommitThenGetRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdCommitThenGetRequest) ProtoMessage()               {}
func (*CmdCommitThenGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *CmdCommitThenGetRequest) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

func (m *CmdCommitThenGetRequest) GetLockVersion() uint64 {
	if m != nil && m.LockVersion != nil {
		return *m.LockVersion
	}
	return 0
}

func (m *CmdCommitThenGetRequest) GetCommitVersion() uint64 {
	if m != nil && m.CommitVersion != nil {
		return *m.CommitVersion
	}
	return 0
}

func (m *CmdCommitThenGetRequest) GetGetVersion() uint64 {
	if m != nil && m.GetVersion != nil {
		return *m.GetVersion
	}
	return 0
}

type CmdCommitThenGetResponse struct {
	Ok               *bool  `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CmdCommitThenGetResponse) Reset()                    { *m = CmdCommitThenGetResponse{} }
func (m *CmdCommitThenGetResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdCommitThenGetResponse) ProtoMessage()               {}
func (*CmdCommitThenGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *CmdCommitThenGetResponse) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func (m *CmdCommitThenGetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Request struct {
	Type             *MessageType               `protobuf:"varint,1,opt,name=type,enum=kvrpcpb.MessageType" json:"type,omitempty"`
	CmdGetReq        *CmdGetRequest             `protobuf:"bytes,2,opt,name=cmd_get_req" json:"cmd_get_req,omitempty"`
	CmdScanReq       *CmdScanRequest            `protobuf:"bytes,3,opt,name=cmd_scan_req" json:"cmd_scan_req,omitempty"`
	CmdPrewriteReq   *CmdPrewriteRequest        `protobuf:"bytes,4,opt,name=cmd_prewrite_req" json:"cmd_prewrite_req,omitempty"`
	CmdCommitReq     *CmdCommitRequest          `protobuf:"bytes,5,opt,name=cmd_commit_req" json:"cmd_commit_req,omitempty"`
	CmdCleanupReq    *CmdCleanupRequest         `protobuf:"bytes,6,opt,name=cmd_cleanup_req" json:"cmd_cleanup_req,omitempty"`
	CmdRbGetReq      *CmdRollbackThenGetRequest `protobuf:"bytes,7,opt,name=cmd_rb_get_req" json:"cmd_rb_get_req,omitempty"`
	CmdCommitGetReq  *CmdCommitThenGetRequest   `protobuf:"bytes,8,opt,name=cmd_commit_get_req" json:"cmd_commit_get_req,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *Request) GetType() MessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MessageType_CmdGet
}

func (m *Request) GetCmdGetReq() *CmdGetRequest {
	if m != nil {
		return m.CmdGetReq
	}
	return nil
}

func (m *Request) GetCmdScanReq() *CmdScanRequest {
	if m != nil {
		return m.CmdScanReq
	}
	return nil
}

func (m *Request) GetCmdPrewriteReq() *CmdPrewriteRequest {
	if m != nil {
		return m.CmdPrewriteReq
	}
	return nil
}

func (m *Request) GetCmdCommitReq() *CmdCommitRequest {
	if m != nil {
		return m.CmdCommitReq
	}
	return nil
}

func (m *Request) GetCmdCleanupReq() *CmdCleanupRequest {
	if m != nil {
		return m.CmdCleanupReq
	}
	return nil
}

func (m *Request) GetCmdRbGetReq() *CmdRollbackThenGetRequest {
	if m != nil {
		return m.CmdRbGetReq
	}
	return nil
}

func (m *Request) GetCmdCommitGetReq() *CmdCommitThenGetRequest {
	if m != nil {
		return m.CmdCommitGetReq
	}
	return nil
}

type Response struct {
	Type             *MessageType                `protobuf:"varint,1,opt,name=type,enum=kvrpcpb.MessageType" json:"type,omitempty"`
	CmdGetResp       *CmdGetResponse             `protobuf:"bytes,2,opt,name=cmd_get_resp" json:"cmd_get_resp,omitempty"`
	CmdScanResp      *CmdScanResponse            `protobuf:"bytes,3,opt,name=cmd_scan_resp" json:"cmd_scan_resp,omitempty"`
	CmdPrewriteResp  *CmdPrewriteResponse        `protobuf:"bytes,4,opt,name=cmd_prewrite_resp" json:"cmd_prewrite_resp,omitempty"`
	CmdCommitResp    *CmdCommitResponse          `protobuf:"bytes,5,opt,name=cmd_commit_resp" json:"cmd_commit_resp,omitempty"`
	CmdCleanupResp   *CmdCleanupResponse         `protobuf:"bytes,6,opt,name=cmd_cleanup_resp" json:"cmd_cleanup_resp,omitempty"`
	CmdRbGetResp     *CmdRollbackThenGetResponse `protobuf:"bytes,7,opt,name=cmd_rb_get_resp" json:"cmd_rb_get_resp,omitempty"`
	CmdCommitGetResp *CmdCommitThenGetResponse   `protobuf:"bytes,8,opt,name=cmd_commit_get_resp" json:"cmd_commit_get_resp,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *Response) GetType() MessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MessageType_CmdGet
}

func (m *Response) GetCmdGetResp() *CmdGetResponse {
	if m != nil {
		return m.CmdGetResp
	}
	return nil
}

func (m *Response) GetCmdScanResp() *CmdScanResponse {
	if m != nil {
		return m.CmdScanResp
	}
	return nil
}

func (m *Response) GetCmdPrewriteResp() *CmdPrewriteResponse {
	if m != nil {
		return m.CmdPrewriteResp
	}
	return nil
}

func (m *Response) GetCmdCommitResp() *CmdCommitResponse {
	if m != nil {
		return m.CmdCommitResp
	}
	return nil
}

func (m *Response) GetCmdCleanupResp() *CmdCleanupResponse {
	if m != nil {
		return m.CmdCleanupResp
	}
	return nil
}

func (m *Response) GetCmdRbGetResp() *CmdRollbackThenGetResponse {
	if m != nil {
		return m.CmdRbGetResp
	}
	return nil
}

func (m *Response) GetCmdCommitGetResp() *CmdCommitThenGetResponse {
	if m != nil {
		return m.CmdCommitGetResp
	}
	return nil
}

func init() {
	proto.RegisterType((*LockInfo)(nil), "kvrpcpb.LockInfo")
	proto.RegisterType((*ResultType)(nil), "kvrpcpb.ResultType")
	proto.RegisterType((*KeyAddress)(nil), "kvrpcpb.KeyAddress")
	proto.RegisterType((*CmdGetRequest)(nil), "kvrpcpb.CmdGetRequest")
	proto.RegisterType((*CmdGetResponse)(nil), "kvrpcpb.CmdGetResponse")
	proto.RegisterType((*CmdScanRequest)(nil), "kvrpcpb.CmdScanRequest")
	proto.RegisterType((*Item)(nil), "kvrpcpb.Item")
	proto.RegisterType((*CmdScanResponse)(nil), "kvrpcpb.CmdScanResponse")
	proto.RegisterType((*Mutation)(nil), "kvrpcpb.Mutation")
	proto.RegisterType((*CmdPrewriteRequest)(nil), "kvrpcpb.CmdPrewriteRequest")
	proto.RegisterType((*CmdPrewriteResponse)(nil), "kvrpcpb.CmdPrewriteResponse")
	proto.RegisterType((*CmdCommitRequest)(nil), "kvrpcpb.CmdCommitRequest")
	proto.RegisterType((*CmdCommitResponse)(nil), "kvrpcpb.CmdCommitResponse")
	proto.RegisterType((*CmdCleanupRequest)(nil), "kvrpcpb.CmdCleanupRequest")
	proto.RegisterType((*CmdCleanupResponse)(nil), "kvrpcpb.CmdCleanupResponse")
	proto.RegisterType((*CmdRollbackThenGetRequest)(nil), "kvrpcpb.CmdRollbackThenGetRequest")
	proto.RegisterType((*CmdRollbackThenGetResponse)(nil), "kvrpcpb.CmdRollbackThenGetResponse")
	proto.RegisterType((*CmdCommitThenGetRequest)(nil), "kvrpcpb.CmdCommitThenGetRequest")
	proto.RegisterType((*CmdCommitThenGetResponse)(nil), "kvrpcpb.CmdCommitThenGetResponse")
	proto.RegisterType((*Request)(nil), "kvrpcpb.Request")
	proto.RegisterType((*Response)(nil), "kvrpcpb.Response")
	proto.RegisterEnum("kvrpcpb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("kvrpcpb.Operator", Operator_name, Operator_value)
	proto.RegisterEnum("kvrpcpb.ResultType_Type", ResultType_Type_name, ResultType_Type_value)
}

var fileDescriptor0 = []byte{
	// 980 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0x5d, 0x73, 0xdb, 0x44,
	0x14, 0x1d, 0x5b, 0xf2, 0xd7, 0x95, 0x95, 0x28, 0x9b, 0x90, 0xb8, 0xe6, 0xab, 0x08, 0x98, 0x29,
	0xa5, 0x35, 0x43, 0x98, 0xd2, 0x99, 0x4e, 0x61, 0xa6, 0x53, 0x0a, 0x04, 0x0a, 0xc9, 0xa4, 0x79,
	0x82, 0x07, 0x8f, 0x6c, 0x5f, 0x52, 0x8f, 0x25, 0x4b, 0x5d, 0xad, 0xc3, 0xf8, 0x8d, 0x07, 0x86,
	0x77, 0xfe, 0x02, 0xff, 0x80, 0x7f, 0xc8, 0xdd, 0x0f, 0xc9, 0x92, 0xe5, 0x84, 0x26, 0xd3, 0x97,
	0xc4, 0x5a, 0xed, 0x39, 0x7b, 0xee, 0x3d, 0xf7, 0x5e, 0x2d, 0xb8, 0xb3, 0x0b, 0x9e, 0x8c, 0x93,
	0xd1, 0x20, 0xe1, 0xb1, 0x88, 0x59, 0xcb, 0x3c, 0xf6, 0xbb, 0x11, 0x8a, 0x20, 0x5b, 0xee, 0xbb,
	0xc8, 0x79, 0xcc, 0xb3, 0x47, 0xff, 0x4b, 0x68, 0x3f, 0x8f, 0xc7, 0xb3, 0xa3, 0xf9, 0x6f, 0x31,
	0xdb, 0x83, 0x6e, 0xc2, 0xa7, 0x51, 0xc0, 0x97, 0xc3, 0x90, 0xd6, 0x7a, 0xb5, 0xdb, 0xb5, 0x3b,
	0x5d, 0xb9, 0x2a, 0x9f, 0x86, 0x17, 0xc8, 0xd3, 0x69, 0x3c, 0xef, 0xd5, 0x69, 0xd5, 0xf6, 0xff,
	0xa8, 0x03, 0x9c, 0x62, 0xba, 0x08, 0xc5, 0xd9, 0x32, 0x41, 0x76, 0x1f, 0x6c, 0x41, 0xff, 0x15,
	0x64, 0xeb, 0xb0, 0x37, 0xc8, 0xa4, 0xac, 0xb6, 0x0c, 0xe4, 0x9f, 0x47, 0x8d, 0x63, 0xf1, 0x12,
	0x39, 0x73, 0xc0, 0x8a, 0xd2, 0x73, 0x45, 0xd5, 0x61, 0x1f, 0x41, 0x47, 0x1d, 0x30, 0x25, 0x0d,
	0x3d, 0x8b, 0x96, 0x9c, 0xc3, 0x9d, 0x9c, 0x20, 0x17, 0x77, 0x0f, 0x9c, 0x10, 0x83, 0x09, 0x72,
	0xbd, 0xcf, 0x56, 0xfb, 0x0e, 0x06, 0x59, 0x34, 0x3f, 0xc7, 0xe2, 0xb9, 0x7a, 0xfd, 0x4c, 0x2e,
	0xf8, 0x23, 0xb0, 0x95, 0xae, 0x26, 0xd4, 0x8f, 0x67, 0x5e, 0x8d, 0xb9, 0xd0, 0x39, 0x45, 0xc1,
	0x97, 0xc1, 0x28, 0x44, 0xaf, 0xce, 0x00, 0x9a, 0x92, 0x18, 0x27, 0x9e, 0x25, 0x5f, 0x3d, 0x8d,
	0xa3, 0x68, 0x2a, 0x04, 0x3d, 0xda, 0x6c, 0x8b, 0xe2, 0x8a, 0xc3, 0x10, 0x27, 0xa3, 0x60, 0x3c,
	0xf3, 0x1a, 0xf2, 0x75, 0xce, 0xed, 0x35, 0x59, 0x07, 0x74, 0x08, 0x5e, 0xc7, 0xff, 0x1e, 0xe0,
	0x47, 0x5c, 0x3e, 0x99, 0x4c, 0x38, 0xa6, 0xa9, 0x0c, 0x69, 0x86, 0x4b, 0x93, 0xb3, 0x1d, 0xe8,
	0x70, 0x3c, 0xa7, 0x6c, 0x0d, 0xa7, 0x13, 0x9d, 0x30, 0xd6, 0x07, 0x3b, 0x41, 0xe4, 0x26, 0xc0,
	0xee, 0xc0, 0x98, 0x72, 0x42, 0x6b, 0xfe, 0x0f, 0xe0, 0x3e, 0x8d, 0x26, 0xdf, 0xa1, 0x38, 0xc5,
	0x57, 0x0b, 0x4c, 0x05, 0xbb, 0x03, 0x0e, 0x91, 0x0d, 0x03, 0xcd, 0xad, 0x48, 0x9d, 0xc3, 0xdd,
	0x3c, 0x29, 0x85, 0x63, 0xb7, 0xa1, 0x55, 0x36, 0xe6, 0x5b, 0xd8, 0xca, 0xb8, 0xd2, 0x24, 0x9e,
	0xa7, 0xc8, 0x3e, 0x86, 0x36, 0x6d, 0x1d, 0xe6, 0xfe, 0x14, 0x99, 0x0a, 0x16, 0xba, 0xd0, 0xb8,
	0x08, 0xc2, 0x05, 0x2a, 0x9e, 0xae, 0xff, 0x8b, 0xe2, 0x79, 0x31, 0x0e, 0xe6, 0xd7, 0x17, 0x45,
	0x54, 0xe1, 0x94, 0x32, 0xaa, 0xa8, 0xdc, 0xa2, 0x46, 0x4b, 0x69, 0x3c, 0x02, 0xfb, 0x48, 0x60,
	0xf4, 0xba, 0xca, 0x4c, 0x6a, 0x95, 0xae, 0x95, 0x4c, 0x4b, 0xc9, 0xfc, 0x0a, 0xb6, 0x73, 0x99,
	0x26, 0x5e, 0x80, 0x7a, 0xac, 0x8b, 0xb7, 0xcd, 0xde, 0x83, 0x16, 0x57, 0x44, 0x29, 0xc1, 0x2d,
	0x3a, 0xc0, 0xcd, 0x0f, 0x90, 0x0a, 0xfc, 0x67, 0xd0, 0xfe, 0x69, 0x21, 0x02, 0x41, 0xda, 0xd8,
	0xbb, 0x84, 0x4b, 0x4c, 0x05, 0xaf, 0x0a, 0xf0, 0x38, 0x41, 0x1e, 0x88, 0x98, 0x5f, 0xa9, 0xe2,
	0xef, 0x1a, 0x30, 0x92, 0x71, 0xc2, 0xf1, 0x77, 0x3e, 0x15, 0x98, 0x65, 0x8c, 0x2a, 0x3b, 0x32,
	0xec, 0x32, 0x5f, 0x56, 0xa9, 0xb2, 0xf3, 0x73, 0xd7, 0xdb, 0x4e, 0x9f, 0xf0, 0x16, 0xb8, 0xa9,
	0x08, 0xb8, 0x18, 0x96, 0x52, 0xb7, 0x6e, 0x82, 0x7d, 0xa9, 0x09, 0xfe, 0x13, 0xd8, 0x2d, 0x49,
	0xba, 0x41, 0x76, 0x42, 0xf0, 0x88, 0x42, 0x77, 0x47, 0x16, 0x53, 0x45, 0x57, 0x4d, 0xe9, 0xfa,
	0x04, 0xba, 0xa4, 0x2b, 0xcd, 0x85, 0x69, 0xbe, 0x8d, 0xd5, 0xb1, 0x0f, 0x5b, 0x63, 0x45, 0x59,
	0x0e, 0xcd, 0x7f, 0x1f, 0x76, 0x0a, 0xa7, 0x55, 0xe5, 0xfa, 0x67, 0x7a, 0x03, 0x8d, 0x81, 0xf9,
	0x22, 0xb9, 0x7e, 0x55, 0x56, 0x94, 0xeb, 0x86, 0x79, 0xa1, 0xac, 0xcb, 0x59, 0xaf, 0xd7, 0x34,
	0xd5, 0x58, 0x34, 0xe9, 0xaf, 0x70, 0x8b, 0x48, 0xe5, 0x20, 0x91, 0x63, 0xe4, 0xec, 0x25, 0xce,
	0x6f, 0xd4, 0xdd, 0x9b, 0x67, 0xef, 0x43, 0xe8, 0x6f, 0x22, 0xdf, 0x60, 0xf0, 0x5a, 0x4f, 0xff,
	0x55, 0x83, 0x83, 0x3c, 0xc5, 0x6f, 0x56, 0xd4, 0x65, 0xae, 0xb2, 0x5d, 0x70, 0xce, 0x71, 0xb5,
	0x68, 0xab, 0x08, 0x1e, 0x40, 0xaf, 0xaa, 0xe3, 0xff, 0xf5, 0xff, 0x63, 0x41, 0x2b, 0xd3, 0xeb,
	0x97, 0xbe, 0x38, 0x7b, 0xab, 0xb6, 0x22, 0x89, 0xc1, 0x39, 0x2a, 0x77, 0x3e, 0x05, 0x67, 0x1c,
	0x4d, 0x86, 0xf2, 0x7c, 0x8e, 0xaf, 0x14, 0x89, 0x73, 0xb8, 0x9f, 0x6f, 0x2d, 0xcf, 0xdc, 0xfb,
	0xd0, 0x95, 0x9b, 0x53, 0x1a, 0x25, 0x6a, 0xb7, 0x65, 0xbe, 0x30, 0x85, 0xdd, 0xc5, 0x69, 0xf8,
	0x00, 0x3c, 0xb9, 0x3d, 0x31, 0xfd, 0xa5, 0x20, 0xba, 0x1b, 0xdf, 0x2e, 0x42, 0xd6, 0x47, 0xc2,
	0xe7, 0x94, 0x26, 0x82, 0x99, 0x54, 0x49, 0x50, 0x43, 0x81, 0x6e, 0x15, 0x41, 0xe5, 0x8e, 0xfb,
	0x02, 0xb6, 0x15, 0x44, 0x57, 0xa8, 0xc2, 0x34, 0x15, 0xa6, 0x5f, 0xc2, 0x94, 0xdb, 0xe2, 0x91,
	0x3e, 0x87, 0x8f, 0xf2, 0xe8, 0x5b, 0x0a, 0xe3, 0x17, 0x31, 0x97, 0xd4, 0xe7, 0x63, 0x60, 0x05,
	0x8d, 0x19, 0xbe, 0xad, 0xf0, 0xb7, 0xab, 0x3a, 0xcb, 0x68, 0xff, 0x5f, 0x0b, 0xda, 0xb9, 0x99,
	0xaf, 0xe3, 0x92, 0x49, 0xbc, 0x3e, 0x27, 0x4d, 0x8c, 0x4d, 0x07, 0x15, 0x9b, 0x0c, 0xe5, 0x67,
	0xe0, 0x16, 0x7c, 0xa2, 0xfd, 0xda, 0xa8, 0x5e, 0xd5, 0x28, 0x03, 0x78, 0x08, 0x3b, 0x6b, 0x4e,
	0x11, 0x48, 0x5b, 0xf5, 0xce, 0x66, 0xab, 0x0c, 0x30, 0x4b, 0x7c, 0xe6, 0x15, 0xc1, 0x1a, 0x1b,
	0x12, 0x5f, 0x1e, 0x58, 0xa6, 0x2e, 0x56, 0x6e, 0x11, 0xaa, 0x59, 0xad, 0x8b, 0xf5, 0x79, 0xf3,
	0x58, 0x9f, 0x95, 0xfb, 0x45, 0x28, 0x6d, 0xd8, 0x87, 0x57, 0x1a, 0x66, 0xd0, 0x5f, 0xc3, 0x6e,
	0xc5, 0x31, 0x62, 0xd0, 0x96, 0x7d, 0x70, 0x85, 0x65, 0x1a, 0x7f, 0xf7, 0xcf, 0x1a, 0x38, 0x45,
	0x4b, 0xe8, 0x7e, 0xa4, 0xb3, 0x4e, 0x57, 0x27, 0x07, 0x5a, 0x26, 0xa3, 0x74, 0x71, 0xda, 0x06,
	0xa7, 0x90, 0x29, 0x73, 0x7b, 0xca, 0x58, 0xf5, 0xed, 0x69, 0x15, 0x1c, 0xdd, 0x9e, 0xf6, 0xd5,
	0x70, 0x5d, 0x93, 0x4d, 0xd7, 0xa8, 0xbd, 0xc2, 0x97, 0x25, 0x5b, 0x6d, 0xdd, 0xbd, 0x07, 0xed,
	0xfc, 0x73, 0x2b, 0x2f, 0x5a, 0xc9, 0xc9, 0x42, 0x2a, 0x50, 0x3f, 0xbf, 0xc1, 0x50, 0x5f, 0xdc,
	0x8e, 0x13, 0x79, 0x75, 0xf3, 0xac, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xe5, 0x8e, 0x56,
	0xf0, 0x0a, 0x00, 0x00,
}
