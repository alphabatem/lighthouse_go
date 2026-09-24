package lighthouse

import (
	"reflect"

	fuzz "github.com/gagliardetto/gofuzz"
)

func assertionFuzzer() *fuzz.Fuzzer {
	// Assertions are tagged unions: only one value field may be populated.
	keepOne := func(value interface{}, c fuzz.Continue) {
		c.FuzzNoCustom(value)
		fields := reflect.ValueOf(value).Elem()
		selected := c.Intn(fields.NumField() - 1) // The final field is Operator.
		for i := 0; i < fields.NumField()-1; i++ {
			if i != selected {
				fields.Field(i).SetZero()
			}
		}
	}
	return fuzz.New().NilChance(0).Funcs(
		func(value *TokenAccountAssertion, c fuzz.Continue) {
			keepOne(value, c)
			if value.State != nil {
				*value.State = []byte{uint8(c.Intn(3))}
			}
		},
		func(value *AccountInfoAssertion, c fuzz.Continue) { keepOne(value, c) },
	)
}
