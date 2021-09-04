.PHONY: all
all: pre fmt lint

.PHONY: pre
pre:
	go mod tidy

.PHONY: fmt
fmt:
	goimports -l -local "github.com/kzmake/examples/basic-state-management/microservices/api/gateway" -w .

.PHONY: lint
lint:
	golangci-lint run
