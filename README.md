# gRPC Types - unary, server, client, Bi-directional streaming Implementation

## Run Process

## 1. Generate code  
```go
protoc --go_out=. --go-grpc_out=. proto/greet.proto    
```

## Run Server
```go
cd server

go run *.go  
```

## Run Client
```go
cd client

go run *.go  
```
