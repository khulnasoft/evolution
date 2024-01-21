SHELL=/bin/bash -o pipefail

GO ?= go

all: bin/utils

clean:
	@rm -fr bin

bin/utils:
	@mkdir -p bin
	@$(GO) build -o bin/utils ./cmd/...
