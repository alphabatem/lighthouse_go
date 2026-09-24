package lighthouse

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"testing"

	solana "github.com/fluxrpc/solana-go"
	"github.com/fluxrpc/solana-go/binary"
	"github.com/stretchr/testify/require"
)

func TestBinaryInstructionCompatibility(t *testing.T) {
	// Golden bytes were captured from the original bindings before migration.
	raw, err := os.ReadFile("testdata/instructions.golden")
	require.NoError(t, err)
	var golden map[string]string
	require.NoError(t, json.Unmarshal(raw, &golden))
	fixtures := binaryInstructionFixtures()
	require.Len(t, golden, len(fixtures))
	for name, inst := range fixtures {
		t.Run(name, func(t *testing.T) {
			data, err := inst.Data()
			require.NoError(t, err)
			require.Equal(t, golden[name], hex.EncodeToString(data))
			decoded, err := DecodeInstruction(nil, data)
			require.NoError(t, err)
			require.Equal(t, inst.TypeID, decoded.TypeID)
			require.Equal(t, inst.Impl, decoded.Impl)
			for end := 0; end < len(data); end++ {
				_, err := DecodeInstruction(nil, data[:end])
				require.Error(t, err, "truncated at byte %d", end)
			}
		})
	}
}

func TestBinaryAccountTypes(t *testing.T) {
	want := TestAccountV1{
		U8: 255, I8: -128, U16: 65535, I16: -32768,
		U32: 0xffffffff, I32: -2147483648, U64: ^uint64(0), I64: -9223372036854775808,
		U128: binary.Uint128{Lo: 1, Hi: 2}, I128: binary.Int128FromInt64(-1),
		Bytes: [32]byte{1, 2, 3}, TrueField: true,
		OptionU8: fixturePtr(uint8(7)), OptionU16: fixturePtr(uint16(258)),
		Pubkey: solana.PublicKey{4, 5, 6}, Vec: []byte{7, 8, 9},
	}
	encoder := binary.NewEncoder(nil)
	require.NoError(t, want.MarshalWithEncoder(encoder))
	data := encoder.Bytes()
	require.Equal(t, "01000000000000000200000000000000ffffffffffffffffffffffffffffffff", hex.EncodeToString(data[30:62]))
	// Decoding must clear absent options even when the destination is reused.
	got := TestAccountV1{OptionU8None: fixturePtr(uint8(9)), OptionU16None: fixturePtr(uint16(9))}
	require.NoError(t, got.UnmarshalWithDecoder(binary.NewDecoder(data)))
	require.Equal(t, want, got)
	for end := 0; end < len(data); end++ {
		var truncated TestAccountV1
		require.Error(t, truncated.UnmarshalWithDecoder(binary.NewDecoder(data[:end])))
	}
}

func TestBinaryAssertionVariants(t *testing.T) {
	key := solana.PublicKey{1, 2, 3}
	token := []TokenAccountAssertion{
		{Mint: &key}, {Owner: &key}, {Amount: fixturePtr(uint64(42))},
		{Delegate: &key}, {State: fixturePtr([]byte{2})}, {IsNative: fixturePtr(true)},
		{DelegatedAmount: fixturePtr(uint64(42))}, {CloseAuthority: &key},
	}
	for tag, want := range token {
		encoder := binary.NewEncoder(nil)
		require.NoError(t, want.MarshalWithEncoder(encoder))
		require.Equal(t, uint8(tag), encoder.Bytes()[0])
		got := TokenAccountAssertion{Amount: fixturePtr(uint64(99))}
		require.NoError(t, got.UnmarshalWithDecoder(binary.NewDecoder(encoder.Bytes())))
		require.Equal(t, want, got)
	}
	info := []AccountInfoAssertion{
		{Lamports: fixturePtr(uint64(42))}, {DataLength: fixturePtr(uint64(42))},
		{Owner: &key}, {RentEpoch: fixturePtr(uint64(42))},
		{IsSigner: fixturePtr(true)}, {IsWritable: fixturePtr(false)}, {Executable: fixturePtr(true)},
	}
	for _, want := range info {
		encoder := binary.NewEncoder(nil)
		require.NoError(t, want.MarshalWithEncoder(encoder))
		got := AccountInfoAssertion{Owner: &key}
		require.NoError(t, got.UnmarshalWithDecoder(binary.NewDecoder(encoder.Bytes())))
		require.Equal(t, want, got)
	}
}

func TestBinaryInvalidData(t *testing.T) {
	for _, data := range [][]byte{
		{255},                         // Unknown instruction.
		{9, 0, 255},                   // Unknown token assertion.
		{5, 0, 5, 2, 0},               // Invalid boolean.
		{10, 0, 255},                  // Assertion count exceeds available bytes.
		{3, 0, 255, 255, 255, 255},    // Oversized assertion vector.
		{7, 0, 0, 255, 255, 255, 255}, // Oversized byte vector.
	} {
		_, err := DecodeInstruction(nil, data)
		require.Error(t, err)
	}
	encoder := binary.NewEncoder(nil)
	require.Error(t, (TokenAccountAssertion{}).MarshalWithEncoder(encoder))
	require.Error(t, (AccountInfoAssertion{Lamports: fixturePtr(uint64(1)), RentEpoch: fixturePtr(uint64(2))}).MarshalWithEncoder(encoder))
	require.Error(t, (TokenAccountAssertions{nil}).MarshalWithEncoder(encoder))
	require.Error(t, make(TokenAccountAssertions, 256).MarshalWithEncoder(encoder))
	require.Error(t, make(AccountInfoAssertions, 256).MarshalWithEncoder(encoder))
}
