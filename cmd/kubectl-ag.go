package main

import (
	cmd2 "aerogear.org/aerogear-cli/pkg/cmd"
	"fmt"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
)

func showHelp(cliOpts *genericclioptions.IOStreams) {
	cids := []int{
		cmd2.VersionCmd,
	}
	for _, cid := range cids {
		cmd, err := cmd2.Build(cid, *cliOpts)
		if err != nil {
			if _, bufError := cliOpts.ErrOut.Write([]byte(fmt.Sprint(err))); bufError != nil {
				panic(bufError)
			}
			os.Exit(1)
		}

		cmd.Bootstrap(*cliOpts)
		subcmd := cmd2.BuildCommand(cmd)
		_, err = cliOpts.Out.Write([]byte(fmt.Sprintf("%s\n", subcmd.Example)))
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	cliOpts := genericclioptions.IOStreams{
		In: os.Stdin, Out: os.Stdout,
		ErrOut: os.Stderr,
	}

	if len(os.Args) == 1 {
		showHelp(&cliOpts)
		os.Exit(1)
	}

	var cmd cmd2.Command
	var err error

	switch os.Args[1] {
	case "version":
		cmd, err = cmd2.Build(cmd2.VersionCmd, cliOpts)
		if err != nil {
			if _, bufError := cliOpts.ErrOut.Write([]byte(fmt.Sprint(err))); bufError != nil {
				panic(bufError)
			}
			os.Exit(1)
		}
	default:
		showHelp(&cliOpts)
		os.Exit(1)
	}

	cmd.Bootstrap(cliOpts)
	subcmd := cmd2.BuildCommand(cmd)
	if err := subcmd.Execute(); err != nil {
		if _, bufError := cliOpts.ErrOut.Write([]byte(fmt.Sprint(err))); bufError != nil {
			panic(bufError)
		}
		os.Exit(1)
	}

	os.Exit(0)
}
