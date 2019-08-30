package app_gen

import (
	"aerogear.org/aerogear-cli/pkg/ag/app"
	"aerogear.org/aerogear-cli/pkg/ag/resources/manifest"
	"aerogear.org/aerogear-cli/pkg/apis/aerogear/v1alpha1"
	"aerogear.org/aerogear-cli/pkg/kube/client"
	context2 "aerogear.org/aerogear-cli/pkg/kube/context"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/types"
)

var genCtx *context2.CtxOptions

func init() {
	genCtx = context2.DefaultContext()
	genCtx.AddFlags(AppGenCmd.Flags())
}

var AppGenCmd = &cobra.Command{
	Use:   "gen",
	Short: "app gen",
	Long:  `app gen`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := genCtx.SetupContext(cmd, args)
		if err != nil {
			return err
		}

		appName, err := app.GetAppname()
		if err != nil {
			return err
		}

		c, err := client.GetClient()
		if err != nil {
			return err
		}

		mobServ := &manifest.MobServ{}

		mobClient := &v1alpha1.MobileClient{}
		key := types.NamespacedName{
			Name:      appName,
			Namespace: genCtx.GetNamespace(),
		}
		err = c.Get(context.TODO(), key, mobClient)
		if err != nil {
			return err
		}

		mobServ.Version = mobClient.Generation
		mobServ.Services = mobClient.Status.Services
		mobServ.ClientId = mobClient.Name
		mobServ.Namespace = mobClient.Namespace

		raw, err := json.Marshal(mobServ)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, raw, "", "\t")
		if err != nil {
			return err
		}

		err = ioutil.WriteFile("mobile-services.json", prettyJSON.Bytes(), 0755)
		if err != nil {
			return err
		}
		fmt.Println("==> mobile-services.json generated")

		return nil
	},
}
