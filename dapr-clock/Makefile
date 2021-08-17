SHELL = /bin/bash

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
install:
	make -C .devcontainer install

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
	kind get clusters -q | grep "dapr-clock" || kind create cluster --config kind.yaml

.PHONY: clean
clean:
	kind delete cluster --name dapr-clock

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
