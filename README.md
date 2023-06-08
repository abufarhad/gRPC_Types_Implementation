# gRPC Types - unary, server, client, Bi-directional streaming Implementation

## See the demo here - https://youtu.be/FnG5_S9dyN4 

# Run Process

## Generate code  
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
