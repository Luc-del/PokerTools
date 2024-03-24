lint: ## Run the linter
	golangci-lint run

unit: ## Run unit tests
	go test -race ./...

gen: ## Generate files
	go generate ./...