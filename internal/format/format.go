// Package format provides plain-text labels for Lighthouse instruction trees.
package format

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	solana "github.com/fluxrpc/solana-go"
)

func Program(name string, id solana.PublicKey) string {
	return fmt.Sprintf("Program: %s %s", name, id)
}

func Instruction(name string) string { return "Instruction: " + name }

func Param(name string, value interface{}) string {
	valueText := strings.TrimSpace(spew.Sdump(value))
	return name + ": " + strings.ReplaceAll(valueText, "\n", "\n"+strings.Repeat(" ", len(name)+2))
}

func Meta(name string, meta *solana.AccountMeta) string {
	if meta == nil {
		return name + ": <nil>"
	}
	var roles []string
	if meta.IsWritable {
		roles = append(roles, "WRITE")
	}
	if meta.IsSigner {
		roles = append(roles, "SIGN")
	}
	return fmt.Sprintf("%s: %s [%s]", name, meta.PublicKey, strings.Join(roles, ", "))
}
