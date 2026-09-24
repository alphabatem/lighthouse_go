package lighthouse_go

import (
	"testing"

	"github.com/alphabatem/lighthouse_go/generated/lighthouse"
	solana "github.com/fluxrpc/solana-go"
	"github.com/stretchr/testify/require"
)

func TestFluxInstructionTransaction(t *testing.T) {
	svc := LighthouseService{}
	require.NoError(t, svc.Start())
	svc.SetLogLevel(lighthouse.LogLevel(0))
	account := solana.PublicKey{1}
	payer := solana.PublicKey{2}
	ix, err := svc.AssertTokenAccountAmountInstruction(account, 42, lighthouse.IntegerOperator_Equal)
	require.NoError(t, err)
	require.Equal(t, lighthouse.ProgramID, ix.ProgramID())
	require.Equal(t, []*solana.AccountMeta{{PublicKey: account}}, ix.Accounts())
	data, err := ix.Data()
	require.NoError(t, err)
	// Instruction 10, log level 0, one amount assertion, little-endian 42, equal.
	require.Equal(t, []byte{10, 0, 1, 2, 42, 0, 0, 0, 0, 0, 0, 0, 0}, data)
	decodedInstruction, err := lighthouse.DecodeInstruction(ix.Accounts(), data)
	require.NoError(t, err)
	require.Equal(t, ix.Accounts(), decodedInstruction.Accounts())
	require.Equal(t, uint64(42), *decodedInstruction.Impl.(*lighthouse.AssertTokenAccountMulti).Assertions[0].Amount)

	tx, err := solana.NewTransaction([]solana.Instruction{ix}, solana.Hash{3}, solana.TransactionPayer(payer))
	require.NoError(t, err)
	wire, err := tx.MarshalBinary()
	require.NoError(t, err)
	decoded, err := solana.TransactionFromBytes(wire)
	require.NoError(t, err)
	require.Len(t, decoded.Message.Instructions, 1)
	compiled := decoded.Message.Instructions[0]
	require.Equal(t, data, []byte(compiled.Data))
	require.Equal(t, lighthouse.ProgramID, decoded.Message.AccountKeys[compiled.ProgramIDIndex])
	require.Len(t, compiled.Accounts, 1)
	require.Equal(t, account, decoded.Message.AccountKeys[compiled.Accounts[0]])
}
