package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run unit tests",
	Run:   test,
}

func init() {
	RootCmd.AddCommand(testCmd)
}

// Run unit tests with test coverage
func test(cmd *cobra.Command, args []string) {
	command := exec.Command("go", "test", "-v", "-cover", "./src/...")
	output, err := command.Output()
	fmt.Println(string(output))
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}
