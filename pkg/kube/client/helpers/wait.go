package helpers

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

func DefaultTimeOpts() TimeOpts {
	return TimeOpts{
		Interval: time.Second * 5,
		Timeout:  time.Minute * 10,
	}
}

func WaitForDeletion(instance runtime.Object, dynclient client.Client, opts *TimeOpts) error {
	key, err := client.ObjectKeyFromObject(instance)
	if err != nil {
		return err
	}

	ctx, cleanup := context.WithTimeout(context.Background(), opts.Timeout)
	defer cleanup()

	return wait.Poll(opts.Interval, opts.Timeout, func() (done bool, err error) {
		innerErr := dynclient.Get(ctx, key, instance)
		if errors.IsNotFound(innerErr) {
			return true, nil
		}

		if innerErr != nil {
			return true, innerErr
		}

		return false, nil
	})
}

func WaitForReadiness(instance runtime.Object, dynclient client.Client, opts *TimeOpts, fn ReadinessFn) error {
	key, err := client.ObjectKeyFromObject(instance)
	if err != nil {
		return err
	}

	ctx, cleanup := context.WithTimeout(context.Background(), opts.Timeout)
	defer cleanup()

	return wait.Poll(opts.Interval, opts.Timeout, func() (done bool, err error) {
		innerErr := dynclient.Get(ctx, key, instance)
		if innerErr != nil {
			return true, err
		}

		err = dynclient.Get(context.TODO(), key, instance)
		if err != nil {
			return true, err
		}

		return fn(instance)
	})
}
