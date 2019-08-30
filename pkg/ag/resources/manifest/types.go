package manifest

import "aerogear.org/aerogear-cli/pkg/apis/aerogear/v1alpha1"

type Component struct {
	Host    string `json:"host"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Mobile  bool   `json:"mobile"`
}

type Manifest struct {
	Components []Component `json:"components"`
}

type MobServ struct {
	Version   int64                          `json:"version"`
	Namespace string                         `json:"namespace"`
	ClientId  string                         `json:"clientId"`
	Services  []v1alpha1.MobileClientService `json:"services"`
}
