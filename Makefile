.PHONY: help setup clean codegen

help: ## show this help
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

setup: ## install dependencies
	cd host && go get .
	cd codegen && go get .

# no help here, because this stubs the host (which will be hand-edited)
codegen:
	cd codegen && go run main.go

null0: host/null0.go ## build native host
	cd host && go build -o ../null0 .

CART_DIRS    := $(wildcard carts/*)
CART_NAMES   := $(notdir $(CART_DIRS))
CART_TARGETS := $(addsuffix .null0,$(CART_NAMES))

%.null0: carts/%/*.go
	cd carts/$* && tinygo build -target wasm -o main.wasm main.go && zip -r ../../$@ . -x "*.go" -x ".DS_Store" -x "*.mod"

carts: $(CART_TARGETS) ## build example carts

clean: ## clean up any built files
	rm -f carts/*/main.wasm null0 *.null0