package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
	"testing"
)

func TestVersionCmd_GetUse(t *testing.T) {
	cmd := &versionCmd{}
	assert.Equal(t, "version", cmd.GetUse())
}

func TestVersionCmd_GetDescription(t *testing.T) {
	cmd := &versionCmd{}
	assert.Equal(t, "Displays the current aerogear cli/plugin version", cmd.GetDescription())
}

func TestVersionCmd_GetExample(t *testing.T) {
	cmd := &versionCmd{}
	current := `
version
	# get the current aerogear cli/plugin version
	kubectl ag version
	oc plugin ag version`
	assert.Equal(t, current, cmd.GetExample())
}

func TestVersionCmd_Bootstrap(t *testing.T) {
	cmd := &versionCmd{}
	assert.Nil(t, cmd.ctx)

	streams := genericclioptions.IOStreams{
		In: os.Stdin, Out: os.Stdout,
		ErrOut: os.Stderr,
	}
	cmd.Bootstrap(streams)
	assert.NotNil(t, cmd.ctx)
	assert.Equal(t, streams, cmd.ctx.IOStreams)
}

func TestVersionCmd_Validate(t *testing.T) {
	var err error
	streams := genericclioptions.IOStreams{
		In: os.Stdin, Out: os.Stdout,
		ErrOut: os.Stderr,
	}
	cmd := &versionCmd{}
	cmd.Bootstrap(streams)

	err = cmd.ctx.Validate()
	assert.NotNil(t, err)

	err = cmd.Complete(BuildCommand(cmd), make([]string, 0))
	assert.Nil(t, err)

	err = cmd.ctx.Validate()
	assert.Nil(t, err)
}

func TestVersionCmd_Run(t *testing.T) {
	var err error
	var buf bytes.Buffer
	streams := genericclioptions.IOStreams{
		In: os.Stdin,
		Out: &buf,
		ErrOut: os.Stderr,
	}

	cmd := &versionCmd{}
	cmd.Bootstrap(streams)

	err = cmd.Complete(BuildCommand(cmd), make([]string, 0))
	assert.Nil(t, err)

	err = cmd.Run()
	assert.Nil(t, err)
	assert.Equal(t, "aerogear-mobile-cli v0.0.1\n", buf.String())
}
