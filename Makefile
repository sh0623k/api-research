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
