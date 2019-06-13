SHELL = /bin/bash
BIN_DIR = ${GOPATH}/bin
BIN_NAME = kubectl-ag
PLUGIN_FOLDER = ${HOME}/.kube/plugins/ag
.DEFAULT_GOAL = all

.PHONY: clean
clean:
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

.PHONY: test
test: test/unit

.PHONY: build/cli
build/cli:
	@go build -o bin/${BIN_NAME} ./cmd/kubectl-ag.go

.PHONY: build
build: clean build/cli

.PHONY: install
install: build
	@cp bin/kubectl-ag ${BIN_DIR}/${BIN_NAME}
	@mkdir -p ${PLUGIN_FOLDER}
	@cp extras/plugin.yaml ${PLUGIN_FOLDER}/plugin.yaml

.PHONY: all
all: clean code/fix test install