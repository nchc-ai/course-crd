IMAGE := ghcr.io/slok/kube-code-generator:v0.2.0

DIRECTORY := $(PWD)
DEPS_CMD := go mod tidy

default: generate

.PHONY: generate
generate:
	@docker run -it --rm -v $(DIRECTORY):/app $(IMAGE) \
	--apis-in ./pkg/apis \
	--go-gen-out ./pkg/client \
	--crd-gen-out ./manifests 

.PHONY: deps
deps:
	$(DEPS_CMD)

.PHONY: clean
clean:
	echo "Cleaning generated files..."
	rm -rf ./manifests
	rm -rf ./pkg/client
	rm -rf ./pkg/apis/coursecontroller/v1alpha1/zz_generated.deepcopy.go