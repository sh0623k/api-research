PATH_OF_REST_SCHEMA = schemas/rest
PATH_OF_REST_GENERATED_CODE = generated/rest

.PHONY: grpc-go-code
grpc-go-code:
	protoc \
 		./schemas/grpc/todo/v1/todo.proto \
		--go_out=./generated/grpc/todo/v1

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
generated-code: grpc-go-code graphql-go-code oapi-go-code
