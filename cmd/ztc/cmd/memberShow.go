package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var memberShowCmd = &cobra.Command{
	Use:   "show [node-id]",
	Short: "List members or show details about one",
	Long:  `List members or show details about one.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("memberShow called")
	},
}

func init() {
	memberCmd.AddCommand(memberShowCmd)
}
