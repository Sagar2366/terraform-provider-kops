PROVIDER_VERSION := "0.0.1"
OS := $(shell echo `uname` | tr '[:upper:]' '[:lower:]')

.PHONY: all
all: clean gen fmt build vet test

.PHONY: clean
clean:
	@rm -f terraform-provider-kops
	@rm -rf ./pkg/schemas/config
	@rm -rf ./pkg/schemas/datasources
	@rm -rf ./pkg/schemas/kops
	@rm -rf ./pkg/schemas/kube
	@rm -rf ./pkg/schemas/resources
	@rm -rf ./pkg/schemas/utils
	@rm -rf ./docs/data-sources/*.md
	@rm -rf ./docs/provider-config/*.md
	@rm -rf ./docs/resources/*.md

.PHONY: install-goimports
install-goimports:
ifeq (, $(shell which goimports))
	go install golang.org/x/tools/cmd/goimports@latest
endif

.PHONY: gen-tf-code
gen-tf-code: clean install-goimports
	@go run ./hack/gen-tf-code/...
	@go fmt ./pkg/schemas/...
	@goimports -w ./pkg/schemas

.PHONY: gen
gen: gen-tf-code

.PHONY: build
build: gen
	@CGO_ENABLED=0 go build -ldflags="-s -w -X 'github.com/eddycharly/terraform-provider-kops/pkg/version.BuildVersion=v${PROVIDER_VERSION}'" ./cmd/terraform-provider-kops

.PHONY: fmt
fmt: build
	@go fmt ./cmd/...
	@go fmt ./pkg/...

.PHONY: test
test: fmt
	@go test ./...

.PHONY: vet
vet: fmt
	@go vet ./...

.PHONY: install
install: all
	@mkdir -p ${HOME}/.terraform.d/plugins/github/eddycharly/kops/${PROVIDER_VERSION}/${OS}_amd64
	@cp terraform-provider-kops $(HOME)/.terraform.d/plugins/github/eddycharly/kops/${PROVIDER_VERSION}/${OS}_amd64/terraform-provider-kops

# EXAMPLES FOR TERRAFORM >= 0.15

.PHONY: examples
examples: example-basic example-aws-profile example-aws-assume-role example-bastion example-klog

.PHONY: example-basic
example-basic: install
	@terraform -chdir=./examples/basic init
	@terraform -chdir=./examples/basic validate
	@terraform -chdir=./examples/basic plan

.PHONY: example-aws-profile
example-aws-profile: install
	@terraform -chdir=./examples/aws-profile init
	@terraform -chdir=./examples/aws-profile validate
	@terraform -chdir=./examples/aws-profile plan

.PHONY: example-aws-assume-role
example-aws-assume-role: install
	@terraform -chdir=./examples/aws-assume-role init
	@terraform -chdir=./examples/aws-assume-role validate

.PHONY: example-bastion
example-bastion: install
	@terraform -chdir=./examples/bastion init
	@terraform -chdir=./examples/bastion validate
	@terraform -chdir=./examples/bastion plan

.PHONY: example-klog
example-klog: install
	@terraform -chdir=./examples/klog init
	@terraform -chdir=./examples/klog validate
	@terraform -chdir=./examples/klog plan

# INTEGRATION TESTS

.PHONY: integration
integration: integration-basic integration-external-policies

.PHONY: integration-reset
integration-reset:
	@rm -rf ./store
	@rm -f 	./terraform.tfstate

.PHONY: integration-basic
integration-basic: integration-reset
	@terraform -chdir=./tests/basic init
	@terraform -chdir=./tests/basic validate
	@terraform -chdir=./tests/basic plan
	@terraform -chdir=./tests/basic apply -auto-approve

.PHONY: integration-external-policies
integration-external-policies: integration-reset
	@terraform -chdir=./tests/external-policies init
	@terraform -chdir=./tests/external-policies validate
	@terraform -chdir=./tests/external-policies plan
	@terraform -chdir=./tests/external-policies apply -auto-approve
