tests:
	echo Waiting for MySQL container...
	docker-compose down && docker-compose up -d
	sleep 10
	go test -v ./... -race