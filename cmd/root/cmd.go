package root

import (
	"aerogear.org/aerogear-cli/cmd/app"
	"aerogear.org/aerogear-cli/cmd/link"
	"aerogear.org/aerogear-cli/cmd/push"
	"aerogear.org/aerogear-cli/cmd/services"
	"aerogear.org/aerogear-cli/cmd/version"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ag",
	Short: "ag is a kubernetes plugin",
	Long:  `kubectl plugin to manage mobile apps in a kubernetes cluster`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(
		app.AppCmd,
		link.LinkCmd,
		push.PushCmd,
		services.ServicesCmd,
		version.VersionCmd,
	)
}
