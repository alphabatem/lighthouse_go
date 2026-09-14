package lighthouse_go

import (
	"context"
	"github.com/alphabatem/lighthouse_go/generated/lighthouse"
	solana "github.com/fluxrpc/solana-go"
	bin "github.com/fluxrpc/solana-go/binary"
	"github.com/fluxrpc/solana-go/rpc"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func integrationRPC(t *testing.T, needsKeypair bool) *rpc.Client {
	t.Helper()
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}
	if os.Getenv("RPC_URL") == "" {
		t.Skip("set RPC_URL to run integration tests")
	}
	if needsKeypair && os.Getenv("TEST_KEYPAIR") == "" {
		t.Skip("set TEST_KEYPAIR to run simulation tests")
	}
	return rpc.New(os.Getenv("RPC_URL"))
}

func TestLighthouseService_AssertTokenAccountInstruction_Decode(t *testing.T) {
	sig := solana.MustSignatureFromBase58("2CJ8oLPbtYmjtcXMjn1u89QVX3EJWJDzjR1HQWw1nDSqQpHWuLQhbecDMSRXXg9bkSWFy8AmckE7Ue4RpjTddQtd")

	c := integrationRPC(t, false)

	z := uint64(0)
	txn, err := c.GetTransactionWithOpts(context.TODO(), sig, &rpc.GetTransactionOpts{Commitment: rpc.CommitmentConfirmed, MaxSupportedTransactionVersion: &z})
	if err != nil {
		t.Fatal(err)
	}

	tx, err := txn.Transaction.GetTransaction()
	if err != nil {
		t.Fatal(err)
	}

	tokenAtaAssertIx := tx.Message.Instructions[3]

	pk := solana.PublicKeyFromBytes(tokenAtaAssertIx.Data[len(tokenAtaAssertIx.Data)-33 : len(tokenAtaAssertIx.Data)-1])

	var dix lighthouse.AssertTokenAccountMulti
	err = dix.UnmarshalWithDecoder(bin.NewDecoder(tokenAtaAssertIx.Data[1:]))
	if err != nil {
		t.Fatal(err)
	}

	expected := uint64(971124800672)
	//op := ">="
	if *dix.Assertions[0].Amount != expected {
		t.Fatal()
	}
	if dix.Assertions[0].Operator != 4 {
		t.Fatal()
	}

	pk = solana.PublicKey{}
	if *dix.Assertions[1].Delegate != pk {
		t.Fatal()
	}
	if dix.Assertions[1].Operator != 0 {
		t.Fatal()
	}

	if *dix.Assertions[2].DelegatedAmount != 0 {
		t.Fatal()
	}
	if dix.Assertions[2].Operator != 4 {
		t.Fatal()
	}

	if dix.Assertions[3].Owner.String() != "D4m7Wj3HQ1UmrPdGtzUc1e6gy74LWr6AcLoT2HCqW1jD" {
		t.Fatal()
	}
	if dix.Assertions[3].Operator != 0 {
		t.Fatal()
	}
}

