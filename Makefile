build:
	cd cmd/feelthemovies &&	env GOOS=linux GOARCH=amd64 go build -mod vendor -o ../../dist/feelthemovies