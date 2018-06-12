// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package rpcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type NonParamsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonParamsRequest) Reset()         { *m = NonParamsRequest{} }
func (m *NonParamsRequest) String() string { return proto.CompactTextString(m) }
func (*NonParamsRequest) ProtoMessage()    {}
func (*NonParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{0}
}
func (m *NonParamsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonParamsRequest.Unmarshal(m, b)
}
func (m *NonParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonParamsRequest.Marshal(b, m, deterministic)
}
func (dst *NonParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonParamsRequest.Merge(dst, src)
}
func (m *NonParamsRequest) XXX_Size() int {
	return xxx_messageInfo_NonParamsRequest.Size(m)
}
func (m *NonParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NonParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NonParamsRequest proto.InternalMessageInfo

type GetStateResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStateResponse) Reset()         { *m = GetStateResponse{} }
func (m *GetStateResponse) String() string { return proto.CompactTextString(m) }
func (*GetStateResponse) ProtoMessage()    {}
func (*GetStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{1}
}
func (m *GetStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateResponse.Unmarshal(m, b)
}
func (m *GetStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateResponse.Marshal(b, m, deterministic)
}
func (dst *GetStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateResponse.Merge(dst, src)
}
func (m *GetStateResponse) XXX_Size() int {
	return xxx_messageInfo_GetStateResponse.Size(m)
}
func (m *GetStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateResponse proto.InternalMessageInfo

func (m *GetStateResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

// assets
type ListAssetsRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAssetsRequest) Reset()         { *m = ListAssetsRequest{} }
func (m *ListAssetsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAssetsRequest) ProtoMessage()    {}
func (*ListAssetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{2}
}
func (m *ListAssetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAssetsRequest.Unmarshal(m, b)
}
func (m *ListAssetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAssetsRequest.Marshal(b, m, deterministic)
}
func (dst *ListAssetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAssetsRequest.Merge(dst, src)
}
func (m *ListAssetsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAssetsRequest.Size(m)
}
func (m *ListAssetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAssetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAssetsRequest proto.InternalMessageInfo

func (m *ListAssetsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ListAssetsResponse struct {
	Assets               []*Asset `protobuf:"bytes,1,rep,name=assets" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAssetsResponse) Reset()         { *m = ListAssetsResponse{} }
func (m *ListAssetsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAssetsResponse) ProtoMessage()    {}
func (*ListAssetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{3}
}
func (m *ListAssetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAssetsResponse.Unmarshal(m, b)
}
func (m *ListAssetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAssetsResponse.Marshal(b, m, deterministic)
}
func (dst *ListAssetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAssetsResponse.Merge(dst, src)
}
func (m *ListAssetsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAssetsResponse.Size(m)
}
func (m *ListAssetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAssetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAssetsResponse proto.InternalMessageInfo

func (m *ListAssetsResponse) GetAssets() []*Asset {
	if m != nil {
		return m.Assets
	}
	return nil
}

type Asset struct {
	AssetID              string   `protobuf:"bytes,1,opt,name=assetID" json:"assetID,omitempty"`
	Amount               uint64   `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{4}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (dst *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(dst, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetAssetID() string {
	if m != nil {
		return m.AssetID
	}
	return ""
}

func (m *Asset) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// transactions
type Input struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	AssetID              string   `protobuf:"bytes,2,opt,name=assetID" json:"assetID,omitempty"`
	Amount               uint64   `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address" json:"address,omitempty"`
	SpentOutputID        string   `protobuf:"bytes,5,opt,name=spentOutputID" json:"spentOutputID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{5}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (dst *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(dst, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Input) GetAssetID() string {
	if m != nil {
		return m.AssetID
	}
	return ""
}

func (m *Input) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Input) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Input) GetSpentOutputID() string {
	if m != nil {
		return m.SpentOutputID
	}
	return ""
}

type Output struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	AssetID              string   `protobuf:"bytes,2,opt,name=assetID" json:"assetID,omitempty"`
	Amount               uint64   `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address" json:"address,omitempty"`
	OutputID             string   `protobuf:"bytes,5,opt,name=OutputID" json:"OutputID,omitempty"`
	Position             int32    `protobuf:"varint,6,opt,name=position" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{6}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (dst *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(dst, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Output) GetAssetID() string {
	if m != nil {
		return m.AssetID
	}
	return ""
}

func (m *Output) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Output) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Output) GetOutputID() string {
	if m != nil {
		return m.OutputID
	}
	return ""
}

func (m *Output) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

type TX struct {
	ID                     string    `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Timestamp              uint64    `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	BlockID                string    `protobuf:"bytes,3,opt,name=blockID" json:"blockID,omitempty"`
	BlockHeight            uint64    `protobuf:"varint,4,opt,name=blockHeight" json:"blockHeight,omitempty"`
	Position               uint32    `protobuf:"varint,5,opt,name=position" json:"position,omitempty"`
	BlockTransactionsCount uint32    `protobuf:"varint,6,opt,name=blockTransactionsCount" json:"blockTransactionsCount,omitempty"`
	Confirmation           uint64    `protobuf:"varint,7,opt,name=confirmation" json:"confirmation,omitempty"`
	StatusFail             bool      `protobuf:"varint,8,opt,name=statusFail" json:"statusFail,omitempty"`
	Inputs                 []*Input  `protobuf:"bytes,9,rep,name=inputs" json:"inputs,omitempty"`
	Outputs                []*Output `protobuf:"bytes,10,rep,name=outputs" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}  `json:"-"`
	XXX_unrecognized       []byte    `json:"-"`
	XXX_sizecache          int32     `json:"-"`
}

func (m *TX) Reset()         { *m = TX{} }
func (m *TX) String() string { return proto.CompactTextString(m) }
func (*TX) ProtoMessage()    {}
func (*TX) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{7}
}
func (m *TX) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TX.Unmarshal(m, b)
}
func (m *TX) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TX.Marshal(b, m, deterministic)
}
func (dst *TX) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TX.Merge(dst, src)
}
func (m *TX) XXX_Size() int {
	return xxx_messageInfo_TX.Size(m)
}
func (m *TX) XXX_DiscardUnknown() {
	xxx_messageInfo_TX.DiscardUnknown(m)
}

