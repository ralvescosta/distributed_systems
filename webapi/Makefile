compile-proto:
	protoc --go_out=plugins=grpc:proto ./pkg/infra/grpc_clients/proto/inventory.proto

run:
	GO_ENV=development GIN_MODE=debug go run main.go

unit-test:
	GO_ENV=test GIN_MODE=debug go test -race ./pkg/... -v

test-cov:
	if ! [ -d "coverage" ]; then \
		echo "Creating covorage folder" ; \
		mkdir coverage; \
	fi
	GO_ENV=development go test ./... -cover -v -coverprofile ./coverage/c.out && go tool cover -html=./coverage/c.out -o ./coverage/coverage.html

build:
	go build -ldflags "-s -w" main.go

private-key:
	if ! [ -d "cert" ]; then \
		echo "Creating covorage folder" ; \
		mkdir cert; \
	fi
	openssl genrsa -out cert/id_rsa 4096

public-key:
	if ! [ -d "cert" ]; then \
		echo "Creating covorage folder" ; \
		mkdir cert; \
	fi
	openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub