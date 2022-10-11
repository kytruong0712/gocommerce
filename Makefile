run:
	cd api && go run cmd/serverd/main.go

build:
	cd api && go build -o bin/main cmd/serverd/main.go

test:
	cd api && go test ./...

docker-build:
	docker build -t gocomm-api .

docker-run:
	docker run -d -p 3000:3000 gocomm-api:latest

docker-multistage-build:
	docker build -t gocomm-api:multistage -f Dockerfile.multistage .

docker-multistage-run:
	docker run -d -p 3000:3000 gocomm-api:multistage