var xxx_messageInfo_TX proto.InternalMessageInfo

func (m *TX) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *TX) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TX) GetBlockID() string {
	if m != nil {
		return m.BlockID
	}
	return ""
}

func (m *TX) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *TX) GetPosition() uint32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *TX) GetBlockTransactionsCount() uint32 {
	if m != nil {
		return m.BlockTransactionsCount
	}
	return 0
}

func (m *TX) GetConfirmation() uint64 {
	if m != nil {
		return m.Confirmation
	}
	return 0
}

func (m *TX) GetStatusFail() bool {
	if m != nil {
		return m.StatusFail
	}
	return false
}

func (m *TX) GetInputs() []*Input {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TX) GetOutputs() []*Output {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type ListTransactionsRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	AssetID              string   `protobuf:"bytes,2,opt,name=assetID" json:"assetID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTransactionsRequest) Reset()         { *m = ListTransactionsRequest{} }
func (m *ListTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTransactionsRequest) ProtoMessage()    {}
func (*ListTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{8}
}
func (m *ListTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTransactionsRequest.Unmarshal(m, b)
}
func (m *ListTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTransactionsRequest.Marshal(b, m, deterministic)
}
func (dst *ListTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTransactionsRequest.Merge(dst, src)
}
func (m *ListTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_ListTransactionsRequest.Size(m)
}
func (m *ListTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTransactionsRequest proto.InternalMessageInfo

func (m *ListTransactionsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ListTransactionsRequest) GetAssetID() string {
	if m != nil {
		return m.AssetID
	}
	return ""
}

type ListTransactionsResponse struct {
	Transactions         []*TX    `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTransactionsResponse) Reset()         { *m = ListTransactionsResponse{} }
