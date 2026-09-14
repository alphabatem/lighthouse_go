# [WIP] lighthouse_go

Uses `github.com/fluxrpc/solana-go` v0.1.9 and requires Go 1.26.4 or newer.

Binary serialization uses `github.com/fluxrpc/solana-go/binary`. Create codecs
with `binary.NewEncoder(nil)` and `binary.NewDecoder(data)`, then call the
bindings' `MarshalWithEncoder` and `UnmarshalWithDecoder` methods. Instruction
payload decoders consume parameters only; `lighthouse.DecodeInstruction` consumes
the full instruction, including its one-byte discriminator. Instruction IDs and
`BaseVariant.TypeID` are now `uint8`; explicit dispatch replaces the old variant
registry. Golden tests cover byte compatibility for all 18 instruction types.

The service and generated bindings accept Flux public keys and return Flux
instructions. Use `lighthouse.DecodeInstruction` directly; Flux does not provide
the previous SDK's global instruction decoder registry. `EncodeToTree` remains
available with plain-text labels; the previous SDK-specific `TextEncode` method
has been removed.

The bindings were adapted from anchor-go output. When regenerating from
`lighthouse.json`, preserve the Flux imports, explicit binary codecs, instruction
dispatch, account accessors, and local tree formatting helpers.

Run `go test ./...` for tests. Live RPC tests require `RPC_URL`; simulation tests
also require `TEST_KEYPAIR` pointing to a Solana keypair file. These can be set in
the environment or `.env`; tests skip when the required configuration is absent.
The offline Flux transaction test is `go test . -run TestFluxInstructionTransaction`.

### Progress
- [x] Token Account Multi
