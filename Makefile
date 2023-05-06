PROTOC_CMD = protoc
PROTOC_ARGS = --proto_path=./api/v2/ --go_opt=paths=source_relative --go_out=./go/kvspec/ --go-grpc_out=./go/kvspec/ ./api/v2/kvspec.proto

HTOML_TAG_FIX_CMD = htoml-tag-fix
HTOML_TAG_FIX_ARGS = go/kvspec/kvspec.pb.go

BUILDCOLOR="\033[34;1m"
BINCOLOR="\033[37;1m"
ENDCOLOR="\033[0m"

ifndef V
	QUIET_BUILD = @printf '%b %b\n' $(BUILDCOLOR)BUILD$(ENDCOLOR) $(BINCOLOR)$@$(ENDCOLOR) 1>&2;
	QUIET_INSTALL = @printf '%b %b\n' $(BUILDCOLOR)INSTALL$(ENDCOLOR) $(BINCOLOR)$@$(ENDCOLOR) 1>&2;
endif


all: go-v2
	@echo ""
	@echo "build complete"
	@echo ""

go-v2:
	# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$(QUIET_BUILD)go install github.com/golang/protobuf/protoc-gen-go$(CCLINK)
	$(QUIET_BUILD)go install google.golang.org/grpc/cmd/protoc-gen-go-grpc$(CCLINK)
	$(QUIET_BUILD)go install github.com/hooto/htoml4g/cmd/htoml-tag-fix$(CCLINK)
	$(QUIET_BUILD)$(PROTOC_CMD) $(PROTOC_ARGS) $(CCLINK)
	$(QUIET_BUILD)$(HTOML_TAG_FIX_CMD) $(HTOML_TAG_FIX_ARGS) $(CCLINK)

