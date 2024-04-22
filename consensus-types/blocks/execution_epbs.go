package blocks

import (
	fastssz "github.com/prysmaticlabs/fastssz"
	consensus_types "github.com/prysmaticlabs/prysm/v5/consensus-types"
	"github.com/prysmaticlabs/prysm/v5/consensus-types/interfaces"
	"github.com/prysmaticlabs/prysm/v5/math"
	enginev1 "github.com/prysmaticlabs/prysm/v5/proto/engine/v1"
	"google.golang.org/protobuf/proto"
)

// executionPayloadEpbs is a convenience wrapper around a ExecutionPayloadEPBS data structure.
type executionPayloadEpbs struct {
	p *enginev1.ExecutionPayloadEPBS
}

// WrappedExecutionPayloadEpbs is a constructor which wraps a protobuf signed execution payload envelope into an interface.
func WrappedExecutionPayloadEpbs(p *enginev1.ExecutionPayloadEPBS) (interfaces.ExecutionData, error) {
	w := executionPayloadEpbs{p: p}
	if w.IsNil() {
		return nil, consensus_types.ErrNilObjectWrapped
	}
	return w, nil
}

// IsNil checks if the underlying data is nil.
func (e executionPayloadEpbs) IsNil() bool {
	return e.p == nil
}

// MarshalSSZ --
func (e executionPayloadEpbs) MarshalSSZ() ([]byte, error) {
	return e.p.MarshalSSZ()
}

// MarshalSSZTo --
func (e executionPayloadEpbs) MarshalSSZTo(dst []byte) ([]byte, error) {
	return e.p.MarshalSSZTo(dst)
}

// SizeSSZ --
func (e executionPayloadEpbs) SizeSSZ() int {
	return e.p.SizeSSZ()
}

// UnmarshalSSZ --
func (e executionPayloadEpbs) UnmarshalSSZ(buf []byte) error {
	return e.p.UnmarshalSSZ(buf)
}

// HashTreeRoot --
func (e executionPayloadEpbs) HashTreeRoot() ([32]byte, error) {
	return e.p.HashTreeRoot()
}

// HashTreeRootWith --
func (e executionPayloadEpbs) HashTreeRootWith(hh *fastssz.Hasher) error {
	return e.p.HashTreeRootWith(hh)
}

// Proto --
func (e executionPayloadEpbs) Proto() proto.Message {
	return e.p
}

// ParentHash --
func (e executionPayloadEpbs) ParentHash() []byte {
	return e.p.ParentHash
}

// FeeRecipient --
func (e executionPayloadEpbs) FeeRecipient() []byte {
	return e.p.FeeRecipient
}

// StateRoot --
func (e executionPayloadEpbs) StateRoot() []byte {
	return e.p.StateRoot
}

// ReceiptsRoot --
func (e executionPayloadEpbs) ReceiptsRoot() []byte {
	return e.p.ReceiptsRoot
}

// LogsBloom --
func (e executionPayloadEpbs) LogsBloom() []byte {
	return e.p.LogsBloom
}

// PrevRandao --
func (e executionPayloadEpbs) PrevRandao() []byte {
	return e.p.PrevRandao
}

// BlockNumber --
func (e executionPayloadEpbs) BlockNumber() uint64 {
	return e.p.BlockNumber
}

// GasLimit --
func (e executionPayloadEpbs) GasLimit() uint64 {
	return e.p.GasLimit
}

// GasUsed --
func (e executionPayloadEpbs) GasUsed() uint64 {
	return e.p.GasUsed
}

// Timestamp --
func (e executionPayloadEpbs) Timestamp() uint64 {
	return e.p.Timestamp
}

// ExtraData --
func (e executionPayloadEpbs) ExtraData() []byte {
	return e.p.ExtraData
}

// BaseFeePerGas --
func (e executionPayloadEpbs) BaseFeePerGas() []byte {
	return e.p.BaseFeePerGas
}

// BlockHash --
func (e executionPayloadEpbs) BlockHash() []byte {
	return e.p.BlockHash
}

// Transactions --
func (e executionPayloadEpbs) Transactions() ([][]byte, error) {
	return e.p.Transactions, nil
}

// TransactionsRoot --
func (e executionPayloadEpbs) TransactionsRoot() ([]byte, error) {
	return nil, consensus_types.ErrUnsupportedField
}

// Withdrawals --
func (e executionPayloadEpbs) Withdrawals() ([]*enginev1.Withdrawal, error) {
	return e.p.Withdrawals, nil
}

// WithdrawalsRoot --
func (e executionPayloadEpbs) WithdrawalsRoot() ([]byte, error) {
	return nil, consensus_types.ErrUnsupportedField
}

func (e executionPayloadEpbs) BlobGasUsed() (uint64, error) {
	return e.p.BlobGasUsed, nil
}

func (e executionPayloadEpbs) ExcessBlobGas() (uint64, error) {
	return e.p.ExcessBlobGas, nil
}

// PbBellatrix --
func (e executionPayloadEpbs) PbBellatrix() (*enginev1.ExecutionPayload, error) {
	return nil, consensus_types.ErrUnsupportedField
}

// PbCapella --
func (e executionPayloadEpbs) PbCapella() (*enginev1.ExecutionPayloadCapella, error) {
	return nil, consensus_types.ErrUnsupportedField
}

// PbDeneb --
func (e executionPayloadEpbs) PbDeneb() (*enginev1.ExecutionPayloadDeneb, error) {
	return nil, consensus_types.ErrUnsupportedField
}

// ValueInWei --
func (e executionPayloadEpbs) ValueInWei() (math.Wei, error) {
	return nil, consensus_types.ErrUnsupportedField
}

// ValueInGwei --
func (e executionPayloadEpbs) ValueInGwei() (uint64, error) {
	return 0, consensus_types.ErrUnsupportedField
}

// IsBlinded --
func (e executionPayloadEpbs) IsBlinded() bool {
	return false
}
