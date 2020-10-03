build:
	@docker build -t mollie-go:latest -f Dockerfile .
.PHONY: build

run:
	@docker run --rm mollie-go:latest
.PHONY: run

lint:
	@go version

	@echo "Running go lint"
	@golint ./...

	@echo "Running go vet"
	@go vet ./...
.PHONY: lint

test: run
.PHONY: test

test-local:
	@go test -v ./mollie/... -coverprofile cover.out
.PHONY: test-local

coverage:
	@go test ./mollie/... -coverprofile cover.out
	@go tool cover -func=cover.out
.PHONY:  coverage

clean:
	@go mod verify
	@go mod tidy
.PHONY: clean