package app

import (
	app_gen2 "aerogear.org/aerogear-cli/cmd/app/app_gen"
	app_init2 "aerogear.org/aerogear-cli/cmd/app/app_init"
	"github.com/spf13/cobra"
)

func init() {
	AppCmd.AddCommand(
		app_init2.AppInitCmd,
		app_gen2.AppGenCmd,
	)
}

var AppCmd = &cobra.Command{
	Use:   "app",
	Short: "App",
	Long:  `Mobile Aplication`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
