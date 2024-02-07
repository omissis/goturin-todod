package cobra_test

import (
	"testing"

	"github.com/omissis/goturin-todod/internal/x/cobra"
)

func TestInitEnvs(t *testing.T) {
	v := cobra.InitEnvs("test")

	if v == nil {
		t.Error("InitEnvs() returned nil")
	}
}
