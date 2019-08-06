package push

import (
	"aerogear.org/aerogear-cli/pkg/kube/client"
	context2 "aerogear.org/aerogear-cli/pkg/kube/context"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
)

var kctx *context2.CtxOptions

func init() {
	kctx = context2.DefaultContext()
	kctx.AddFlags(PushCmd.Flags())

}

var PushCmd = &cobra.Command{
	Use:   "push",
	Short: "push changer to server",
	Long:  `push changes to server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := kctx.SetupContext(cmd, args)
		if err != nil {
			return err
		}

		c, err := client.GetClient()
		if err != nil {
			return err
		}

		//mobile client
		app, err := getAppDef()
		if err != nil {
			return err
		}
		app.Namespace = kctx.GetNamespace()
		err = c.Create(context.TODO(), app)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				return err
			}
			fmt.Println("==> [SKIP] mobile client already exists")
		} else {
			fmt.Println("==> mobile client created")
		}

		//auth secret
		secret, err := getAuthSecret()
		if err != nil {
			return err
		}
		secret.Namespace = kctx.GetNamespace()
		err = c.Create(context.TODO(), secret)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				return err
			}
			fmt.Println("==> [SKIP] auth secret already exists")
		} else {
			fmt.Println("==> auth secret created")
		}

		//auth service binding
		binding, err := getAuthBinding()
		if err != nil {
			return err
		}
		binding.Namespace = kctx.GetNamespace()
		err = c.Create(context.TODO(), binding)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				return err
			}
			fmt.Println("==> [SKIP] auth service binding already exists")
		} else {
			fmt.Println("==> auth service binding created")
		}

		return nil
	},
}
