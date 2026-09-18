// tools/generate.go
package generate

//go:generate protoc -I proto --go_out=proto/generated --go_opt=paths=source_relative --go-grpc_out=proto/generated --go-grpc_opt=paths=source_relative kvstore.proto
