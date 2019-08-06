package context

import (
	"fmt"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"os/user"
	"path/filepath"
	"strings"
)

func IsContextEqual(ctxA, ctxB *api.Context) bool {
	if ctxA == nil || ctxB == nil {
		return false
	}

	if ctxA.Cluster != ctxB.Cluster {
		return false
	}

	if ctxA.Namespace != ctxB.Namespace {
		return false
	}

	if ctxA.AuthInfo != ctxB.AuthInfo {
		return false
	}

	return true
}

func GenerateContextName(fromContext *api.Context) string {
	name := fromContext.Namespace
	if len(fromContext.Cluster) > 0 {
		name = fmt.Sprintf("%s/%s", name, fromContext.Cluster)
	}

	if len(fromContext.AuthInfo) > 0 {
		cleanAuthInfo := strings.Split(fromContext.AuthInfo, "/")[0]
		name = fmt.Sprintf("%s/%s", name, cleanAuthInfo)
	}

	return name
}

func GetUser() (*user.User, error) {
	return user.Current()
}

func GetRestConfig(usr *user.User) (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", filepath.Join(usr.HomeDir, ".kube", "config"))
}
