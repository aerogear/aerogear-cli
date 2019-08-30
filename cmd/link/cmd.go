package link

import (
	"aerogear.org/aerogear-cli/cmd/link/link_auth"
	"github.com/spf13/cobra"
)

var clientTypeVar string

func init() {
	LinkCmd.AddCommand(link_auth.LinkAuthCmd)
}

var LinkCmd = &cobra.Command{
	Use:   "link",
	Short: "link",
	Long:  `link a service`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
