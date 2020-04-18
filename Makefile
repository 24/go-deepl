wire:
	GO111MODULE=off go get github.com/google/wire/cmd/wire
	wire ./...

test:
	go test -cover ./...

fmt:
	go fmt ./...

lint: fmt
	GO111MODULE=off go get golang.org/x/lint/golint
	go vet ./...
	golint --set_exit_status ./... || exit "$$?"

.PHONY: wire test fmt lint