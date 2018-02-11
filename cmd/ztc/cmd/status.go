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
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zaninime/ztc/api"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display the controller status",
	Long:  "Display the controller status.",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := url.Parse(viper.GetString("baseUrl"))

		if err != nil {
			log.Fatal(err)
		}

		cntrl := api.Controller{
			BaseURL:   *url,
			AuthToken: viper.GetString("authToken"),
		}

		status, err := cntrl.GetStatus()
		if err != nil {
			log.Fatal(err)
		}

		time := time.Unix(0, int64(status.Clock)*1000000)

		fmt.Printf("API Version: %d\nCurrent Time: %s\n", status.APIVersion, time)
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
