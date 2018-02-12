package cmd

import (
	"errors"
	"fmt"
	"log"

	yaml "github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var memberAddCmd = &cobra.Command{
	Use:   "add node-id",
	Short: "Add a node to the member list",
	Long:  `Add a node to the member list.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("this command accepts exactly one argument")
		}

		// TODO: Validate node ID here
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		networkID, err := cmd.Flags().GetString("net")

		if err != nil {
			panic(err)
		}

		cntrl := getAPIController()
		nodeID := args[0]

		member, err := cntrl.AddMember(networkID, nodeID, nil)

		if err != nil {
			log.Fatal(err)
		}

		encodedMember, err := yaml.Marshal(member)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(encodedMember))
	},
}

func init() {
	memberCmd.AddCommand(memberAddCmd)
}
