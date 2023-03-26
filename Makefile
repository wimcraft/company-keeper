lint:
	golangci-lint run

test:
	go test -v ./handler -parallel 4