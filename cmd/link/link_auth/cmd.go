package link_auth

import (
	"aerogear.org/aerogear-cli/pkg/ag/app"
	"aerogear.org/aerogear-cli/pkg/kube/client"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"sigs.k8s.io/yaml"
)

var clientTypeVar string

func init() {
	LinkAuthCmd.Flags().StringVarP(&clientTypeVar, "client-type", "c", "public", "auth client type")
}

var LinkAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "auth",
	Long:  `authentication service linking`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := os.Stat(".ag/mobileclient.yaml")
		if os.IsNotExist(err) {
			return errors.New("application not initialized")
		}

		appName, err := app.GetAppname()
		if err != nil {
			return err
		}

		c, err := client.GetClient()
		if err != nil {
			return err
		}

		authService, err := getAuthService(c)
		if err != nil {
			return err
		}

		bindParams := buildBindParams(authService, appName, clientTypeVar)
		raw, err := yaml.Marshal(bindParams)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(".ag/auth-secret.yaml", raw, 0775)
		if err != nil {
			return err
		}
		fmt.Println("==> .ag/auth-secret.yaml.")

		serviceBinidng := buildServiceBinding(authService, bindParams, appName)
		raw, err = yaml.Marshal(serviceBinidng)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(".ag/auth-service-binding.yaml", raw, 0775)
		if err != nil {
			return err
		}
		fmt.Println("==> .ag/auth-service-binding.yaml")

		return nil
	},
}
