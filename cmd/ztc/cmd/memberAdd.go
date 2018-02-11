package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var memberAddCmd = &cobra.Command{
	Use:   "add node-id",
	Short: "Add a node to the member list",
	Long:  `Add a node to the member list.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("memberAdd called")
	},
}

func init() {
	memberCmd.AddCommand(memberAddCmd)
}
