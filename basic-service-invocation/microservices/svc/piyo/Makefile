.PHONY: all
all: pre fmt lint

.PHONY: pre
pre:
	go mod tidy

.PHONY: fmt
fmt:
	goimports -l -local "github.com/kzmake/examples/basic-service-invocation/microservices/svc/piyo" -w .

.PHONY: lint
lint:
	golangci-lint run