func TestLighthouseService_AssertTokenAccountAmountInstruction(t *testing.T) {
	c := integrationRPC(t, true)

	hash, err := c.GetLatestBlockhash(context.TODO(), rpc.CommitmentConfirmed)
	if err != nil {
		t.Fatal(err)
	}

	svc := LighthouseService{}
	if err := svc.Start(); err != nil {
		t.Fatal(err)
	}

	kp, err := solana.PrivateKeyFromSolanaKeygenFile(os.Getenv("TEST_KEYPAIR"))
	if err != nil {
		t.Fatal(err)
	}

	tokenAta := solana.MustPublicKeyFromBase58("8xdxXqGnMWqWLNP6xRsJBynEHk3t1o577jap1UNiHdxf")

	ix, err := svc.AssertTokenAccountAmountInstruction(tokenAta, 18827277720, lighthouse.IntegerOperator_Equal)
	if err != nil {
		t.Fatal(err)
	}

	d, _ := ix.Data()
	t.Logf("Data: %v", d)

	txn, err := solana.NewTransaction([]solana.Instruction{
		NewComputeBudgetSetUnitPriceInstruction(200),
		NewComputeBudgetSetUnitLimitInstruction(10_000),
		ix,
	}, hash.Value.Blockhash, solana.TransactionPayer(kp.PublicKey()))
	if err != nil {
		t.Fatal(err)
	}

	txn.Signatures = []solana.Signature{solana.Signature{}}

	sim, err := c.SimulateTransactionWithOpts(context.TODO(), txn, &rpc.SimulateTransactionOpts{
		SigVerify:  false,
		Commitment: rpc.CommitmentProcessed,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Sim: %+v\n", sim.Value)

	if sim.Value.Err != nil {
		t.Fatal(sim.Value.Err)
	}
}

func TestLighthouseService_AssertAccountInfoInstruction(t *testing.T) {
	c := integrationRPC(t, true)

	hash, err := c.GetLatestBlockhash(context.TODO(), rpc.CommitmentConfirmed)
	if err != nil {
		t.Fatal(err)
	}

	svc := LighthouseService{}
	if err := svc.Start(); err != nil {
		t.Fatal(err)
	}

	kp, err := solana.PrivateKeyFromSolanaKeygenFile(os.Getenv("TEST_KEYPAIR"))
	if err != nil {
		t.Fatal(err)
	}

	accInfo, err := c.GetAccountInfo(context.TODO(), kp.PublicKey())
	if err != nil {
		t.Fatal(err)
	}

	lamports := accInfo.Value.Lamports - 5002 //Remove tx cost & gas

	t.Logf("Asserting lamports == %v", lamports)
	ix, err := svc.AssertAccountInfoInstruction(kp.PublicKey(), lighthouse.AccountInfoAssertions{
		{
			Lamports: &lamports,
			Operator: uint8(lighthouse.IntegerOperator_Equal),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	d, _ := ix.Data()
	t.Logf("Data: %v", d)

	txn, err := solana.NewTransaction([]solana.Instruction{
		NewComputeBudgetSetUnitPriceInstruction(200),
		NewComputeBudgetSetUnitLimitInstruction(10_000),
		ix,
	}, hash.Value.Blockhash, solana.TransactionPayer(kp.PublicKey()))
	if err != nil {
		t.Fatal(err)
	}

	txn.Signatures = []solana.Signature{solana.Signature{}}
	//_, err = txn.Sign(func(key solana.PublicKey) *solana.PrivateKey {
	//	return &kp
	//})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//sig, err := c.SendTransactionWithOpts(context.TODO(), txn, rpc.TransactionOpts{SkipPreflight: true})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//log.Printf("Sig: %s", sig)

	sim, err := c.SimulateTransactionWithOpts(context.TODO(), txn, &rpc.SimulateTransactionOpts{
		SigVerify:  false,
		Commitment: rpc.CommitmentProcessed,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Sim: %+v\n", sim.Value)

	if sim.Value.Err != nil {
		t.Fatal(sim.Value.Err)
	}
}

func NewComputeBudgetSetUnitLimitInstruction(units uint32) solana.Instruction {
	computeBudget := solana.MustPublicKeyFromBase58("ComputeBudget111111111111111111111111111111")

	borshEncoder := bin.NewEncoder(nil)

	borshEncoder.WriteUint8(2) //2 = Set Compute Unit Limit
	borshEncoder.WriteUint32(units)

	inst2 := solana.NewInstruction(computeBudget, solana.AccountMetaSlice{}, borshEncoder.Bytes())

	return inst2
}

func NewComputeBudgetSetUnitPriceInstruction(mLamports uint64) solana.Instruction {
	computeBudget := solana.MustPublicKeyFromBase58("ComputeBudget111111111111111111111111111111")

	borshEncoder := bin.NewEncoder(nil)

	borshEncoder.WriteUint8(3) //3 = Set Compute Unit Bids
	borshEncoder.WriteUint64(mLamports)

	inst2 := solana.NewInstruction(computeBudget, solana.AccountMetaSlice{}, borshEncoder.Bytes())

	return inst2
}
