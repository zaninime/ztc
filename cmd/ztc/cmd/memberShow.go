package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	yaml "github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var memberShowCmd = &cobra.Command{
	Use:   "show [node-id]",
	Short: "List members or show details about one",
	Long:  `List members or show details about one.`,
	Args: func(cmd *cobra.Command, args []string) error {
		editableOnly, err := cmd.Flags().GetBool("editable-only")
		if err != nil {
			panic(err)
		}

		if editableOnly && len(args) == 0 {
			return errors.New("--editable-only can be used only when specifying the node id")
		}

		if len(args) > 1 {
			return errors.New("this command accepts at most one argument")
		}

		// TODO: Validate node ID here

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cntrl := getAPIController()
		networkID, err := cmd.Flags().GetString("net")

		if err != nil {
			panic(err)
		}

		if len(args) == 0 {
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
			return
		}

		nodeID := args[0]
		editableOnly, err := cmd.Flags().GetBool("editable-only")

		if err != nil {
			panic(err)
		}

		node, err := cntrl.GetMember(networkID, nodeID)

		if err != nil {
			log.Fatal(err)
		}

		var encodedMember []byte

		if editableOnly {
			encodedMember, err = yaml.Marshal(node.EditableMember)
		} else {
			encodedMember, err = yaml.Marshal(node)
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(encodedMember))
	},
}

func init() {
	memberCmd.AddCommand(memberShowCmd)

	memberShowCmd.Flags().Bool("editable-only", false, "Only return the editable part of the network")
}
