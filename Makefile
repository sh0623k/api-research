PATH_OF_GRPC_SCHEMA = ./schemas/grpc/todo/v1
PATH_OF_GRPC_GENERATED_GO_CODE = ./generated/grpc
PATH_OF_GRPC_GENERATED_PYTHON_CODE = ./generated/grpc/todo/v1
PATH_OF_REST_SCHEMA = ./schemas/rest
PATH_OF_REST_GENERATED_CODE = ./generated/rest

.PHONY: grpc-code
grpc-code:
	rm -rf ${PATH_OF_GRPC_GENERATED_GO_CODE}
	mkdir ${PATH_OF_GRPC_GENERATED_GO_CODE}
	protoc \
 		${PATH_OF_GRPC_SCHEMA}/todo.proto \
		--go_out=${PATH_OF_GRPC_GENERATED_GO_CODE} \
		--go-grpc_out=${PATH_OF_GRPC_GENERATED_GO_CODE}
	python3 -m grpc_tools.protoc \
		-I=${PATH_OF_GRPC_SCHEMA} \
		--python_out=${PATH_OF_GRPC_GENERATED_PYTHON_CODE} \
		--grpc_python_out=${PATH_OF_GRPC_GENERATED_PYTHON_CODE} \
		todo.proto

.PHONY: server-graphql-go-code
server-graphql-go-code:
	go run github.com/99designs/gqlgen generate

.PHONY: client-graphql-go-code
client-graphql-go-code:
	go run github.com/Khan/genqlient

.PHONY: graphql-go-code
graphql-go-code: server-graphql-go-code client-graphql-go-code

.PHONY: oapi-go-code
oapi-go-code:
	rm -rf ${PATH_OF_REST_GENERATED_CODE}
	mkdir ${PATH_OF_REST_GENERATED_CODE}
	mkdir ${PATH_OF_REST_GENERATED_CODE}/types
	mkdir ${PATH_OF_REST_GENERATED_CODE}/server
	oapi-codegen --config ${PATH_OF_REST_SCHEMA}/types.cfg.yaml --package types ${PATH_OF_REST_SCHEMA}/todo.yaml
	oapi-codegen --config ${PATH_OF_REST_SCHEMA}/server.cfg.yaml --package server ${PATH_OF_REST_SCHEMA}/todo.yaml

.PHONY: generated-code
generated-code: grpc-code graphql-go-code oapi-go-code

.PHONY: start-containers
start-containers:
	docker compose down --volumes --remove-orphans
	docker compose up -d --build

.PHONY: create-test-data
create-test-data:
	go run locust/cmd/main.go

.PHONY: run-locust-to-graphql
run-locust-to-graphql:
	cd locust/graphql; locust -H http://localhost:8080

.PHONY: run-locust-to-grpc
run-locust-to-grpc:
	cd locust/grpc; locust -H http://localhost:50051

.PHONY: run-locust-to-rest
run-locust-to-rest:
	cd locust/rest; locust -H http://localhost:3000
