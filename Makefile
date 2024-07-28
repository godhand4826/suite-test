mocks:
	mockery

.PHONY: test
test: mocks
	go test -v ./...
