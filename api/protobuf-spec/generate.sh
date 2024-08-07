protoc --go_out=. hello.proto
protoc --go-grpc_out=. hello.proto

#protoc --go_out=plugins=grpc:. hello.proto
