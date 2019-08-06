package version

import (
	"aerogear.org/aerogear-cli/version"
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ag",
	Long:  `All software has versions. This is ag.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("AeroGear CLI %s", version.Version)
	},
}
