package lighthouse

import solana "github.com/fluxrpc/solana-go"

func fixturePtr[T any](value T) *T { return &value }

func binaryInstructionFixtures() map[string]*Instruction {
	return map[string]*Instruction{
		"AssertAccountData":                   (&AssertAccountData{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertion: fixturePtr(AccountDataAssertion{Offset: 257, Assertion: []byte{1, 2, 3}})}).Build(),
		"AssertAccountDataMulti":              (&AssertAccountDataMulti{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertions: fixturePtr(AccountDataAssertions{{Offset: 257, Assertion: []byte{1, 2, 3}}})}).Build(),
		"AssertAccountDelta":                  (&AssertAccountDelta{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertion: fixturePtr(AccountDeltaAssertion{Typ: 2, Data: []byte{3, 4, 5}})}).Build(),
		"AssertAccountInfo":                   (&AssertAccountInfo{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertion: fixturePtr(AccountInfoAssertion{Owner: fixturePtr(solana.PublicKey{1, 2, 3}), Operator: 1})}).Build(),
		"AssertAccountInfoMulti":              (&AssertAccountInfoMulti{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertions: AccountInfoAssertions{&AccountInfoAssertion{Owner: fixturePtr(solana.PublicKey{1, 2, 3}), Operator: 1}}}).Build(),
		"AssertBubblegumTreeConfigAccount":    (&AssertBubblegumTreeConfigAccount{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertion: fixturePtr(BubblegumTreeConfigAssertion{Typ: 2, Data: []byte{3, 4, 5}})}).Build(),
		"AssertMerkleTreeAccount":             (&AssertMerkleTreeAccount{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertion: fixturePtr(MerkleTreeAssertion{Typ: 2, Data: []byte{3, 4, 5}})}).Build(),
		"AssertMintAccount":                   (&AssertMintAccount{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertion: fixturePtr(MintAccountAssertion{Typ: 2, Data: []byte{3, 4, 5}})}).Build(),
		"AssertMintAccountMulti":              (&AssertMintAccountMulti{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertions: fixturePtr(MintAccountAssertions{Typ: 2, Data: []byte{3, 4, 5}})}).Build(),
		"AssertStakeAccount":                  (&AssertStakeAccount{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertion: fixturePtr(StakeAccountAssertion{Typ: 2, Data: []byte{3, 4, 5}})}).Build(),
		"AssertStakeAccountMulti":             (&AssertStakeAccountMulti{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertions: fixturePtr(StakeAccountAssertions{Typ: 2, Data: []byte{3, 4, 5}})}).Build(),
		"AssertSysvarClock":                   (&AssertSysvarClock{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertion: fixturePtr(SysvarClockAssertion{Typ: 2, Data: []byte{3, 4, 5}})}).Build(),
		"AssertTokenAccount":                  (&AssertTokenAccount{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertion: fixturePtr(TokenAccountAssertion{Amount: fixturePtr(uint64(42)), Operator: 4})}).Build(),
		"AssertTokenAccountMulti":             (&AssertTokenAccountMulti{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertions: TokenAccountAssertions{&TokenAccountAssertion{Amount: fixturePtr(uint64(42)), Operator: 4}}}).Build(),
		"AssertUpgradeableLoaderAccount":      (&AssertUpgradeableLoaderAccount{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertion: fixturePtr(UpgradeableLoaderStateAssertion{Typ: 2, Data: []byte{3, 4, 5}})}).Build(),
		"AssertUpgradeableLoaderAccountMulti": (&AssertUpgradeableLoaderAccountMulti{LogLevel: fixturePtr(LogLevel_FailedPlaintextMessage), Assertions: fixturePtr(UpgradeableLoaderStateAssertions{Typ: 2, Data: []byte{3, 4, 5}})}).Build(),
		"MemoryClose":                         (&MemoryClose{MemoryId: fixturePtr(uint8(7)), MemoryBump: fixturePtr(uint8(7))}).Build(),
		"MemoryWrite":                         (&MemoryWrite{MemoryId: fixturePtr(uint8(7)), MemoryBump: fixturePtr(uint8(7)), WriteOffset: fixturePtr(CompactU64(257)), WriteType: fixturePtr(WriteType{Typ: 2, Data: []byte{3, 4, 5}})}).Build(),
	}
}
