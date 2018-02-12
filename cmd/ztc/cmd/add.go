package cmd

import (
	"errors"
	"fmt"
	"log"

	yaml "github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [network-id]",
	Short: "Add a network to manage",
	Long:  `Add a network to manage.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("this command accepts at most one argument")
		}

		// TODO: Validate network ID here
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		networkID := ""

		if len(args) == 1 {
			networkID = args[0]
		}

		cntrl := getAPIController()

		network, err := cntrl.AddNetwork(networkID, nil)

		if err != nil {
			log.Fatal(err)
		}

		encodedNetwork, err := yaml.Marshal(network)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(encodedNetwork))
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
