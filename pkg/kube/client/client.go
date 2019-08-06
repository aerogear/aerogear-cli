package client

import (
	"aerogear.org/aerogear-cli/pkg/apis"
	"aerogear.org/aerogear-cli/pkg/kube/context"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _client client.Client

func GetClient() (client.Client, error) {
	var err error

	if _client == nil {
		err = bootstrap()
	}

	return _client, err
}

func bootstrap() error {
	user, err := context.GetUser()
	if err != nil {
		return err
	}

	cfg, err := context.GetRestConfig(user)
	if err != nil {
		return err
	}

	opts, err := getOptions()
	if err != nil {
		return err
	}

	_client, err = client.New(cfg, opts)
	if err != nil {
		return err
	}

	return nil
}

func getOptions() (client.Options, error) {
	opts := client.Options{
		Scheme: runtime.NewScheme(),
	}

	err := apis.AddToScheme(opts.Scheme)
	if err != nil {
		return opts, err
	}

	return opts, nil
}
