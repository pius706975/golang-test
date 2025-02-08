package cmd

import (
	serve "github.com/pius706975/golang-test/api"
	"github.com/pius706975/golang-test/package/database"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "backend",
	Long: "Go backend service",
}

func init() {
	initCommand.AddCommand(serve.ServeCMD)
	initCommand.AddCommand(database.MigrationCMD)
	initCommand.AddCommand(database.SeedCMD)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}