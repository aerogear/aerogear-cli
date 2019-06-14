module aerogear.org/aerogear-cli

go 1.12

require (
	github.com/pkg/errors v0.8.0
	github.com/spf13/cobra v0.0.0-20180319062004-c439c4fa0937
	github.com/spf13/pflag v1.0.1
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.3.0
	k8s.io/cli-runtime v0.0.0-20190531135611-d60f41fb4dc3
	k8s.io/client-go v0.0.0-20190531132438-d58e65e5f4b1
)

replace (
	golang.org/x/sync => golang.org/x/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190209173611-3b5209105503
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190313210603-aa82965741a9
	k8s.io/api => k8s.io/api v0.0.0-20190531132109-d3f5f50bdd94
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190531131812-859a0ba5e71a
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190531135611-d60f41fb4dc3
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190531132438-d58e65e5f4b1
)

replace k8s.io/component-base => k8s.io/component-base v0.0.0-20190531133342-103ccccb7a11
