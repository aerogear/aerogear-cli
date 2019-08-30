package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	MobileClientType = metav1.TypeMeta{
		APIVersion: "mdc.aerogear.org/v1alpha1",
		Kind:       "MobileClient",
	}
)

type MobileClientSpec struct {
	ApiKey string `json:"apiKey"`
	DmzUrl string `json:"dmzUrl"`
	Name   string `json:"name"`
}

type MobileClientStatus struct {
	Services []MobileClientService `json:"services"`
}

type MobileClientService struct {
	Id     string                 `json:"id"`
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Url    string                 `json:"url"`
	Config map[string]interface{} `json:"config"`
}

// +k8s:openapi-gen=true
type MobileClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MobileClientSpec   `json:"spec,omitempty"`
	Status MobileClientStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MobileClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MobileClient `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MobileClient{}, &MobileClientList{})
}
