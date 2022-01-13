.PHONY: all
all: .generator .env
	@pre-commit run --all-files --hook-stage=manual openapi-generator || true
	@pre-commit run --all-files --hook-stage=manual docs || echo "modified files"
	@pre-commit run --all-files --hook-stage=manual lint || echo "modified files"

.PHONY: .env
.env:
	@echo "export UID=$(shell id -u)\nexport GID=$(shell id -g)" > $@
