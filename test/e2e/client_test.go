package e2e

import (
	"aerogear.org/aerogear-cli/pkg/apis/aerogear/v1alpha1"
	client2 "aerogear.org/aerogear-cli/pkg/kube/client"
	"aerogear.org/aerogear-cli/pkg/kube/client/helpers"
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"testing"

	"context"
)

func TestClient(t *testing.T) {
	client, err := client2.GetClient()
	assert.NotNil(t, client)
	assert.Nil(t, err)

	labels := map[string]string{"app": "memcached"}
	replicas := new(int32)
	*replicas = 1

	dep := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "memcached",
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "memcached:1.4.36-alpine",
						Name:    "memcached",
						Command: []string{"memcached", "-m=64", "-o", "modern", "-v"},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 11211,
							Name:          "memcached",
						}},
					}},
				},
			},
		},
	}

	err = client.Create(context.TODO(), dep)
	assert.Nil(t, err)

	timeOpts := helpers.DefaultTimeOpts()
	emptyDep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "memcached",
			Namespace: "default",
		},
	}
	err = helpers.WaitForReadiness(emptyDep, client, &timeOpts, func(ro runtime.Object) (b bool, e error) {
		if emptyDep.Status.AvailableReplicas == 1 {
			return true, nil
		}

		return false, nil
	})

	err = client.Delete(context.TODO(), dep)
	assert.Nil(t, err)

	err = helpers.WaitForDeletion(dep, client, &timeOpts)
	assert.Nil(t, err)
}

func TestMobileClientCreation(t *testing.T) {
	client, err := client2.GetClient()
	assert.NotNil(t, client)
	assert.Nil(t, err)

	clientApp := &v1alpha1.MobileClient{}

	clientApp.Namespace = "default"
	clientApp.Name = "myapp"
	clientApp.Spec.Name = "myapp"
	clientApp.Spec.ApiKey = "11111111-1111-1111-1111-11111111111"

	err = client.Create(context.TODO(), clientApp)
	assert.Nil(t, err)

	timeOpts := helpers.DefaultTimeOpts()
	emptyMobileClient := &v1alpha1.MobileClient{}
	err = helpers.WaitForReadiness(emptyMobileClient, client, &timeOpts, func(ro runtime.Object) (b bool, e error) {
		return true, nil
	})

	err = client.Delete(context.TODO(), clientApp)
	assert.Nil(t, err)

	err = helpers.WaitForDeletion(clientApp, client, &timeOpts)
	assert.Nil(t, err)

}
