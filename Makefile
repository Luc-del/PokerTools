lint: ## Run the linter
	golangci-lint run

unit: ## Run unit tests
	go test -race ./...
