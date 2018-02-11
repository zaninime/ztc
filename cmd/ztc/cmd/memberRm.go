package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var memberRmCmd = &cobra.Command{
	Use:   "rm node-id",
	Short: "Remove a node from the member list",
	Long:  `Remove a node from the member list.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("memberRm called")
	},
}

func init() {
	memberCmd.AddCommand(memberRmCmd)
}
