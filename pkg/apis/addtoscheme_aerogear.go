package apis

import (
	"aerogear.org/aerogear-cli/pkg/apis/aerogear/v1alpha1"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func init() {
	AddToSchemes = append(
		AddToSchemes,
		v1.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1.AddToScheme,
		corev1.AddToScheme)
}
