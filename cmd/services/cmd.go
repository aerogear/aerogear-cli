package services

import (
	manifest2 "aerogear.org/aerogear-cli/pkg/ag/resources/manifest"
	"aerogear.org/aerogear-cli/pkg/kube/client"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"os"
	"text/tabwriter"
)

var ServicesCmd = &cobra.Command{
	Use:   "services",
	Short: "services",
	Long:  `list available services`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.GetClient()
		if err != nil {
			return err
		}

		secret := &v1.Secret{}
		key := types.NamespacedName{
			Name:      "manifest",
			Namespace: "webapp",
		}
		err = c.Get(context.TODO(), key, secret)
		if err != nil {
			return err
		}

		manifestBytes := secret.Data["generated_manifest"]

		manifest := manifest2.Manifest{}
		err = json.Unmarshal(manifestBytes, manifest)
		if err != nil {
			return err
		}

		if len(manifest.Components) == 0 {
			return errors.New("No resources found.")
		}

		w := tabwriter.NewWriter(os.Stdout, 5, 0, 10, ' ', tabwriter.TabIndent)
		_, err = fmt.Fprintln(w, "NAME\tHOST\tVERSION\t")
		if err != nil {
			return err
		}
		for _, component := range manifest.Components {
			if component.Mobile {
				continue
			}
			_, err = fmt.Fprintf(w, "%s\t%s\t%s\n", component.Name, component.Host, component.Version)
			if err != nil {
				return err
			}
		}
		err = w.Flush()
		if err != nil {
			return err
		}

		return nil
	},
}
