// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package lighthouse

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AssertAccountDelta is the `AssertAccountDelta` instruction.
type AssertAccountDelta struct {
	LogLevel  *LogLevel
	Assertion *AccountDeltaAssertion

	// [0] = [] accountA
	// ··········· Account A where the delta is calculated from
	//
	// [1] = [] accountB
	// ··········· Account B where the delta is calculated to
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAssertAccountDeltaInstructionBuilder creates a new `AssertAccountDelta` instruction builder.
func NewAssertAccountDeltaInstructionBuilder() *AssertAccountDelta {
	nd := &AssertAccountDelta{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetLogLevel sets the "logLevel" parameter.
func (inst *AssertAccountDelta) SetLogLevel(logLevel LogLevel) *AssertAccountDelta {
	inst.LogLevel = &logLevel
	return inst
}

// SetAssertion sets the "assertion" parameter.
func (inst *AssertAccountDelta) SetAssertion(assertion AccountDeltaAssertion) *AssertAccountDelta {
	inst.Assertion = &assertion
	return inst
}

// SetAccountAAccount sets the "accountA" account.
// Account A where the delta is calculated from
func (inst *AssertAccountDelta) SetAccountAAccount(accountA ag_solanago.PublicKey) *AssertAccountDelta {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(accountA)
	return inst
}

// GetAccountAAccount gets the "accountA" account.
// Account A where the delta is calculated from
func (inst *AssertAccountDelta) GetAccountAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAccountBAccount sets the "accountB" account.
// Account B where the delta is calculated to
func (inst *AssertAccountDelta) SetAccountBAccount(accountB ag_solanago.PublicKey) *AssertAccountDelta {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(accountB)
	return inst
}

// GetAccountBAccount gets the "accountB" account.
// Account B where the delta is calculated to
func (inst *AssertAccountDelta) GetAccountBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst AssertAccountDelta) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AssertAccountDelta,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AssertAccountDelta) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AssertAccountDelta) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.LogLevel == nil {
			return errors.New("LogLevel parameter is not set")
		}
		if inst.Assertion == nil {
			return errors.New("Assertion parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AccountA is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AccountB is not set")
		}
	}
	return nil
}

func (inst *AssertAccountDelta) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AssertAccountDelta")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" LogLevel", *inst.LogLevel))
						paramsBranch.Child(ag_format.Param("Assertion", *inst.Assertion))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("accountA", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("accountB", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj AssertAccountDelta) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LogLevel` param:
	err = encoder.Encode(obj.LogLevel)
	if err != nil {
		return err
	}
	// Serialize `Assertion` param:
	err = encoder.Encode(obj.Assertion)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AssertAccountDelta) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LogLevel`:
	err = decoder.Decode(&obj.LogLevel)
	if err != nil {
		return err
	}
	// Deserialize `Assertion`:
	err = decoder.Decode(&obj.Assertion)
	if err != nil {
		return err
	}
	return nil
}

// NewAssertAccountDeltaInstruction declares a new AssertAccountDelta instruction with the provided parameters and accounts.
func NewAssertAccountDeltaInstruction(
	// Parameters:
	logLevel LogLevel,
	assertion AccountDeltaAssertion,
	// Accounts:
	accountA ag_solanago.PublicKey,
	accountB ag_solanago.PublicKey) *AssertAccountDelta {
	return NewAssertAccountDeltaInstructionBuilder().
		SetLogLevel(logLevel).
		SetAssertion(assertion).
		SetAccountAAccount(accountA).
		SetAccountBAccount(accountB)
}
