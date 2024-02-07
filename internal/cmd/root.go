package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/omissis/goturin-todod/internal/app"
	cobrax "github.com/omissis/goturin-todod/internal/x/cobra"
)

type RootCommand struct {
	*cobra.Command
	cfg app.Config
}

func NewRootCommand(cfg app.Config, versions map[string]string) *RootCommand {
	const envPrefix = ""

	root := &RootCommand{
		Command: &cobra.Command{
			PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
				cobrax.BindFlags(cmd, cobrax.InitEnvs(envPrefix), log.Fatal, envPrefix)

				return nil
			},
			Use:           "todod",
			SilenceUsage:  true,
			SilenceErrors: true,
		},
		cfg: cfg,
	}

	cobrax.BindFlags(root.Command, cobrax.InitEnvs(envPrefix), log.Fatal, envPrefix)

	root.AddCommand(NewVersionCommand(versions))
	root.AddCommand(NewServeCommand(app.NewContainer()))

	return root
}
