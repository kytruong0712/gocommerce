# syntax=docker/dockerfile:1

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.19

# create a working directory inside the image
WORKDIR /gocomm-api

# Copy and download dependency using go mod.
COPY api/go.mod .
COPY api .

RUN go mod download

# compile application
# /gocommapi: directory stores binaries file
# ./cmd/serverd/main.go: lookups the main.go file by the path started with gocommapi/cmd/...
# RUN go build -o /gocomm ./cmd/serverd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /gocommapi ./cmd/serverd/main.go

# tells Docker that the container listens on specified network ports at runtime
EXPOSE 3000

# command to be used to execute when the image is used to start a container
CMD ["/gocommapi"]
