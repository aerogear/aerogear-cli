package main

import (
	"aerogear.org/aerogear-cli/cmd/root"
	"aerogear.org/aerogear-cli/pkg/kube/context"
	"os"
)

func main() {
	ctx := context.DefaultContext()
	if err := root.RootCmd.Execute(); err != nil {
		if _, panicErr := ctx.ErrOut.Write([]byte(err.Error())); panicErr != nil {
			panic(panicErr)
		}
		os.Exit(1)
	}
}
