.PHONY: gen-go gen-python setup-go setup-python run-server run-client run-set run-get clean

# Default target
all: gen-go gen-python

# Setup directories for Go
setup-dirs:
	@mkdir -p github.com/productscience/chain-protos
	@mkdir -p network_node_direct/v1
	@mkdir -p python/network_node/v1

# Generate Go code from the proto files
gen-go: setup-dirs
	@echo "Generating Go code..."
	protoc --go_out=. --go-grpc_out=. network_node/v1/network_node.proto
	@echo "Copying to direct import location..."
	cp github.com/productscience/chain-protos/network_node/v1/* network_node_direct/v1/
	@echo "Go code generation complete"

# Generate Python code from the proto files
gen-python: setup-dirs
	@echo "Generating Python code..."
	mkdir -p python/network_node/v1
	touch python/network_node/__init__.py
	touch python/network_node/v1/__init__.py
	python -m grpc_tools.protoc -I. --python_out=python --grpc_python_out=python network_node/v1/network_node.proto
	@echo "Python code generation complete"

# Setup Go environment
setup-go:
	@echo "Setting up Go environment..."
	go mod init github.com/productscience/chain-protos || true
	cd server && go mod init github.com/productscience/chain-protos/server || true
	cd server && go mod edit -replace github.com/productscience/chain-protos=../
	cd server && go mod tidy
	@echo "Go environment setup complete"

# Setup Python environment
setup-python:
	@echo "Setting up Python environment..."
	python3 -m venv venv || true
	. venv/bin/activate && pip install grpcio grpcio-tools
	@echo "Python environment setup complete"

# Run the Go server
run-server:
	@echo "Running Go server..."
	cd server && go run main.go

# Run the Python client to set a record
run-set:
	@echo "Running Python client to set a record..."
	cd python && . ../venv/bin/activate && python client.py --action set --key mykey --value myvalue

# Run the Python client to get a record
run-get:
	@echo "Running Python client to get a record..."
	cd python && . ../venv/bin/activate && python client.py --action get --key mykey

# Run Python client with arguments
# Usage: make run-client ACTION=set KEY=mykey VALUE=myvalue
# or: make run-client ACTION=get KEY=mykey
run-client:
	@if [ "$(ACTION)" = "set" ] && [ -n "$(KEY)" ] && [ -n "$(VALUE)" ]; then \
		echo "Running Python client with action=$(ACTION), key=$(KEY), value=$(VALUE)"; \
		cd python && . ../venv/bin/activate && python client.py --action $(ACTION) --key $(KEY) --value $(VALUE); \
	elif [ "$(ACTION)" = "get" ] && [ -n "$(KEY)" ]; then \
		echo "Running Python client with action=$(ACTION), key=$(KEY)"; \
		cd python && . ../venv/bin/activate && python client.py --action $(ACTION) --key $(KEY); \
	else \
		echo "Usage for set: make run-client ACTION=set KEY=mykey VALUE=myvalue"; \
		echo "Usage for get: make run-client ACTION=get KEY=mykey"; \
		exit 1; \
	fi

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -rf github.com/productscience/chain-protos/network_node/v1/*.pb.go
	rm -rf network_node_direct/v1/*.pb.go
	rm -rf python/network_node/v1/*_pb2*.py
	@echo "Clean complete"

# Setup everything
setup: setup-dirs setup-go setup-python gen-go gen-python
	@echo "Setup complete"

# Help target
help:
	@echo "Available targets:"
	@echo "  gen-go         - Generate Go code from proto files"
	@echo "  gen-python     - Generate Python code from proto files"
	@echo "  setup-go       - Setup Go environment"
	@echo "  setup-python   - Setup Python environment (creates venv and installs dependencies)"
	@echo "  setup          - Setup everything (directories, Go, Python, and generate code)"
	@echo "  run-server     - Run the Go server"
	@echo "  run-set        - Run Python client with preset set action (key=mykey, value=myvalue)"
	@echo "  run-get        - Run Python client with preset get action (key=mykey)"
	@echo "  run-client     - Run Python client with custom arguments"
	@echo "                   Usage: make run-client ACTION=set KEY=mykey VALUE=myvalue"
	@echo "                   or: make run-client ACTION=get KEY=mykey"
	@echo "  clean          - Clean generated files"
	@echo "  help           - Show this help message" 