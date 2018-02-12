package cmd

import (
	"errors"
	"io/ioutil"
	"log"

	yaml "github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"github.com/zaninime/ztc/api"
)

var memberEditCmd = &cobra.Command{
	Use:   "edit node-id",
	Short: "Edit the membership of a node",
	Long:  `Edit the membership of a node.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("this command accepts exactly two arguments")
		}

		// TODO: validate node id
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cntrl := getAPIController()
		networkID, err := cmd.Flags().GetString("net")

		if err != nil {
			panic(err)
		}

		nodeID := args[0]
		configFile := args[1]

		fileContent, err := ioutil.ReadFile(configFile)

		if err != nil {
			log.Fatal(err)
		}

		var config api.EditableMember
		err = yaml.Unmarshal(fileContent, &config)

		if err != nil {
			log.Fatal(err)
		}

		_, err = cntrl.EditMember(networkID, nodeID, &config)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	memberCmd.AddCommand(memberEditCmd)
}
