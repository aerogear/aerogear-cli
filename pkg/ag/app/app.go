package app

import (
	"aerogear.org/aerogear-cli/pkg/apis/aerogear/v1alpha1"
	"fmt"
	"io/ioutil"
	"sigs.k8s.io/yaml"
)

func GetAppname(folder ...string) (string, error) {
	var agFolder string
	if len(folder) > 0 {
		agFolder = folder[0]
	} else {
		agFolder = ".ag"
	}

	path := fmt.Sprintf("%s/mobileclient.yaml", agFolder)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	appClient := &v1alpha1.MobileClient{}
	err = yaml.Unmarshal(data, appClient)
	if err != nil {
		return "", err
	}

	return appClient.Name, nil
}
