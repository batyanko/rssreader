setup: ## setup tools for testing and linting
	go get github.com/stretchr/testify
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

test: ## run unit tests
	go test github.com/batyanko/rssreader/...

lint: ## run lint check
	golangci-lint run --enable-all --disable=wsl,godox,gosec,lll,scopelint
