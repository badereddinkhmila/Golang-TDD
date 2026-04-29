test:
	@GOFLAGS="-count=1" go test -v -cover -race -coverprofile=coverage.out ./...

escape:
	go build -gcflags="-m"
