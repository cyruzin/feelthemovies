# Start from golang v1.12.x base image.
FROM golang:1.12.4 as build-stage

# Set the Current Working Directory inside the container.
WORKDIR /go/src/github.com/cyruzin/feelthemovies

# Copy everything from the current directory to the 
# PWD (Present Working Directory) inside the container.
COPY . .

# Download all the dependencies.
RUN go get -d -v ./...

# Build the Go App.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/feelthemovies

# Start a new stage from scratch.
FROM alpine:latest  

# Add certificates.
RUN apk add ca-certificates

# Copy the Pre-built binary file from the previous stage.
COPY --from=build-stage /go/bin/feelthemovies .

# This container exposes port 8000 to the outside world.
EXPOSE 8000

# Run the executable.
CMD ["./feelthemovies"]