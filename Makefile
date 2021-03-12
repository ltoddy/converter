.PHONY: clean fmt test view_test_coverage

SRC = $(shell find . -type f -name '*.go')
PACKAGES = $(shell go list ./...)

clean:
	@go clean -cache
	@golangci-lint cache clean
	@rm -rf output

fmt:
	@gofmt -l -w $(SRC)
	@goimports -w $(SRC)

install-tools:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.2
	go get golang.org/x/tools/cmd/goimports
	git checkout go.mod
	rm go.sum

lint:
	@golangci-lint run ./...

test:
	@go test -v -parallel 1 -coverprofile=coverage.out $(PACKAGES)
	@go tool cover -func=coverage.out

view_test_coverage:
	@go tool cover -html=coverage.out
