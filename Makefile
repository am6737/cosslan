
MODEL_LIST = user line node

.PHONY: run
run:
	@go run ./cmd/main.go

#.PHONY: gen-openapi
#gen-openapi: $(addprefix gen-openapi-, $(MODEL_LIST))
#
#gen-openapi-%:
#	-@oapi-codegen -generate echo,spec -o internal/api/http/v1/$*/openapi.gen.go -package v1 internal/api/http/v1/$*/$*.yaml
#	-@oapi-codegen -generate types -o internal/api/http/v1/$*/openapi.types.go -package v1 internal/api/http/v1/$*/$*.yaml

.PHONY: gen-openapi
gen-openapi:
	-@oapi-codegen -generate echo,spec -o internal/api/http/v1/openapi.gen.go -package v1 docs/cosslan.yaml
	-@oapi-codegen -generate types -o internal/api/http/v1/openapi.types.go -package v1 docs/cosslan.yaml


# 定义变量
INPUT_DIRS := $(wildcard  ./internal/api/http/v1/*)
OUTPUT_FILE := ./docs/cosslan.yaml
SWAGGER_MERGER := $(shell command -v go-swagger-merger)

#go get github.com/efureev/go-swagger-merger
#go install github.com/efureev/go-swagger-merger
# 定义合并 OpenAPI 的目标
.PHONY: merge-openapi
merge-openapi: gen-openapi
    # 检查是否存在 go-swagger-merger 命令
	go-swagger-merger -o $(OUTPUT_FILE) $(foreach dir,$(INPUT_DIRS),-i $(dir)/*.yaml)

# Go code format
.PHONY: fmt
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...

# Go code vet
.PHONY: vet
vet:
	@echo "Checking Go code with go vet..."
	@go vet ./...


# Run all checks and formats
.PHONY: check
check: vet fmt
	@echo "All checks and formatting done."
