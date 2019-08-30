package context

import (
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
)

var _ctx *CtxOptions

func DefaultContext() *CtxOptions {
	if _ctx == nil {

		_ctx = NewCtxOptions(genericclioptions.IOStreams{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		})
	}

	return _ctx
}
