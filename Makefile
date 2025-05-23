PROTO_SRC_DIR = proto/network_node
GO_OUT_DIR = go

# Protobuf import paths or other plugins you use
PROTO_INCLUDE = \
  -I$(PROTO_SRC_DIR)

# For gRPC code generation
PROTOC_GEN_GO = \
  --go_out=$(GO_OUT_DIR) \
  --go_opt=module=github.com/product-science/chain-protos/go \
  --go-grpc_out=$(GO_OUT_DIR) \
  --go-grpc_opt=module=github.com/product-science/chain-protos/go

# For gRPC-Gateway (if you’re using it)
PROTOC_GEN_GRPC_GATEWAY = \
  --grpc-gateway_out=$(GO_OUT_DIR) \
  --grpc-gateway_opt=module=github.com/product-science/chain-protos/go

proto-gen-go:
	protoc $(PROTO_INCLUDE) \
	  $(PROTOC_GEN_GO) \
	  $(PROTOC_GEN_GRPC_GATEWAY) \
	  $(PROTO_SRC_DIR)/**/*.proto

proto-gen-desc:
	protoc $(PROTO_INCLUDE) \
       --descriptor_set_out=$(GO_OUT_DIR)/network_node/v1/network_node.pb \
       $(PROTO_SRC_DIR)/v1/network_node.proto
