.PHONY: test
test:
	@go test -v -count 1 ./...

.PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./...
