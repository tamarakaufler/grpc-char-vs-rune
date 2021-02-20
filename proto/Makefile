# Ensure protoc is on the path
PROTOC_INSTALLED := $(shell command -v protoc;)
ifndef PROTOC_INSTALLED
	$(error Install protoc from https://github.com/protocolbuffers/protobuf/releases)
endif

.PHONY: protogen
protogen:
	@echo "Generating protos"
	prototool lint
	prototool generate
	$(MAKE) -C ../client mocks
