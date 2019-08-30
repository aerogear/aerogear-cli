package link_auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client2 "sigs.k8s.io/controller-runtime/pkg/client"
)

func getAuthService(c client2.Client) (*v1beta1.ServiceInstance, error) {
	list := &v1beta1.ServiceInstanceList{}
	opt := &client2.ListOptions{}
	err := c.List(context.TODO(), opt, list)
	if err != nil {
		return nil, err
	}

	for _, instance := range list.Items {
		if instance.Spec.ClusterServiceClassExternalName == "ag-keycloak-identity-management-apb" {
			return &instance, nil
		}
	}

	return nil, errors.New("auth service not found")
}

func buildBindParams(instance *v1beta1.ServiceInstance, appName string, clientType string) *v1.Secret {
	return &v1.Secret{
		TypeMeta: v12.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: v12.ObjectMeta{
			Name: fmt.Sprintf("%s-bind-parameters-ppn2f", instance.Name),
		},
		Type: v1.SecretTypeOpaque,
		StringData: map[string]string{
			"parameters": fmt.Sprintf(`{"CLIENT_ID":"%s","CLIENT_TYPE":"%s"}`, appName, clientType),
		},
	}
}

func buildServiceBinding(instance *v1beta1.ServiceInstance, secret *v1.Secret, appName string) *v1beta1.ServiceBinding {
	return &v1beta1.ServiceBinding{
		TypeMeta: v12.TypeMeta{
			APIVersion: "servicecatalog.k8s.io/v1beta1",
			Kind:       "ServiceBinding",
		},
		ObjectMeta: v12.ObjectMeta{
			Annotations: map[string]string{
				"binding.aerogear.org/consumer": appName,
				"binding.aerogear.org/provider": instance.Name,
			},
			Labels: map[string]string{
				"mdc.aerogear.org/clientId": appName,
			},
			Name: fmt.Sprintf("%s-identitymanagement-p9f9w", instance.Name),
		},
		Spec: v1beta1.ServiceBindingSpec{
			InstanceRef: v1beta1.LocalObjectReference{
				Name: instance.Name,
			},
			ParametersFrom: []v1beta1.ParametersFromSource{
				{
					SecretKeyRef: &v1beta1.SecretKeyReference{
						Name: secret.Name,
						Key:  "parameters",
					},
				},
			},
		},
	}
}
