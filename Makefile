.PHONY: all test clean

containers-up:
	clear && docker-compose down && docker-compose up -d

dependencies:
	go get -t -v ./... 

build:
	cd cmd/feelthemovies &&	env GOOS=linux GOARCH=amd64 go build -o ${HOME}/feelthemovies

test: 
	clear && go test -v ./...