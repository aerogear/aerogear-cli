SHELL = /bin/bash
BIN_DIR = ${GOPATH}/bin
BIN_NAME = ag
PLUGIN_FOLDER = ${HOME}/.kube/plugins/ag
.DEFAULT_GOAL = all

.PHONY: clean
clean:
	@go clean -testcache || true
	@rm -rf bin

.PHONYY: gen/plugin
gen/plugin:
	hack/gen/plugin.sh

.PHONY: code/check
code/check:
	@diff -u <(echo -n) <(gofmt -d `find . -type f -name '*.go' -not -path "./vendor/*"`)

.PHONY: code/fix
code/fix:
	@gofmt -w `find . -type f -name '*.go' -not -path "./vendor/*"`

.PHONY: test/unit
test/unit:
	@go test -v -race -cover ./pkg/...

.PHONY: test/e2e
test/e2e:
	@kubectl apply -f deploy/crds --validate=false
	@go test -v -race -cover ./test/e2e/...

.PHONY: test
test: test/unit test/e2e

.PHONY: build/cli
build/cli:
	@go build -o bin/${BIN_NAME} ./cmd/main.go

.PHONY: build
build: clean build/cli

.PHONY: install
install: build
	@cp bin/ag ${BIN_DIR}/kubectl-ag
	@cp bin/ag ${BIN_DIR}/${BIN_NAME}
	@mkdir -p ${PLUGIN_FOLDER}
	@cp extras/plugin.yaml ${PLUGIN_FOLDER}/plugin.yaml

.PHONY: prepare/crds
prepare/crds:
	@oc apply -f ./deploy/crds/

.PHONY: all
all: clean code/fix test install
