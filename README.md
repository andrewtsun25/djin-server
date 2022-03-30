# djin-server

This is the [gRPC](https://grpc.io/) server implementation of the website [djin.dev](https://djin.dev). 

The server can be accessed under port `localhost:8080` when developing locally

# Frequently used commands

## Building the server
Build an executable version of this server that can be run. 
```go
go build ./...
```

## Running the server
Run this server in listening mode.
```go
go run ./...
```

## Installing the server
Install the server as an executable on your machine.
```go
go install ./...
```

## Tests 
Run all unit tests for this server.
```go
go test ./...
```
