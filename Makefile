test:
	go test -v -cover ./...

lint:
	golangci-lint run ./...

check:
	staticcheck ./...

ready: check lint test
