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

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show [network-id]",
	Short: "List managed networks or show the details about one",
	Long: `Show a list of all the networks managed by the controller or
the details about a specific network.`,
	Args: func(cmd *cobra.Command, args []string) error {
		editableOnly, err := cmd.Flags().GetBool("editable-only")
		if err != nil {
			panic(err)
		}

		if editableOnly && len(args) == 0 {
			return errors.New("--editable-only can be used only when specifying the network id")
		}

		if len(args) > 1 {
			return errors.New("this command accepts at most one argument")
		}

		// TODO: Validate network ID here

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cntrl := getAPIController()

		if len(args) != 1 {
			networks, err := cntrl.GetNetworkList()

			if err != nil {
				log.Fatal(err)
			}

			for _, networkID := range networks {
				network, err := cntrl.GetNetwork(networkID)
				if err != nil {
					log.Fatal(err)
				}

				networkType := "public"

				if network.Private {
					networkType = "private"
				}

				fmt.Printf("%s: %s, %s\n", network.ID, network.EditableNetwork.Name, networkType)
			}

			return
		}

		networkID := args[0]

		network, err := cntrl.GetNetwork(networkID)

		if err != nil {
			log.Fatal(err)
		}

		editableOnly, err := cmd.Flags().GetBool("editable-only")
		if err != nil {
			panic(err)
		}

		var encodedNetwork []byte

		if editableOnly {
			encodedNetwork, err = yaml.Marshal(network.EditableNetwork)
		} else {
			encodedNetwork, err = yaml.Marshal(network)
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(encodedNetwork))
	},
}

func init() {
	RootCmd.AddCommand(showCmd)

	showCmd.Flags().Bool("editable-only", false, "Only return the editable part of the network")
}
