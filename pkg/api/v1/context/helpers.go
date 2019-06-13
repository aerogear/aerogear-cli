package context

import (
	"fmt"
	"k8s.io/client-go/tools/clientcmd/api"
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
