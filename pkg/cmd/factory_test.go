package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
	"strings"
	"testing"
)

func TestBuild(t *testing.T) {
	cases := []struct{
		Name string
		CommandId int
		Streams genericclioptions.IOStreams
		ExpectError bool
		Validate func(t *testing.T, cmd Command)
	}{
		{
			Name: "Should error an error for invalid command id",
			CommandId: 100,
			Streams: genericclioptions.IOStreams{
				In: os.Stdin, Out: os.Stdout,
				ErrOut: os.Stderr,
			},
			ExpectError: true,
			Validate: func(t *testing.T, cmd Command) {
				if cmd != nil {
					t.Fatal("Expected cmd to return a nil pointer")
				}
			},
		},
		{
			Name: "Should the version sub-command",
			CommandId: 0,
			Streams: genericclioptions.IOStreams{
				In: os.Stdin, Out: os.Stdout,
				ErrOut: os.Stderr,
			},
			ExpectError: false,
			Validate: func(t *testing.T, cmd Command) {
				if cmd.GetUse() != "version" {
					t.Fatalf("Expected cmd sub-command to be \"version\" but got \"%s\"", cmd.GetUse())
				}
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			cmd, err := Build(tc.CommandId, tc.Streams)

			if tc.ExpectError && err == nil {
				t.Fatal("Expected error but got none")
			}

			if !tc.ExpectError && err != nil {
				t.Fatalf("Got unexpected error: %v", err)
			}

			tc.Validate(t, cmd)
		})
	}
}

func TestBuildCommand(t *testing.T) {
	cases := []struct{
		Name string
		GetCmd func() Command
		Validate func(t *testing.T, cmd *cobra.Command)
	}{
		{
			Name: "Should validate command",
			GetCmd: func() Command {
				streams := genericclioptions.IOStreams{
					In: os.Stdin, Out: os.Stdout,
					ErrOut: os.Stderr,
				}
				cmd, _ := Build(VersionCmd, streams)

				return cmd
			},
			Validate: func(t *testing.T, cmd *cobra.Command) {
				example := cmd.Example
				expectedExample := `
version
	# get the current aerogear cli/plugin version
	kubectl ag version
	oc plugin ag version`
				if strings.Compare(expectedExample, example) != 0 {
					t.Fatalf("Expected \"%s\" but got \"%s\"", expectedExample, example)
				}
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			cmd := tc.GetCmd()
			cobraCmd := BuildCommand(cmd)
			tc.Validate(t, cobraCmd)
		})
	}
}
