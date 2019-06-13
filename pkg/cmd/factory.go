package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func BuildCommand(vc Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:          vc.GetUse(),
		Short:        vc.GetDescription(),
		Example:      vc.GetExample(),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := vc.Complete(c, args); err != nil {
				return err
			}

			if err := vc.Validate(); err != nil {
				return err
			}

			if err := vc.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	vc.AddFlags(cmd.Flags())

	return cmd
}

func Build(cmdId int, streams genericclioptions.IOStreams) (Command, error) {
	switch cmdId {
	case VersionCmd:
		cmd := &versionCmd{}
		cmd.Bootstrap(streams)
		return cmd, nil
	default:
		return nil, errors.New("Invalid command")
	}
}
