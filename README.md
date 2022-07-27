# api-research

## Purpose

Research GraphQL, gRPC, and REST.

## Development

### Requirements

- Go
  - 1.18 or later
- Python
  - 3.7 or later
- pip
  - 9.0.1 or later
- Locust
  - 2.10.1 or later

Download modules by executing the below command.

```shell
go mod download
```

Refer to the [Quick start](https://grpc.io/docs/languages/python/quickstart/#grpc) to install grpcio.

### IDE

- IntelliJ IDEA

### Development Environment Construction

Generate code by the following command after cloning this repository.
 
```shell
make generated-code
```

## Load testing

Prepare for load testing by the following commands.

```shell
make start-containers
make create-test-data
```

Perform GraphQL load testing.

```shell
make run-locust-to-graphql
```

Perform gRPC load testing.

```shell
make run-locust-to-grpc
```

Perform REST load testing.

```shell
make run-locust-to-rest
```
