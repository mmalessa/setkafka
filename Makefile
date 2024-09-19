
DC = docker compose

.DEFAULT_GOAL      = help

.PHONY: help
help:
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' Makefile | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

### DEV
.PHONY: build
build: build-dev

.PHONY: build-dev
build-dev: ## Build dev image
	$(DC) build

.PHONY: build-prod
build-prod: ## Build prod image
	@TARGET=prod $(DC) build


binary: ## build binary file /bin/setkafka
	@$(DC) exec go bash -c "go build -o bin/setkafka"

up: ## Start the project docker containers
	@$(DC) up -d

down: ## Down the docker containers
	@$(DC) down --timeout 25

shell: ## Run shell in go container
	@$(DC) exec -it -u appuser go bash
