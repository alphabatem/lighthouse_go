package lighthouse_go

import (
	"github.com/alphabatem/common/context"
	"github.com/alphabatem/lighthouse_go/generated/lighthouse"
	solana "github.com/fluxrpc/solana-go"
)

type LighthouseService struct {
	context.DefaultService

	logLevel lighthouse.LogLevel
}

const LIGHTHOUSE_SVC = "lighthouse_svc"

func (svc LighthouseService) Id() string {
	return LIGHTHOUSE_SVC
}

func (svc *LighthouseService) Start() error {
	svc.logLevel = lighthouse.LogLevel_FailedPlaintextMessage

	return nil
}

func (svc *LighthouseService) SetLogLevel(level lighthouse.LogLevel) {
	svc.logLevel = level
}

func (svc *LighthouseService) AssertTokenAccountAmountMinMaxInstruction(tokenAccount solana.PublicKey, minAmount uint64, maxAmount uint64) (solana.Instruction, error) {
	ix := lighthouse.NewAssertTokenAccountMultiInstruction(svc.logLevel, []*lighthouse.TokenAccountAssertion{
		{Amount: &minAmount, Operator: uint8(lighthouse.IntegerOperator_GreaterThanOrEqual)},
		{Amount: &maxAmount, Operator: uint8(lighthouse.IntegerOperator_LessThanOrEqual)},
	}, tokenAccount)

	return ix.Build(), nil
}

func (svc *LighthouseService) AssertTokenAccountAmountInstruction(tokenAccount solana.PublicKey, amount uint64, operator lighthouse.IntegerOperator) (solana.Instruction, error) {
	ix := lighthouse.NewAssertTokenAccountMultiInstruction(svc.logLevel, []*lighthouse.TokenAccountAssertion{
		{Amount: &amount, Operator: uint8(operator)},
	}, tokenAccount)

	return ix.Build(), nil
}

func (svc *LighthouseService) AssertAccountInfoInstruction(account solana.PublicKey, assert []*lighthouse.AccountInfoAssertion) (solana.Instruction, error) {
	ix := lighthouse.NewAssertAccountInfoMultiInstruction(svc.logLevel, assert, account)
	return ix.Build(), nil
}

func (svc *LighthouseService) AssertTokenAccountInstruction(tokenAccount solana.PublicKey, assert []*lighthouse.TokenAccountAssertion) (solana.Instruction, error) {
	ix := lighthouse.NewAssertTokenAccountMultiInstruction(svc.logLevel, assert, tokenAccount)

	return ix.Build(), nil
}
