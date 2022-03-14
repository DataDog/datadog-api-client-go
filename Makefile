.PHONY: all
all: .generator
	@rm -rf api/*
	@pre-commit run --all-files --hook-stage=manual openapi-generator || true
	@rm -rf api/v*/datadog/api/ api/v*/datadog/go.mod api/v*/datadog/go.sum
	@pre-commit run --all-files --hook-stage=manual docs || echo "modified files"
	@pre-commit run --all-files --hook-stage=manual lint || echo "modified files"
	@rm -rf api/v*/datadog/docs/
