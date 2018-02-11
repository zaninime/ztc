package cmd

import (
	"errors"
	"io/ioutil"
	"log"

	yaml "github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"github.com/zaninime/ztc/api"
)

var editCmd = &cobra.Command{
	Use:   "edit network-id config-file.yml",
	Short: "Edit a network",
	Long: `Edit a network. Provide both the network ID and a yaml file
to read the network configuration from.

Tip: dump the current configuration first using
  ztc show network-id --editable-only > config-file.yml`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("this command accepts exactly two arguments")
		}

		// TODO: validate network id
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cntrl := getAPIController()

		networkID := args[0]
		configFile := args[1]

		fileContent, err := ioutil.ReadFile(configFile)

		if err != nil {
			log.Fatal(err)
		}

		var config api.EditableNetwork
		err = yaml.Unmarshal(fileContent, &config)

		if err != nil {
			log.Fatal(err)
		}

		_, err = cntrl.EditNetwork(networkID, &config)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(editCmd)
}
