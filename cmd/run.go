package cmd

import (
	"github.com/devopsgig/restapi/src/server"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the REST API",
	Run:   run,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

// Starts the server
func run(cmd *cobra.Command, args []string) {
	server.Run()
}
