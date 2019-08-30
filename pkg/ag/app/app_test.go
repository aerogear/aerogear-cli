package app

import "testing"

func TestGetAppname(t *testing.T) {
	cases := []struct {
		Name        string
		Folder      string
		ExpectError bool
		Validate    func(t *testing.T, name string)
	}{
		{
			Name:        "Should retrieve app name",
			Folder:      "_testdata",
			ExpectError: false,
			Validate: func(t *testing.T, name string) {
				if name != "myapp" {
					t.Fatalf("Expected name myapp but got %s", name)
				}
			},
		},
		{
			Name:        "Should not retrieve app name",
			Folder:      "_notfound",
			ExpectError: true,
			Validate: func(t *testing.T, name string) {
				if name != "" {
					t.Fatal("app name should be empty")
				}
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			name, err := GetAppname(tc.Folder)
			t.Log(err)

			if tc.ExpectError && err == nil {
				t.Fatal("Expected error but got none")
			}

			if !tc.ExpectError && err != nil {
				t.Fatal(err)
			}

			tc.Validate(t, name)
		})
	}
}
