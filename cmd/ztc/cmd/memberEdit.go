package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var memberEditCmd = &cobra.Command{
	Use:   "edit node-id",
	Short: "Edit the membership of a node",
	Long:  `Edit the membership of a node.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("memberEdit called")
	},
}

func init() {
	memberCmd.AddCommand(memberEditCmd)
}
