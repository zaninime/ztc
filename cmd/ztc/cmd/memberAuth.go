package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

var memberAuthCmd = &cobra.Command{
	Use:   "auth node-id",
	Short: "Authorize or revoke access to a node",
	Long:  `Authorize or revoke access to a node.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("this command accepts exactly one arguments")
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

		revoke, err := cmd.Flags().GetBool("revoke")

		if err != nil {
			panic(err)
		}

		nodeID := args[0]

		member, err := cntrl.GetMember(networkID, nodeID)

		if err != nil {
			log.Fatal(err)
		}

		member.Authorized = !revoke

		_, err = cntrl.EditMember(networkID, nodeID, member.EditableMember)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	memberCmd.AddCommand(memberAuthCmd)
	memberAuthCmd.Flags().BoolP("revoke", "r", false, "Revoke authorization")
}
