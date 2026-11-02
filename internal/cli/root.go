package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Cobras rootCmd
var rootCmd = &cobra.Command{
	Use:   "swix",
	Short: "Cybersecurity Swiss Army knife.",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("a subcommand is required")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

// Pre-run func
func init() {

}
