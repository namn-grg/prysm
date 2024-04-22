package blocks_test

import (
	"testing"

	ssz "github.com/prysmaticlabs/fastssz"
	fieldparams "github.com/prysmaticlabs/prysm/v5/config/fieldparams"
	consensus_types "github.com/prysmaticlabs/prysm/v5/consensus-types"
	"github.com/prysmaticlabs/prysm/v5/consensus-types/blocks"
	"github.com/prysmaticlabs/prysm/v5/consensus-types/interfaces"
	"github.com/prysmaticlabs/prysm/v5/encoding/bytesutil"
	enginev1 "github.com/prysmaticlabs/prysm/v5/proto/engine/v1"
	"github.com/prysmaticlabs/prysm/v5/testing/require"
)

func createWrappedSignedPayloadEnvelope(t testing.TB) interfaces.ExecutionData {
	payload, err := blocks.WrappedExecutionPayloadEpbs(&enginev1.ExecutionPayloadEPBS{
		ParentHash:    bytesutil.PadTo([]byte("parentblockhash"), fieldparams.RootLength),
		FeeRecipient:  bytesutil.PadTo([]byte("feerecipient"), fieldparams.FeeRecipientLength),
		StateRoot:     bytesutil.PadTo([]byte("stateroot"), fieldparams.RootLength),
		ReceiptsRoot:  bytesutil.PadTo([]byte("receiptsroot"), fieldparams.RootLength),
		LogsBloom:     bytesutil.PadTo([]byte("logsbloom"), fieldparams.LogsBloomLength),
		PrevRandao:    bytesutil.PadTo([]byte("prevrandao"), fieldparams.RootLength),
		BlockNumber:   1,
		GasLimit:      2,
		GasUsed:       3,
		Timestamp:     4,
		ExtraData:     bytesutil.PadTo([]byte("extradata"), fieldparams.RootLength),
		BaseFeePerGas: bytesutil.PadTo([]byte("basefeepergas"), fieldparams.RootLength),
		BlockHash:     bytesutil.PadTo([]byte("blockhash"), fieldparams.RootLength),
		Transactions:  [][]byte{{0xa}, {0xb}, {0xc}},
		Withdrawals:   make([]*enginev1.Withdrawal, 0),
		BlobGasUsed:   5,
		ExcessBlobGas: 6,
		InclusionListSummary: [][]byte{
			bytesutil.PadTo([]byte("alice"), fieldparams.FeeRecipientLength),
			bytesutil.PadTo([]byte("blob"), fieldparams.FeeRecipientLength),
			bytesutil.PadTo([]byte("charlie"), fieldparams.FeeRecipientLength),
		},
	})
	require.NoError(t, err)
	return payload
}

func TestWrappedSignedExecutionPayloadEnvelope(t *testing.T) {
	p := createWrappedSignedPayloadEnvelope(t)
	require.Equal(t, false, p.IsNil())
	m, err := p.MarshalSSZ()
	require.NoError(t, err)
	_, err = p.MarshalSSZTo(nil)
	require.NoError(t, err)
	n := p.SizeSSZ()
	require.NotEqual(t, 0, n)
	_, err = p.HashTreeRoot()
	require.NoError(t, err)
	err = p.HashTreeRootWith(ssz.DefaultHasherPool.Get())
	require.NoError(t, err)
	proto := p.Proto()
	require.NotNil(t, proto)
	parentHash := p.ParentHash()
	require.DeepEqual(t, parentHash, bytesutil.PadTo([]byte("parentblockhash"), fieldparams.RootLength))
	feeRecipient := p.FeeRecipient()
	require.DeepEqual(t, feeRecipient, bytesutil.PadTo([]byte("feerecipient"), fieldparams.FeeRecipientLength))
	stateRoot := p.StateRoot()
	require.DeepEqual(t, stateRoot, bytesutil.PadTo([]byte("stateroot"), fieldparams.RootLength))
	receiptsRoot := p.ReceiptsRoot()
	require.DeepEqual(t, receiptsRoot, bytesutil.PadTo([]byte("receiptsroot"), fieldparams.RootLength))
	logsBloom := p.LogsBloom()
	require.DeepEqual(t, logsBloom, bytesutil.PadTo([]byte("logsbloom"), fieldparams.LogsBloomLength))
	prevRandao := p.PrevRandao()
	require.DeepEqual(t, prevRandao, bytesutil.PadTo([]byte("prevrandao"), fieldparams.RootLength))
	extraData := p.ExtraData()
	require.DeepEqual(t, extraData, bytesutil.PadTo([]byte("extradata"), fieldparams.RootLength))
	baseFeePerGas := p.BaseFeePerGas()
	require.DeepEqual(t, baseFeePerGas, bytesutil.PadTo([]byte("basefeepergas"), fieldparams.RootLength))
	require.Equal(t, uint64(1), p.BlockNumber())
	require.Equal(t, uint64(2), p.GasLimit())
	require.Equal(t, uint64(3), p.GasUsed())
	require.Equal(t, uint64(4), p.Timestamp())
	require.DeepEqual(t, p.BlockHash(), bytesutil.PadTo([]byte("blockhash"), fieldparams.RootLength))
	txs, err := p.Transactions()
	require.NoError(t, err)
	require.DeepEqual(t, txs, [][]byte{{0xa}, {0xb}, {0xc}})
	_, err = p.TransactionsRoot()
	require.ErrorIs(t, err, consensus_types.ErrUnsupportedField)
	wd, err := p.Withdrawals()
	require.NoError(t, err)
	require.DeepEqual(t, wd, make([]*enginev1.Withdrawal, 0))
	_, err = p.WithdrawalsRoot()
	require.ErrorIs(t, err, consensus_types.ErrUnsupportedField)
	b, err := p.BlobGasUsed()
	require.NoError(t, err)
	require.DeepEqual(t, uint64(5), b)
	e, err := p.ExcessBlobGas()
	require.NoError(t, err)
	require.DeepEqual(t, uint64(6), e)
	_, err = p.PbBellatrix()
	require.ErrorIs(t, err, consensus_types.ErrUnsupportedField)
	_, err = p.PbCapella()
	require.ErrorIs(t, err, consensus_types.ErrUnsupportedField)
	_, err = p.PbDeneb()
	require.ErrorIs(t, err, consensus_types.ErrUnsupportedField)
	_, err = p.ValueInGwei()
	require.ErrorIs(t, err, consensus_types.ErrUnsupportedField)
	_, err = p.ValueInWei()
	require.ErrorIs(t, err, consensus_types.ErrUnsupportedField)
	require.Equal(t, false, p.IsBlinded())
	require.NoError(t, p.UnmarshalSSZ(m)) // Testing this last because it modifies the object.
}
