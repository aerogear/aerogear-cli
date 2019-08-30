module aerogear.org/aerogear-cli

go 1.12

require (
	github.com/gobuffalo/packr v1.30.1
	github.com/google/pprof v0.0.0-20190723021845-34ac40c74b70 // indirect
	github.com/ianlancetaylor/demangle v0.0.0-20181102032728-5e5cf60278f6 // indirect
	github.com/kubernetes-incubator/service-catalog v0.2.1
	github.com/lucasjones/reggen v0.0.0-20180717132126-cdb49ff09d77
	github.com/magiconair/properties v1.8.0
	github.com/pkg/errors v0.8.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.3.0
	golang.org/x/arch v0.0.0-20190815191158-8a70ba74b3a1 // indirect
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190531132109-d3f5f50bdd94
	k8s.io/apimachinery v0.0.0-20190531131812-859a0ba5e71a
	k8s.io/cli-runtime v0.0.0-20190531135611-d60f41fb4dc3
	k8s.io/client-go v0.0.0-20190531132438-d58e65e5f4b1
	sigs.k8s.io/controller-runtime v0.1.11
	sigs.k8s.io/yaml v1.1.0
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
