package helpers

import "testing"

func TestGetRawTemplate(t *testing.T) {
	cases := []struct {
		Name         string
		TemplateName string
		ExpectError  bool
	}{
		{
			Name:         "Should load template",
			TemplateName: "mobileclient.yaml.tmpl",
			ExpectError:  false,
		},
		{
			Name:         "Should not load template (not found)",
			TemplateName: "404.yaml.tmpl",
			ExpectError:  true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			_, err := GetRawTemplate(tc.TemplateName)

			if tc.ExpectError && err == nil {
				t.Fatal("Expected error but got none")
			}

			if !tc.ExpectError && err != nil {
				t.Fatal(err)
			}
		})
	}
}
