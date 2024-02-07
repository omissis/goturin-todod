package cmd

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"

	"github.com/omissis/goturin-todod/internal/app"
)

const (
	apiServerPort = 8080
	databasePort  = 5432
)

func NewServeCommand(ctr *app.Container) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run todod server",
		Args:  cobra.ExactArgs(0),
		RunE: func(_ *cobra.Command, _ []string) error {
			fmt.Println("Running the Todod server...")

			return ctr.APIServer().Run()
		},
	}

	setupServeCmdFlags(cmd, ctr)

	return cmd
}

func setupServeCmdFlags(cmd *cobra.Command, ctr *app.Container) {
	cmd.Flags().StringVar(&ctr.APIServerHost, "api-host", "0.0.0.0", "server host")
	cmd.Flags().Uint16Var(&ctr.APIServerPort, "api-port", apiServerPort, "server port")
	cmd.Flags().StringSliceVar(&ctr.APIAllowedOrigins, "api-allowed-origins", nil, "server port")

	cmd.Flags().StringVar(&ctr.DBName, "db-name", "todod", "database name")
	cmd.Flags().StringVar(&ctr.DBHost, "db-host", "0.0.0.0", "database host")
	cmd.Flags().StringVar(&ctr.DBPassword, "db-password", "todod", "database password")
	cmd.Flags().Uint16Var(&ctr.DBPort, "db-port", databasePort, "database port")
	cmd.Flags().StringVar(
		&ctr.DBSslMode,
		"db-ssl-mode",
		"disable",
		"databse ssl mode (accepted values: disable, require, verify-ca, verify-full)",
	)
	cmd.Flags().StringVar(&ctr.DBUser, "db-user", "todod", "database user")
}
