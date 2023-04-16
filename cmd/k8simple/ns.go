package k8simple

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nsCmd = &cobra.Command{
	Use:     "ns",
	Aliases: []string{"n", "namespace"},
	Short:   "Switch to a namespace",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		status := k8simple.switchNamespace(args[0])
		if status {
			fmt.Println("Switched to namespace " + args[0])
		} else {
			fmt.Println("Unable to switch into namespace... :(")
		}
	},
}
