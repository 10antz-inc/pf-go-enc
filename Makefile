.ONESHELL:

.PHONY: check
check:
	go fmt ./...
	golangci-lint run
	go test -cover ./... -coverprofile=docs/cover/cover.out
	go tool cover -html=docs/cover/cover.out -o docs/cover/cover.html