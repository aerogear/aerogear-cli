package helpers

import "github.com/gobuffalo/packr"

func GetRawTemplate(name string) ([]byte, error) {
	box := packr.NewBox("../../../res")
	rawData, err := box.Find(name)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}
