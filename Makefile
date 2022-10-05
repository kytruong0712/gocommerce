run:
	cd api && go run cmd/serverd/main.go

build:
	cd api && go build -o bin/main cmd/serverd/main.go

test:
	cd api && go test ./...
