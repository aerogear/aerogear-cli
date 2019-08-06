package helpers

import (
	"k8s.io/apimachinery/pkg/runtime"
	"time"
)

type TimeOpts struct {
	Interval time.Duration
	Timeout  time.Duration
}

type ReadinessFn func(ro runtime.Object) (bool, error)
