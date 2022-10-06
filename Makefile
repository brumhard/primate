SHELL=/bin/bash -e -o pipefail
PWD = $(shell pwd)

# constants
GOLANGCI_VERSION = 1.43.0
DOCKER_REPO = ghcr.io/brumhard/primate
DOCKER_TAG = $(shell git rev-parse --short HEAD)

all: git-hooks generate tidy ## Initializes all tools

out:
	@mkdir -p out

git-hooks:
	@git config --local core.hooksPath .githooks/

download: ## Downloads the dependencies
	@go mod download

tidy: ## Cleans up go.mod and go.sum
	@go mod tidy

fmt: ## Formats all code with go fmt
	@go fmt ./...

run: fmt ## Run the app
	@go run ./cmd/primate/main.go

test-build: ## Tests whether the code compiles
	@go build -o /dev/null ./...

build: app/build/web out/bin ## Builds all binaries

.PHONY: app/build/web
app/build/web:
	cd app && flutter build web

GO_BUILD = mkdir -pv "$(@)" && go build -ldflags="-w -s" -o "$(@)" ./...
.PHONY: out/bin
out/bin:
	$(GO_BUILD)

GOLANGCI_LINT = bin/golangci-lint-$(GOLANGCI_VERSION)
$(GOLANGCI_LINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- -b bin v$(GOLANGCI_VERSION)
	@mv bin/golangci-lint "$(@)"

lint: fmt $(GOLANGCI_LINT) download ## Lints all code with golangci-lint
	@$(GOLANGCI_LINT) run

lint-reports: out/lint.xml

.PHONY: out/lint.xml
out/lint.xml: $(GOLANGCI_LINT) out download
	@$(GOLANGCI_LINT) run ./... --out-format checkstyle | tee "$(@)"

test: ## Runs all tests
	@go test $(ARGS) ./...

coverage: out/report.json ## Displays coverage per func on cli
	go tool cover -func=out/cover.out

html-coverage: out/report.json ## Displays the coverage results in the browser
	go tool cover -html=out/cover.out

test-reports: out/report.json

.PHONY: out/report.json
out/report.json: out
	@go test -count 1 ./... -coverprofile=out/cover.out --json | tee "$(@)"

clean: ## Cleans up everything
	@rm -rf bin out protodeps

docker: ## Builds docker image
	docker buildx build -t $(DOCKER_REPO):$(DOCKER_TAG) .
# Go dependencies versioned through tools.go
GO_DEPENDENCIES = google.golang.org/protobuf/cmd/protoc-gen-go \
				google.golang.org/grpc/cmd/protoc-gen-go-grpc \
				github.com/envoyproxy/protoc-gen-validate \
				github.com/bufbuild/buf/cmd/buf \
                github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking \
                github.com/bufbuild/buf/cmd/protoc-gen-buf-lint

define make-go-dependency
  # target template for go tools, can be referenced e.g. via /bin/<tool>
  bin/$(notdir $1):
	GOBIN=$(PWD)/bin go install $1
endef

.PHONY: api/proto/buf.lock
api/proto/buf.lock: bin/buf
	@bin/buf mod update api/proto

# this creates a target for each go dependency to be referenced in other targets
$(foreach dep, $(GO_DEPENDENCIES), $(eval $(call make-go-dependency, $(dep))))

protolint: api/proto/buf.lock bin/protoc-gen-buf-lint ## Lints your protobuf files
	bin/buf lint

protobreaking: api/proto/buf.lock bin/protoc-gen-buf-breaking ## Compares your current protobuf with the version on master to find breaking changes
	bin/buf breaking --against '.git#branch=main'

generate: ## Generates code from protobuf files
generate: api/proto/buf.lock bin/protoc-gen-go bin/protoc-gen-go-grpc bin/protoc-gen-validate
	PATH=$(PWD)/bin:$$PATH buf generate

ci: lint-reports test-reports ## Executes lint and test and generates reports

help: ## Shows the help
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
        awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ''
