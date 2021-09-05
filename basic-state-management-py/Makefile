SHELL = /bin/bash

SERVICES = \
	microservices/api/gateway \
	microservices/svc/user \

.PHONY: all
all: pre proto fmt lint

.PHONY: install
install:
	make -C .devcontainer install


.PHONY: pre
pre:
	@poetry install --no-root
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


.PHONY: build
build:
	skaffold build


.PHONY: kind
kind:
	kind get clusters -q | grep "basic-service-invocation-py" || kind create cluster --config kind.yaml

.PHONY: clean
clean:
	kind delete cluster --name basic-service-invocation-py

.PHONY: dev
dev:
	skaffold dev


.PHONY: deploy-production
deploy-production:
	docker login ghcr.io
	skaffold run -p production

.PHONY: destroy-production
destroy-production:
	skaffold delete -p production

.PHONY: http
http:
	curl -i localhost:58080/user/v1/users -H "Content-Type: application/json" -d '{"name": "alice"}'

.PHONY: grpc
grpc:
	grpcurl -protoset <(buf build -o -) -plaintext -d '{"name": "alice"}' localhost:55050 basic.user.v1.UserService/Create || true
