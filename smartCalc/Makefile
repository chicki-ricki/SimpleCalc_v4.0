format:
	gofmt -w *.go
	gofmt -w */*.go
#	goimports -local *.go

lint:
	golangci-lint run
	go vet

run:
	bee run