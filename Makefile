run:
	cd api && go run cmd/serverd/main.go

build:
	cd api && go build -o bin/main cmd/serverd/main.go

test:
	cd api && go test ./...

docker-build:
	docker build -t gocomm-api .

docker-run:
	docker run --name gocomm-api -p 3000:3000 -h 0.0.0.0 gocomm-api
