package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var memberShowCmd = &cobra.Command{
	Use:   "show [node-id]",
	Short: "List members or show details about one",
	Long:  `List members or show details about one.`,
	Run: func(cmd *cobra.Command, args []string) {
		cntrl := getAPIController()
		networkID, err := cmd.Flags().GetString("net")

		if err != nil {
			log.Fatal(err)
		}

		members, err := cntrl.GetMemberList(networkID)

		if err != nil {
			log.Fatal(err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
		fmt.Fprintln(w, "ID\tAuthorized\tBridge\tSince")

		for _, nodeID := range members {
			node, err := cntrl.GetMember(networkID, nodeID)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(w, "%s\t%t\t%t\t%s\n", node.ID, node.Authorized, node.ActiveBridge, node.CreationTime)
		}

		w.Flush()
	},
}

func init() {
	memberCmd.AddCommand(memberShowCmd)
}
