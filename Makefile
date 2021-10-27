.PHONY: all
all: .generator .env
	@docker-compose -f docker-compose.generator.yaml up
	@rm api/v*/datadog/go.mod api/v*/datadog/go.sum
	@pre-commit run --all-files --hook-stage=manual lint || echo "modified files"

.PHONY: .env
.env:
	@echo "export UID=$(shell id -u)\nexport GID=$(shell id -g)" > $@