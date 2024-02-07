package main

//go:generate go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0
//go:generate oapi-codegen --config=oapi-codegen.yaml api/openapi.json
//go:generate go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.25.0
//go:generate sqlc generate

import (
	"log"

	"github.com/omissis/goturin-todod/internal/app"
	"github.com/omissis/goturin-todod/internal/cmd"
)

var (
	version   = "unknown"
	gitCommit = "unknown"
	buildTime = "unknown"
	goVersion = "unknown"
	osArch    = "unknown"
)

func main() {
	versions := map[string]string{
		"version":   version,
		"gitCommit": gitCommit,
		"buildTime": buildTime,
		"goVersion": goVersion,
		"osArch":    osArch,
	}

	if err := cmd.NewRootCommand(app.Config{}, versions).Execute(); err != nil {
		log.Fatal(err)
	}
}
