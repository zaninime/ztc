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
	"fmt"
	"log"
	"net/url"

	yaml "github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zaninime/ztc/api"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a network interactively",
	Long: `Edit a network interactively. The network configuration will
be presented as a YAML file in your preferred editor. Save the file and quit to commit.
To abort, save an empty file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("Pass the network id")
		}

		networkID := args[0]

		url, err := url.Parse(viper.GetString("baseUrl"))

		if err != nil {
			log.Fatal(err)
		}

		cntrl := api.Controller{
			BaseURL:   *url,
			AuthToken: viper.GetString("authToken"),
		}

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
