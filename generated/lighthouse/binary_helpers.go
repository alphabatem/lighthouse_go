package lighthouse

import (
	"fmt"

	"github.com/fluxrpc/solana-go/binary"
)

// These helpers retain the Borsh byte-vector layout used by the bindings.
func writeByteSlice(encoder *binary.Encoder, data []byte) error {
	if uint64(len(data)) > uint64(^uint32(0)) {
		return fmt.Errorf("byte vector is too large: %d", len(data))
	}
	encoder.WriteUint32(uint32(len(data)))
	encoder.WriteBytes(data)
	return encoder.Err()
}

func readByteSlice(decoder *binary.Decoder) []byte {
	n := decoder.ReadUint32()
	return decoder.ReadBytesCopy(int(n))
}
