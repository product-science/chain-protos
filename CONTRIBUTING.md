# Contributing to Chain Protos

Thank you for your interest in contributing to Chain Protos! This document provides guidelines and instructions for contributing to this project.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
3. Set up the development environment by installing the required tools:
   - Go (version 1.22 or later)
   - Protocol Buffer Compiler (protoc)
   - Go plugins for protoc

## Repository Structure

```
chain-protos/
├── network_node/            # Service definitions for network node
│   └── v1/                  # Version 1 of network node API
├── gen/                     # Generated code (committed to repo, do not edit manually)
├── examples/                # Usage examples
└── python/                  # Python-specific generated code
```

## Adding New Proto Files

When adding new proto files:

1. Place them in appropriate directories:
   ```
   <service-name>/<version>/<file-name>.proto
   ```

2. Follow naming conventions:
   - Service names should be snake_case
   - Messages should be PascalCase
   - Fields should be snake_case
   - Enums should be PascalCase with SCREAMING_SNAKE_CASE values

3. For Go code generation, include the `go_package` option:
   ```protobuf
   option go_package = "github.com/productscience/chain-protos/<service-name>/<version>;<packagename>";
   ```

4. Add appropriate documentation comments to all public elements

## Making Changes

1. Create a new branch for your changes:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes to the proto files
3. Generate code to ensure everything works:
   ```bash
   make gen-go
   ```
4. Review the generated code
5. **Important:** Commit both the proto changes AND the generated code to the repository
6. Update version and changelog according to semantic versioning:
   - For backwards-compatible additions, increment the minor version
   - For backwards-incompatible changes, increment the major version
   - Update CHANGELOG.md with a summary of your changes

## Generated Code Policy

This repository follows the approach of committing generated code to version control. This ensures that:

1. All users of the library can import the exact same generated code
2. We avoid version skew between different versions of the code generators
3. Users don't need to set up the code generation toolchain themselves

When making changes to proto files, always regenerate the code using `make gen-go` and commit the resulting changes to the `gen/` directory.

## Pull Request Process

1. Ensure your code follows the style guide
2. Update documentation if needed
3. Make sure all tests pass
4. Submit your pull request with:
   - A clear title and description
   - Reference to any relevant issues
   - Explanation of the changes made

## Versioning

We use [Semantic Versioning](https://semver.org/). For proto files, this means:

- MAJOR version increases for backwards-incompatible changes (removing fields, changing types, etc.)
- MINOR version increases for backwards-compatible additions (new fields, new methods, etc.)
- PATCH version increases for backwards-compatible fixes (documentation, comments, etc.)

## Code of Conduct

Please be respectful and considerate of others when contributing to this project. We are committed to providing a welcoming and inclusive environment.

## Communication

If you have any questions about contributing, please:
- Open an issue on GitHub
- Contact the maintainers

Thank you for contributing to Chain Protos! 