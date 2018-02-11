package cmd

import (
	"fmt"
	"log"

	yaml "github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display the controller status",
	Long:  "Display the controller status.",
	Run: func(cmd *cobra.Command, args []string) {
		cntrl := getAPIController()

		status, err := cntrl.GetStatus()
		if err != nil {
			log.Fatal(err)
		}

		yamlStatus, err := yaml.Marshal(status)

		fmt.Println(string(yamlStatus))
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
