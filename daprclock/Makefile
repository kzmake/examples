SHELL = /bin/bash

GOTOOLS = \
	github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
	github.com/ramya-rao-a/go-outline@latest \
	github.com/cweill/gotests/gotests@latest \
	github.com/fatih/gomodifytags@latest \
	github.com/josharian/impl@latest \
	github.com/haya14busa/goplay/cmd/goplay@latest \
	github.com/go-delve/delve/cmd/dlv@latest \
	honnef.co/go/tools/cmd/staticcheck@latest \
	golang.org/x/tools/gopls@latest \
	golang.org/x/tools/cmd/goimports@latest \
	github.com/golangci/golangci-lint/cmd/golangci-lint@latest \

PBTOOLS = \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
	google.golang.org/protobuf/cmd/protoc-gen-go@latest \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
	github.com/bufbuild/buf/cmd/buf@latest \
	github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@latest \
	github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@latest \
	github.com/fullstorydev/grpcurl/cmd/grpcurl@latest \

SERVICES = \
	microservices/api/gateway \
	microservices/svc/clock \
	microservices/svc/hourhand \
	microservices/svc/secondhand \
	microservices/wrk/tick \
	microservices/wrk/secondhand \

.PHONY: all
all: proto pre fmt lint

.PHONY: install
install: install-base install-golang install-proto

.PHONY: install-base
install-base:
	@for t in $(GOTOOLS); do go install $$t; done

.PHONY: install-golang
install-golang:
	@for t in $(GOTOOLS); do go install $$t; done

.PHONY: install-proto
install-proto:
	@for t in $(PBTOOLS); do go install $$t; done


.PHONY: pre
pre:
	@for f in $(SERVICES); do make -C $$f pre; done

.PHONY: fmt
fmt:
	@for f in $(SERVICES); do make -C $$f fmt; done

.PHONY: lint
lint:
	@for f in $(SERVICES); do make -C $$f lint; done


.PHONY: proto
proto:
	buf mod update
	buf generate


.PHONY: kind
kind:
	kind get clusters -q | grep "daprclock" || kind create cluster --config kind.yaml

.PHONY: clean
clean:
	kind delete cluster --name daprclock

.PHONY: dev
dev:
	skaffold run -f skaffold.components.yaml
	skaffold run -f skaffold.dapr.yaml
	skaffold dev


.PHONY: http
http:
	curl -i localhost:58080/clock/now

.PHONY: grpc
grpc:
	grpcurl -protoset <(buf build -o -) -plaintext localhost:55050 daprclock.clock.v1.Clock/Now || true
	grpcurl -protoset <(buf build -o -) -plaintext localhost:55051 daprclock.hourhand.v1.HourHand/Now || true
	grpcurl -protoset <(buf build -o -) -plaintext localhost:55052 daprclock.minutehand.v1.MinuteHand/Now || true
	grpcurl -protoset <(buf build -o -) -plaintext localhost:55053 daprclock.secondhand.v1.SecondHand/Now || true
