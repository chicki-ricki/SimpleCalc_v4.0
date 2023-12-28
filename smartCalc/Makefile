GO=go
# Go=go1.19.2

format:
	@gofmt -w *.go
	@gofmt -w */*.go
#	goimports -local *.go

lint:
	@golangci-lint run
	@${GO} vet

run:
	bee run 

test: testing
	
cover:
	@${GO} test -covermode=count -coverpkg=./model -coverprofile ./tests/cover.out -v ./model > /dev/null
	@${GO} tool cover -html ./tests/cover.out -o ./tests/cover.html > /dev/null
	@open ./tests/cover.html
#	@yandex-browser cover.html > /dev/null 
#	firefox cover.html
#	code cover.html

testing:
	@${GO} test ./model -v -cover 