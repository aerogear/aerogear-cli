package cmd

import (
	v1ctx "aerogear.org/aerogear-cli/pkg/api/v1/context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type versionCmd struct {
	ctx *v1ctx.CtxOptions
}

func (vc *versionCmd) GetUse() string {
	return "version"
}

func (vc *versionCmd) GetDescription() string {
	return "Displays the current aerogear cli/plugin version"
}

func (vc *versionCmd) GetExample() string {
	return fmt.Sprintf(`
%[1]s
	# get the current aerogear cli/plugin version
	kubectl ag version
	oc plugin ag version`, vc.GetUse())
}

func (vc *versionCmd) Bootstrap(streams genericclioptions.IOStreams) {
	vc.ctx = v1ctx.NewCtxOptions(streams)
}

func (vc *versionCmd) Complete(cmd *cobra.Command, args []string) error {
	return vc.ctx.SetupContext(cmd, args)
}

func (vc *versionCmd) Validate() error {
	return vc.ctx.Validate()
}

func (vc *versionCmd) AddFlags(flags *pflag.FlagSet) {
	vc.ctx.AddFlags(flags)
}

func (vc *versionCmd) Run() error {
	version := vc.getVersion()
	output := fmt.Sprintf("%s\n", version)

	if _, err := vc.ctx.Out.Write([]byte(output)); err != nil {
		return err
	}

	return nil
}

func (vc *versionCmd) getVersion() string {
	return "aerogear-mobile-cli v0.0.1"
}
