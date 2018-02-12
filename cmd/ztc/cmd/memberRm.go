package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

var memberRmCmd = &cobra.Command{
	Use:   "rm node-id",
	Short: "Remove a node from the member list",
	Long:  `Remove a node from the member list.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("this command accepts exactly one argument")
		}

		// TODO: Validate node ID here
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cntrl := getAPIController()
		networkID, err := cmd.Flags().GetString("net")

		if err != nil {
			panic(err)
		}

		nodeID := args[0]

		err = cntrl.RemoveMember(networkID, nodeID)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	memberCmd.AddCommand(memberRmCmd)
}
