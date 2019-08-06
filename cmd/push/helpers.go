package push

import (
	"aerogear.org/aerogear-cli/pkg/apis/aerogear/v1alpha1"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

func getAppDef() (*v1alpha1.MobileClient, error) {
	path := ".ag/mobileclient.yaml"

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	appClient := &v1alpha1.MobileClient{}
	err = yaml.Unmarshal(data, appClient)
	if err != nil {
		return nil, err
	}

	return appClient, nil
}

func getAuthSecret() (*v1.Secret, error) {
	path := ".ag/auth-secret.yaml"

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	instance := &v1.Secret{}
	err = yaml.Unmarshal(data, instance)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func getAuthBinding() (*v1beta1.ServiceBinding, error) {
	path := ".ag/auth-service-binding.yaml"

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	instance := &v1beta1.ServiceBinding{}
	err = yaml.Unmarshal(data, instance)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