func (m *ListTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTransactionsResponse) ProtoMessage()    {}
func (*ListTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{9}
}
func (m *ListTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTransactionsResponse.Unmarshal(m, b)
}
func (m *ListTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTransactionsResponse.Marshal(b, m, deterministic)
}
func (dst *ListTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTransactionsResponse.Merge(dst, src)
}
func (m *ListTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTransactionsResponse.Size(m)
}
func (m *ListTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTransactionsResponse proto.InternalMessageInfo

func (m *ListTransactionsResponse) GetTransactions() []*TX {
	if m != nil {
		return m.Transactions
	}
	return nil
}

// submit tx
type SubmitTransactionRequest struct {
	RawTransaction       string   `protobuf:"bytes,1,opt,name=rawTransaction" json:"rawTransaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitTransactionRequest) Reset()         { *m = SubmitTransactionRequest{} }
func (m *SubmitTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionRequest) ProtoMessage()    {}
func (*SubmitTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{10}
}
func (m *SubmitTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionRequest.Unmarshal(m, b)
}
func (m *SubmitTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionRequest.Marshal(b, m, deterministic)
}
func (dst *SubmitTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionRequest.Merge(dst, src)
}
func (m *SubmitTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionRequest.Size(m)
}
func (m *SubmitTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionRequest proto.InternalMessageInfo

func (m *SubmitTransactionRequest) GetRawTransaction() string {
	if m != nil {
		return m.RawTransaction
	}
	return ""
}

type SubmitTransactionResponse struct {
	TxID                 string   `protobuf:"bytes,1,opt,name=txID" json:"txID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitTransactionResponse) Reset()         { *m = SubmitTransactionResponse{} }
func (m *SubmitTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionResponse) ProtoMessage()    {}
func (*SubmitTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{11}
}
func (m *SubmitTransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionResponse.Unmarshal(m, b)
}
func (m *SubmitTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionResponse.Marshal(b, m, deterministic)
}
func (dst *SubmitTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionResponse.Merge(dst, src)
}
func (m *SubmitTransactionResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionResponse.Size(m)
}
func (m *SubmitTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionResponse proto.InternalMessageInfo

func (m *SubmitTransactionResponse) GetTxID() string {
	if m != nil {
		return m.TxID
	}
	return ""
}

// gas
type EstimateTransactionGasRequest struct {
	TransactionTemplate  []byte   `protobuf:"bytes,1,opt,name=transactionTemplate,proto3" json:"transactionTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstimateTransactionGasRequest) Reset()         { *m = EstimateTransactionGasRequest{} }
func (m *EstimateTransactionGasRequest) String() string { return proto.CompactTextString(m) }
func (*EstimateTransactionGasRequest) ProtoMessage()    {}
func (*EstimateTransactionGasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{12}
}
func (m *EstimateTransactionGasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstimateTransactionGasRequest.Unmarshal(m, b)
}
func (m *EstimateTransactionGasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstimateTransactionGasRequest.Marshal(b, m, deterministic)
}
func (dst *EstimateTransactionGasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateTransactionGasRequest.Merge(dst, src)
}
func (m *EstimateTransactionGasRequest) XXX_Size() int {
	return xxx_messageInfo_EstimateTransactionGasRequest.Size(m)
}
func (m *EstimateTransactionGasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateTransactionGasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateTransactionGasRequest proto.InternalMessageInfo

func (m *EstimateTransactionGasRequest) GetTransactionTemplate() []byte {
	if m != nil {
		return m.TransactionTemplate
	}
	return nil
}

type EstimateTransactionGasResponse struct {
	TotalNeu             int64    `protobuf:"varint,1,opt,name=totalNeu" json:"totalNeu,omitempty"`
	StorageNeu           int64    `protobuf:"varint,2,opt,name=storageNeu" json:"storageNeu,omitempty"`
	VmNeu                int64    `protobuf:"varint,3,opt,name=vmNeu" json:"vmNeu,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstimateTransactionGasResponse) Reset()         { *m = EstimateTransactionGasResponse{} }
func (m *EstimateTransactionGasResponse) String() string { return proto.CompactTextString(m) }
func (*EstimateTransactionGasResponse) ProtoMessage()    {}
func (*EstimateTransactionGasResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_a31bdcf93cfd96d0, []int{13}
}
func (m *EstimateTransactionGasResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstimateTransactionGasResponse.Unmarshal(m, b)
}
func (m *EstimateTransactionGasResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstimateTransactionGasResponse.Marshal(b, m, deterministic)
}
func (dst *EstimateTransactionGasResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateTransactionGasResponse.Merge(dst, src)
}
func (m *EstimateTransactionGasResponse) XXX_Size() int {
	return xxx_messageInfo_EstimateTransactionGasResponse.Size(m)
}
func (m *EstimateTransactionGasResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateTransactionGasResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateTransactionGasResponse proto.InternalMessageInfo

func (m *EstimateTransactionGasResponse) GetTotalNeu() int64 {
	if m != nil {
		return m.TotalNeu
	}
	return 0
}

func (m *EstimateTransactionGasResponse) GetStorageNeu() int64 {
	if m != nil {
		return m.StorageNeu
	}
	return 0
}

func (m *EstimateTransactionGasResponse) GetVmNeu() int64 {
	if m != nil {
		return m.VmNeu
	}
	return 0
}

func init() {
	proto.RegisterType((*NonParamsRequest)(nil), "rpcpb.NonParamsRequest")
	proto.RegisterType((*GetStateResponse)(nil), "rpcpb.GetStateResponse")
	proto.RegisterType((*ListAssetsRequest)(nil), "rpcpb.ListAssetsRequest")
	proto.RegisterType((*ListAssetsResponse)(nil), "rpcpb.ListAssetsResponse")
	proto.RegisterType((*Asset)(nil), "rpcpb.Asset")
	proto.RegisterType((*Input)(nil), "rpcpb.Input")
	proto.RegisterType((*Output)(nil), "rpcpb.Output")
	proto.RegisterType((*TX)(nil), "rpcpb.TX")
	proto.RegisterType((*ListTransactionsRequest)(nil), "rpcpb.ListTransactionsRequest")
	proto.RegisterType((*ListTransactionsResponse)(nil), "rpcpb.ListTransactionsResponse")
	proto.RegisterType((*SubmitTransactionRequest)(nil), "rpcpb.SubmitTransactionRequest")
	proto.RegisterType((*SubmitTransactionResponse)(nil), "rpcpb.SubmitTransactionResponse")
	proto.RegisterType((*EstimateTransactionGasRequest)(nil), "rpcpb.EstimateTransactionGasRequest")
	proto.RegisterType((*EstimateTransactionGasResponse)(nil), "rpcpb.EstimateTransactionGasResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiServiceClient interface {
	// state
	GetState(ctx context.Context, in *NonParamsRequest, opts ...grpc.CallOption) (*GetStateResponse, error)
	ListAssets(ctx context.Context, in *ListAssetsRequest, opts ...grpc.CallOption) (*ListAssetsResponse, error)
	ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error)
	SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error)
	EstimateTransactionGas(ctx context.Context, in *EstimateTransactionGasRequest, opts ...grpc.CallOption) (*EstimateTransactionGasResponse, error)
}

type apiServiceClient struct {
	cc *grpc.ClientConn
}

func NewApiServiceClient(cc *grpc.ClientConn) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) GetState(ctx context.Context, in *NonParamsRequest, opts ...grpc.CallOption) (*GetStateResponse, error) {
	out := new(GetStateResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.ApiService/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListAssets(ctx context.Context, in *ListAssetsRequest, opts ...grpc.CallOption) (*ListAssetsResponse, error) {
	out := new(ListAssetsResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.ApiService/ListAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error) {
	out := new(ListTransactionsResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.ApiService/ListTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error) {
	out := new(SubmitTransactionResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.ApiService/SubmitTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) EstimateTransactionGas(ctx context.Context, in *EstimateTransactionGasRequest, opts ...grpc.CallOption) (*EstimateTransactionGasResponse, error) {
	out := new(EstimateTransactionGasResponse)
	err := c.cc.Invoke(ctx, "/rpcpb.ApiService/EstimateTransactionGas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
type ApiServiceServer interface {
	// state
	GetState(context.Context, *NonParamsRequest) (*GetStateResponse, error)
	ListAssets(context.Context, *ListAssetsRequest) (*ListAssetsResponse, error)
	ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error)
	SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error)
	EstimateTransactionGas(context.Context, *EstimateTransactionGasRequest) (*EstimateTransactionGasResponse, error)
}

func RegisterApiServiceServer(s *grpc.Server, srv ApiServiceServer) {
	s.RegisterService(&_ApiService_serviceDesc, srv)
}

func _ApiService_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ApiService/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetState(ctx, req.(*NonParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ApiService/ListAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListAssets(ctx, req.(*ListAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ApiService/ListTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListTransactions(ctx, req.(*ListTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_SubmitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).SubmitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ApiService/SubmitTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).SubmitTransaction(ctx, req.(*SubmitTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_EstimateTransactionGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstimateTransactionGasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).EstimateTransactionGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ApiService/EstimateTransactionGas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).EstimateTransactionGas(ctx, req.(*EstimateTransactionGasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _ApiService_GetState_Handler,
		},
		{
			MethodName: "ListAssets",
			Handler:    _ApiService_ListAssets_Handler,
		},
		{
			MethodName: "ListTransactions",
			Handler:    _ApiService_ListTransactions_Handler,
		},
		{
			MethodName: "SubmitTransaction",
			Handler:    _ApiService_SubmitTransaction_Handler,
		},
		{
			MethodName: "EstimateTransactionGas",
			Handler:    _ApiService_EstimateTransactionGas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_rpc_a31bdcf93cfd96d0) }

var fileDescriptor_rpc_a31bdcf93cfd96d0 = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcb, 0x52, 0x1b, 0x47,
	0x14, 0xad, 0x19, 0x3d, 0x90, 0x2e, 0x82, 0x40, 0x87, 0x88, 0x61, 0x42, 0xc8, 0xd0, 0x45, 0x82,
	0x8a, 0x2a, 0xa1, 0x84, 0x54, 0xa5, 0x2a, 0xec, 0x48, 0x48, 0x88, 0xaa, 0x12, 0x48, 0x06, 0x2d,
	0xd8, 0x78, 0xd1, 0x12, 0x6d, 0x79, 0xca, 0x9a, 0x87, 0xa7, 0x7b, 0xb0, 0xd9, 0x7a, 0x6f, 0x6f,
	0xfc, 0x03, 0xfe, 0x13, 0x7f, 0x84, 0xcb, 0x7f, 0xe0, 0x0f, 0x71, 0xf5, 0x9d, 0x9e, 0xa1, 0x25,
	0x24, 0xb3, 0xf2, 0xae, 0xef, 0x3d, 0x67, 0xce, 0x3d, 0xba, 0x75, 0xba, 0x05, 0xcd, 0x34, 0x19,
	0x1d, 0x26, 0x69, 0x2c, 0x63, 0x52, 0x4b, 0x93, 0x51, 0x32, 0x74, 0xb7, 0xc7, 0x71, 0x3c, 0x9e,
	0xf0, 0x1e, 0x4b, 0x82, 0x1e, 0x8b, 0xa2, 0x58, 0x32, 0x19, 0xc4, 0x91, 0xc8, 0x49, 0x94, 0xc0,
	0xda, 0x79, 0x1c, 0xfd, 0xc7, 0x52, 0x16, 0x0a, 0x9f, 0x3f, 0xcb, 0xb8, 0x90, 0xf4, 0x00, 0xd6,
	0xce, 0xb8, 0xbc, 0x94, 0x4c, 0x72, 0x9f, 0x8b, 0x24, 0x8e, 0x04, 0x27, 0x6d, 0xa8, 0x0b, 0xc9,
	0x64, 0x26, 0x1c, 0xcb, 0xb3, 0x3a, 0x4d, 0x5f, 0x57, 0xb4, 0x0b, 0xeb, 0xff, 0x04, 0x42, 0x9e,
	0x08, 0xc1, 0x65, 0x21, 0x40, 0x1c, 0x58, 0x62, 0xd7, 0xd7, 0x29, 0x17, 0x05, 0xbb, 0x28, 0xe9,
	0x31, 0x10, 0x93, 0xae, 0xc5, 0xf7, 0xa0, 0xce, 0xb0, 0xe3, 0x58, 0x5e, 0xa5, 0xb3, 0x7c, 0xd4,
	0x3a, 0x44, 0xeb, 0x87, 0x48, 0xf3, 0x35, 0x46, 0x7f, 0x83, 0x1a, 0x36, 0x50, 0x5e, 0x1d, 0xfa,
	0xa7, 0xa5, 0x7c, 0x5e, 0x2a, 0x97, 0x2c, 0x8c, 0xb3, 0x48, 0x3a, 0xb6, 0x67, 0x75, 0xaa, 0xbe,
	0xae, 0xe8, 0x2b, 0x0b, 0x6a, 0xfd, 0x28, 0xc9, 0x24, 0x21, 0x50, 0x95, 0xb7, 0x09, 0xd7, 0x1f,
	0xe2, 0xd9, 0xd4, 0xb3, 0x17, 0xe9, 0x55, 0x4c, 0x3d, 0xf3, 0x07, 0x56, 0xa7, 0x7e, 0x20, 0xd9,
	0x83, 0x15, 0x91, 0xf0, 0x48, 0x5e, 0x64, 0x32, 0xc9, 0x94, 0x62, 0x0d, 0xf1, 0xe9, 0x26, 0x7d,
	0x6b, 0x41, 0x3d, 0x2f, 0xbe, 0xb8, 0x21, 0x17, 0x1a, 0x33, 0x5e, 0xca, 0x5a, 0x61, 0x49, 0x2c,
	0x02, 0x95, 0x07, 0xa7, 0xee, 0x59, 0x9d, 0x9a, 0x5f, 0xd6, 0xf4, 0x83, 0x0d, 0xf6, 0xe0, 0x8a,
	0xac, 0x82, 0x5d, 0xae, 0xd9, 0xee, 0x9f, 0x92, 0x6d, 0x68, 0xca, 0x20, 0xe4, 0x42, 0xb2, 0x30,
	0xd1, 0x4b, 0xbe, 0x6b, 0x28, 0x1b, 0xc3, 0x49, 0x3c, 0x7a, 0xda, 0x3f, 0x45, 0x7f, 0x4d, 0xbf,
	0x28, 0x89, 0x07, 0xcb, 0x78, 0xfc, 0x9b, 0x07, 0xe3, 0x27, 0x12, 0x4d, 0x56, 0x7d, 0xb3, 0x35,
	0x65, 0x46, 0x19, 0x5d, 0xb9, 0x33, 0x43, 0x7e, 0x85, 0x36, 0x52, 0x07, 0x29, 0x8b, 0x04, 0x1b,
	0x61, 0x80, 0xff, 0xc0, 0x35, 0xd4, 0x91, 0xb9, 0x00, 0x25, 0x14, 0x5a, 0xa3, 0x38, 0x7a, 0x1c,
	0xa4, 0x21, 0x86, 0xde, 0x59, 0xc2, 0xb1, 0x53, 0x3d, 0xb2, 0x03, 0x90, 0x67, 0xf9, 0x2f, 0x16,
	0x4c, 0x9c, 0x86, 0x67, 0x75, 0x1a, 0xbe, 0xd1, 0x51, 0xe1, 0x0c, 0x54, 0x74, 0x84, 0xd3, 0x9c,
	0x0a, 0x27, 0xe6, 0xc9, 0xd7, 0x18, 0xd9, 0x87, 0xa5, 0x18, 0xd7, 0x2a, 0x1c, 0x40, 0xda, 0x8a,
	0xa6, 0xe5, 0xcb, 0xf6, 0x0b, 0x94, 0xfe, 0x0b, 0x9b, 0xea, 0x06, 0x98, 0x5e, 0x1f, 0xbc, 0x36,
	0x8b, 0x03, 0x41, 0xfb, 0xe0, 0xdc, 0x97, 0xd3, 0xd7, 0xaa, 0x0b, 0x2d, 0x69, 0xf4, 0xf5, 0xe5,
	0x6a, 0x6a, 0x63, 0x83, 0x2b, 0x7f, 0x0a, 0xa6, 0xbf, 0x83, 0x73, 0x99, 0x0d, 0xc3, 0xc0, 0x14,
	0x2b, 0xac, 0xfd, 0x08, 0xab, 0x29, 0x7b, 0x6e, 0x00, 0xda, 0xe1, 0x4c, 0x97, 0xf6, 0x60, 0x6b,
	0x8e, 0x86, 0xf6, 0xa3, 0xa2, 0xfe, 0xa2, 0x4c, 0x13, 0x9e, 0xe9, 0xff, 0xf0, 0xdd, 0x9f, 0x42,
	0x06, 0x21, 0x93, 0xdc, 0xf8, 0xe4, 0x8c, 0x95, 0x4b, 0xf9, 0x09, 0xbe, 0x36, 0x5c, 0x0e, 0x78,
	0x98, 0x4c, 0x98, 0xcc, 0xaf, 0x4b, 0xcb, 0x9f, 0x07, 0xd1, 0x14, 0x76, 0x16, 0x49, 0x6a, 0x23,
	0x2e, 0x34, 0x64, 0x2c, 0xd9, 0xe4, 0x9c, 0x67, 0x28, 0x54, 0xf1, 0xcb, 0x3a, 0x8f, 0x43, 0x9c,
	0xb2, 0x31, 0x57, 0xa8, 0x8d, 0xa8, 0xd1, 0x21, 0x1b, 0x50, 0xbb, 0x09, 0x15, 0x54, 0x41, 0x28,
	0x2f, 0x8e, 0xde, 0x55, 0x01, 0x4e, 0x92, 0xe0, 0x92, 0xa7, 0x37, 0xc1, 0x88, 0x93, 0x0b, 0x68,
	0x14, 0x2f, 0x28, 0xd9, 0xd4, 0xfb, 0x9e, 0x7d, 0x66, 0xdd, 0x02, 0x98, 0x7d, 0x6b, 0xe9, 0xfa,
	0xcb, 0xf7, 0x1f, 0xdf, 0xd8, 0xcb, 0xa4, 0xd9, 0xbb, 0xf9, 0xb9, 0x27, 0x50, 0xe4, 0x11, 0xc0,
	0xdd, 0xbb, 0x49, 0x1c, 0xfd, 0xe5, 0xbd, 0x97, 0xd7, 0xdd, 0x9a, 0x83, 0x68, 0x55, 0x17, 0x55,
	0x37, 0xe8, 0x57, 0x4a, 0x75, 0x12, 0x08, 0xd9, 0xcd, 0xdf, 0xd5, 0x63, 0xeb, 0x80, 0x64, 0xb0,
	0x36, 0x9b, 0x22, 0xb2, 0x63, 0x48, 0xcd, 0x49, 0xab, 0xfb, 0xfd, 0x42, 0x5c, 0x0f, 0xf4, 0x70,
	0xa0, 0x4b, 0xbf, 0x29, 0x07, 0x9a, 0x71, 0x53, 0x63, 0x6f, 0x61, 0xfd, 0x5e, 0x5a, 0x48, 0xa1,
	0xbb, 0x28, 0x8b, 0xae, 0xb7, 0x98, 0xa0, 0x27, 0xef, 0xe2, 0xe4, 0x6f, 0x69, 0x1b, 0x17, 0x88,
	0x34, 0x73, 0xb6, 0x1a, 0xfd, 0xda, 0x82, 0xf6, 0xfc, 0x94, 0x90, 0x3d, 0xad, 0xff, 0xd9, 0x5c,
	0xba, 0x3f, 0x3c, 0xc0, 0xd2, 0x56, 0xf6, 0xd1, 0xca, 0x2e, 0xdd, 0x56, 0x56, 0xb8, 0xe6, 0x9a,
	0x66, 0xba, 0x63, 0xa6, 0x76, 0x31, 0xac, 0xe3, 0xff, 0xf1, 0x2f, 0x9f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x4a, 0xdd, 0x46, 0x74, 0xc1, 0x07, 0x00, 0x00,
}
