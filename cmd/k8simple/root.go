package k8simple

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k8simple",
	Short: "k8simple - a simple cli makes k8s cli options more easy",
	Long: `k8simple - Cli tool helps you to easily perform k8s operations
One can use k8simple to switch modify clusters, deploys and pods etc`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "^^^Whoops, There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
