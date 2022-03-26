// proto defines a gRPC interface for interaction with the backend
// API.
package proto

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./main.proto
