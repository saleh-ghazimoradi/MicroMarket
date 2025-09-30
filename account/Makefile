generate-grpc:
	protoc --go_out=. --go-grpc_out=. gateway/proto/account.proto

docker-up:
	docker compose up -d

docker-down:
	docker compose down

fmt:
	go fmt ./...

vet:
	go vet ./...

run: fmt vet
	go run main.go

test:
	go test -v ./...