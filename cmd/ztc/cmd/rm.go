package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm network-id",
	Short: "Stop managing a network",
	Long:  `Stop managing a network.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("this command accepts exactly one argument")
		}

		// TODO: Validate network ID here
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cntrl := getAPIController()
		networkID := args[0]

		err := cntrl.RemoveNetwork(networkID)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
