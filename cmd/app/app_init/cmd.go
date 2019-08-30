package app_init

import (
	"aerogear.org/aerogear-cli/pkg/ag/resources/mobileclient"
	"fmt"
	"github.com/lucasjones/reggen"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var appName string

var AppInitCmd = &cobra.Command{
	Use:   "init",
	Short: "app init",
	Long:  `app init`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("App name is required")
		}

		appName = args[0]

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := os.Mkdir(".ag", 0775)
		if err != nil {
			if !os.IsExist(err) {
				return err
			}

		}
		fmt.Println("==> .ag folder created.")

		apiKey, err := reggen.Generate("^(\\w{8}-\\w{4}-\\w{4}-\\w{4}-\\w{11})$", 11)
		if err != nil {
			return err
		}

		tmplData := mobileclient.MobileClientData{
			Name:   appName,
			ApiKey: apiKey,
		}
		data, err := mobileclient.NewMobileClient(tmplData)
		if err != nil {
			return err
		}

		_, err = os.Stat(".ag/mobileclient.yaml")
		if err != nil {
			if os.IsPermission(err) || os.IsTimeout(err) {
				return err
			}
		}

		if err == nil {
			fmt.Println("==> .ag/mobileclient.yaml already exists")
			return nil
		}

		err = ioutil.WriteFile(".ag/mobileclient.yaml", data, 0775)
		if err != nil {
			return err
		}

		fmt.Println("==> .ag/mobileclient.yaml created.")

		return nil
	},
}
