# grpc-web-example

This repository demonstrates the usage of grpc-web with vuejs and a Go microservice.

## Directories:
* `api` - Contains the gRPC/proto definitions
* `envoy` - Contains the envoy configuration for grpc-web to grpc proxying
* `time` - time microservice in Go, listens on port 9090 for gRPC connections
* `frontend-ts` - simple vuejs 3 application using connect-es plugins to communicate with the backend using gRPC-Web

## Usage:
* `make proto` - Generate proto clients
* `make run-frontend` - Start frontend
* `make run-servers` - Start Time services and Envoy proxy
