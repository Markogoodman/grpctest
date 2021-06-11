# grpctest

protoc-gen-go v1.26.0
protoc v3.15.8

// gen code
protoc --go_out=. --go_opt=paths=source_relative \
 --go-grpc_out=. --go-grpc_opt=paths=source_relative \
 proto/helloworld.proto

grpcurl

interceptor: like middleware
