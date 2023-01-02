build:
	@docker build -t mollie-go:latest -f Dockerfile .
.PHONY: build

run:
	@docker run --rm mollie-go:latest
.PHONY: run

ci-lint:
    # colored-line-number|line-number|json|tab|checkstyle|code-climate|junit-xml|github-actions
	@golangci-lint run --out-format=tab --issues-exit-code=0 --sort-results --skip-dirs-use-default --tests=false --presets=bugs,complexity,format,performance,style,unused
.PHONY: ci-lint

pr-lint:
	@golangci-lint run --issues-exit-code=0 --out-format=github-actions --new=true --sort-results --skip-dirs-use-default --tests=false --presets=bugs,complexity,format,performance,style,unused
.PHONY: pr-lint

ci-all-presets:
	@golangci-lint run --out-format=tab --issues-exit-code=0 --sort-results --skip-dirs-use-default --tests=false --presets=bugs,comment,complexity,error,format,import,metalinter,module,performance,sql,style,test,unused
.PHONY: ci-all-presets

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
	@go test -v -failfast ./mollie/... -coverprofile cover.out
.PHONY: test-local

coverage: test-local
	@go tool cover -func=cover.out
.PHONY:  coverage

cover-report: test-local
	@go tool cover -html=cover.out
.PHONY: cover-report

clean:
	@go mod verify
	@go mod tidy
.PHONY: clean

update-docs:
	@godocdown ./mollie > docs/README.md 
.PHONY: update-docs