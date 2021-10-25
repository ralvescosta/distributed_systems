### To compile the .proto files we need 

- protoc cli compiler for unix
- go-proto plugin: ➜ go install github.com/golang/protobuf/protoc-gen-go
- to compile run this command: ➜ protoc --go_out=plugins=grpc:proto inventory.proto