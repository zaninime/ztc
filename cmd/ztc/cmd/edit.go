// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"log"

	yaml "github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
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
		network, err := cntrl.GetNetwork(networkID)

		if err != nil {
			log.Fatal(err)
		}

		network.EditableNetwork.Name = "hello world"

		network2, err := cntrl.EditNetwork(networkID, network.EditableNetwork)

		if err != nil {
			log.Fatal(err)
		}

		yaml, err := yaml.Marshal(network2)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(yaml))
	},
}

func init() {
	RootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
