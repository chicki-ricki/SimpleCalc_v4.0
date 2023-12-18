GO=go
#Go=go1.19.2

format:
	gofmt -w *.go
	gofmt -w */*.go
#	goimports -local *.go

lint:
	golangci-lint run
	${GO} vet

run:
	bee run

test: testing
	

testing:
	${GO} test ./... -cover 