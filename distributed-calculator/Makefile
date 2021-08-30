SHELL = /bin/bash

SERVICES = \
	microservices/api/gateway \
	microservices/svc/greeter \
	microservices/svc/adder \
	microservices/svc/subtractor \
	microservices/svc/multiplier \
	microservices/svc/divider \

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
	kind get clusters -q | grep "distributed-calculator" || kind create cluster --config kind.yaml

.PHONY: clean
clean:
	kind delete cluster --name distributed-calculator

.PHONY: dev
dev:
	skaffold run -f skaffold.dapr.yaml
	skaffold dev


.PHONY: http
http:
	curl -i localhost:58080/adder/add -H "Content-Type: application/json" -d '{"operand_one": "0.1234", "operand_two": "5.6789"}'
	curl -i localhost:58080/subtractor/sub -H "Content-Type: application/json" -d '{"operand_one": "0.1234", "operand_two": "5.6789"}'
	curl -i localhost:58080/multiplier/mul -H "Content-Type: application/json" -d '{"operand_one": "0.1234", "operand_two": "5.6789"}'
	curl -i localhost:58080/divider/div -H "Content-Type: application/json" -d '{"operand_one": "0.1234", "operand_two": "5.6789"}'

.PHONY: grpc
grpc:
	grpcurl -protoset <(buf build -o -) -plaintext -d '{"operand_one": "0.1234", "operand_two": "5.6789"}' localhost:55050 distributedcalculator.adder.v1.Adder/Add || true
	grpcurl -protoset <(buf build -o -) -plaintext -d '{"operand_one": "0.1234", "operand_two": "5.6789"}' localhost:55051 distributedcalculator.subtractor.v1.Subtractor/Sub || true
	grpcurl -protoset <(buf build -o -) -plaintext -d '{"operand_one": "0.1234", "operand_two": "5.6789"}' localhost:55052 distributedcalculator.multiplier.v1.Multiplier/Mul || true
	grpcurl -protoset <(buf build -o -) -plaintext -d '{"operand_one": "0.1234", "operand_two": "5.6789"}' localhost:55053 distributedcalculator.divider.v1.Divider/Div || true
