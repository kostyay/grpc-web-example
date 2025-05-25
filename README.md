# grpc-web-example

This repository demonstrates the usage of grpc-web with vuejs and a Go microservice.

```mermaid
flowchart LR
    Browser[Browser]:::blue-->|HTTP REQUEST|VueJS["VueJS Application"]:::green
    VueJS-->|grpc-web|Envoy[Envoy]:::orange
    Envoy-->|grpc|TimeService["time service (go)"]:::purple

    classDef blue fill:#2196F3,stroke:#000,stroke-width:2px,color:#fff
    classDef green fill:#4CAF50,stroke:#000,stroke-width:2px,color:#fff
    classDef orange fill:#FF9800,stroke:#000,stroke-width:2px,color:#fff
    classDef purple fill:#9C27B0,stroke:#000,stroke-width:2px,color:#fff
```

## Directories:
* `api` - Contains the gRPC/proto definitions
* `envoy` - Contains the envoy configuration for grpc-web to grpc proxying
* `time` - time microservice in Go, listens on port 9090 for gRPC connections
* `frontend-ts` - simple vuejs 3 application using connect-es plugins to communicate with the backend using gRPC-Web

## Usage:
* `make proto` - Generate proto clients
* `make run-frontend` - Start frontend
* `make run-servers` - Start Time services and Envoy proxy
