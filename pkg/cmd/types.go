package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

const (
	VersionCmd = 0
)

type Command interface {
	GetUse() string
	GetDescription() string
	GetExample() string
	AddFlags(flags *pflag.FlagSet)
	Complete(cmd *cobra.Command, args []string) error
	Validate() error
	Bootstrap(streams genericclioptions.IOStreams)
	Run() error
}
