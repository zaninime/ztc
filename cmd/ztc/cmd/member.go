package cmd

import (
	"github.com/spf13/cobra"
)

// memberCmd represents the member command
var memberCmd = &cobra.Command{
	Use:   "member",
	Short: "Manage members of a network",
	Long:  `Manage members of a network.`,
}

func init() {
	RootCmd.AddCommand(memberCmd)

	memberCmd.PersistentFlags().String("net", "", "Network ID used for the operations")
	memberCmd.MarkPersistentFlagRequired("net")
}
