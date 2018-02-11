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

// addCmd represents the add command
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

		network, err := cntrl.AddNetwork(networkID)

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
