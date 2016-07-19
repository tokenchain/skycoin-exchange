// Code generated by protoc-gen-go.
// source: pp.withdrawal.proto
// DO NOT EDIT!

package pp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WithdrawalReq struct {
	AccountId        []byte  `protobuf:"bytes,1,opt,name=account_id" json:"account_id,omitempty"`
	CoinType         *string `protobuf:"bytes,2,opt,name=coin_type" json:"coin_type,omitempty"`
	Coins            *uint64 `protobuf:"varint,3,opt,name=coins" json:"coins,omitempty"`
	OutputAddress    *string `protobuf:"bytes,4,opt,name=output_address" json:"output_address,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WithdrawalReq) Reset()                    { *m = WithdrawalReq{} }
func (m *WithdrawalReq) String() string            { return proto.CompactTextString(m) }
func (*WithdrawalReq) ProtoMessage()               {}
func (*WithdrawalReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *WithdrawalReq) GetAccountId() []byte {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *WithdrawalReq) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

func (m *WithdrawalReq) GetCoins() uint64 {
	if m != nil && m.Coins != nil {
		return *m.Coins
	}
	return 0
}

func (m *WithdrawalReq) GetOutputAddress() string {
	if m != nil && m.OutputAddress != nil {
		return *m.OutputAddress
	}
	return ""
}

type WithdrawalRes struct {
	AccountId        []byte  `protobuf:"bytes,10,opt,name=account_id" json:"account_id,omitempty"`
	NewTxid          *string `protobuf:"bytes,20,opt,name=new_txid" json:"new_txid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WithdrawalRes) Reset()                    { *m = WithdrawalRes{} }
func (m *WithdrawalRes) String() string            { return proto.CompactTextString(m) }
func (*WithdrawalRes) ProtoMessage()               {}
func (*WithdrawalRes) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *WithdrawalRes) GetAccountId() []byte {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *WithdrawalRes) GetNewTxid() string {
	if m != nil && m.NewTxid != nil {
		return *m.NewTxid
	}
	return ""
}

func init() {
	proto.RegisterType((*WithdrawalReq)(nil), "pp.WithdrawalReq")
	proto.RegisterType((*WithdrawalRes)(nil), "pp.WithdrawalRes")
}

func init() { proto.RegisterFile("pp.withdrawal.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xd0, 0x2b,
	0xcf, 0x2c, 0xc9, 0x48, 0x29, 0x4a, 0x2c, 0x4f, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2a, 0x28, 0x50, 0x8a, 0xe5, 0xe2, 0x0d, 0x87, 0x8b, 0x07, 0xa5, 0x16, 0x0a, 0x09, 0x71,
	0x71, 0x25, 0x26, 0x27, 0xe7, 0x97, 0xe6, 0x95, 0xc4, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xf0, 0x08, 0x09, 0x72, 0x71, 0x26, 0xe7, 0x67, 0xe6, 0xc5, 0x97, 0x54, 0x16, 0xa4, 0x4a, 0x30,
	0x01, 0x85, 0x38, 0x85, 0x78, 0xb9, 0x58, 0x41, 0x42, 0xc5, 0x12, 0xcc, 0x40, 0x2e, 0x8b, 0x90,
	0x18, 0x17, 0x5f, 0x7e, 0x69, 0x49, 0x41, 0x69, 0x49, 0x7c, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71,
	0xb1, 0x04, 0x0b, 0x48, 0x99, 0x92, 0x29, 0xaa, 0xf1, 0xc5, 0x68, 0xc6, 0x73, 0x81, 0x8d, 0x17,
	0xe0, 0xe2, 0xc8, 0x4b, 0x2d, 0x8f, 0x2f, 0xa9, 0x00, 0x8a, 0x88, 0x80, 0xb4, 0x01, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x2e, 0xf3, 0x08, 0x73, 0xaf, 0x00, 0x00, 0x00,
}
