.PHONY: gen lint test release tag clean help

# Generate code using buf
gen:
	@echo "Generating Go code..."
	@mkdir -p gen
	buf generate
	@echo "Go code generation complete"
	@echo "Note: The generated Go code should be committed to the repository"

# Lint proto files
lint:
	@echo "Linting proto files..."
	buf lint
	@echo "Linting complete"

# Run all tests
test:
	@echo "Running tests..."
	go test -v ./...
	@echo "Tests complete"

# Tag the current state with the version in VERSION file
tag:
	@VERSION=$$(cat VERSION); \
	echo "Tagging version v$$VERSION"; \
	git tag -a "v$$VERSION" -m "Release v$$VERSION"; \
	echo "To push the tag, run: git push origin v$$VERSION"

# Release a new version - updates VERSION and CHANGELOG
# Usage: make release VERSION=1.0.0
release:
	@if [ -z "$(VERSION)" ]; then \
		echo "Please specify a version: make release VERSION=1.0.0"; \
		exit 1; \
	fi; \
	CURRENT_VERSION=$$(cat VERSION); \
	echo "Updating from v$$CURRENT_VERSION to v$(VERSION)"; \
	echo "$(VERSION)" > VERSION; \
	DATE=$$(date +%Y-%m-%d); \
	sed -i.bak "0,/^## \[.*\]/s/^## \[.*\]/## [$(VERSION)] - $$DATE\n\n### Added\n- \n\n## [$$CURRENT_VERSION]/" CHANGELOG.md; \
	rm CHANGELOG.md.bak; \
	echo "Updated VERSION and CHANGELOG.md. Review the changes, commit, and then run 'make tag' to tag this release."

# Clean generated files (but remember they should be committed normally)
clean:
	@echo "Cleaning generated files..."
	rm -rf gen
	@echo "Clean complete"

# Default target shows help
.DEFAULT_GOAL := help

# Help target
help:
	@echo "Available targets:"
	@echo "  gen       - Generate Go code using buf (commit the output)"
	@echo "  lint      - Lint proto files with buf" 
	@echo "  test      - Run all tests"
	@echo "  release   - Update VERSION and CHANGELOG for a new release"
	@echo "             Usage: make release VERSION=1.0.0"
	@echo "  tag       - Tag the repository with the version in VERSION file"
	@echo "  clean     - Clean generated files (use with caution)"
	@echo "  help      - Show this help message" 