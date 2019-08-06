package mobileclient

import (
	"aerogear.org/aerogear-cli/pkg/apis/aerogear/v1alpha1"
	"github.com/magiconair/properties/assert"
	"sigs.k8s.io/yaml"
	"testing"
)

func TestNewMobileClient(t *testing.T) {
	cases := []struct {
		Name        string
		Data        *MobileClientData
		ExpectError bool
		Validate    func(t *testing.T, data []byte)
	}{
		{
			Name: "Should parse mobile client",
			Data: &MobileClientData{
				Name:   "myapp",
				ApiKey: "ce485ad0-6b75-44aa-86a4-8e74355a5b4c",
			},
			ExpectError: false,
			Validate: func(t *testing.T, data []byte) {
				client := &v1alpha1.MobileClient{}

				err := yaml.Unmarshal(data, client)
				if err != nil {
					t.Fatalf("Failed to parse mobile client: %v", err)
				}

				assert.Equal(t, client.Spec.Name, "myapp")
				assert.Equal(t, client.Name, "myapp")
				assert.Equal(t, client.Spec.ApiKey, "ce485ad0-6b75-44aa-86a4-8e74355a5b4c")
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			data, err := NewMobileClient(*tc.Data)

			if tc.ExpectError && err == nil {
				t.Fatal("Expected error but got none")
			}

			if !tc.ExpectError && err != nil {
				t.Fatal(err)
			}

			tc.Validate(t, data)
		})
	}
}
