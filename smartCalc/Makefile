# GO=go
Go=go1.19.2

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
	${GO} test ./... -v -cover 

cover:
	${GO} test -covermode=count -coverpkg=./model -coverprofile ./tests/cover.out -v ./model
	${GO} tool cover -html ./tests/cover.out -o ./tests/cover.html
	open ./tests/cover.html	
