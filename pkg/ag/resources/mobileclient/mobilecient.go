package mobileclient

import (
	"aerogear.org/aerogear-cli/pkg/ag/helpers"
	"bytes"
	"html/template"
)

func NewMobileClient(data MobileClientData) ([]byte, error) {
	rawData, err := helpers.GetRawTemplate("mobileclient.yaml.tmpl")
	if err != nil {
		return nil, err
	}

	tmpl := template.New("mobileclient-template")
	parsed, err := tmpl.Parse(string(rawData[:]))
	if err != nil {
		return nil, err
	}

	buff := &bytes.Buffer{}
	err = parsed.Execute(buff, data)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
