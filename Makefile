## =========================== Variables =================================== ##

SHELL := /bin/bash -o pipefail

export GOBIN  := $(CURDIR)/bin

## ============================== Commands =================================== ##


DAGGER_CMD     := $(CURDIR)/bin/dagger
DAGGER_VERSION := 0.6.1

GALE_CMD     := $(CURDIR)/bin/gale
GALE_VERSION := main

## ============================== Targets ==================================== ##

.PHONY: lint
lint: $(DAGGER_CMD) $(GALE_CMD);
	@$(DAGGER_CMD) run $(GALE_CMD) --workflow ci --job lint --disable-checkout --export

.PHONY: test
test: $(DAGGER_CMD) $(GALE_CMD);
	@$(DAGGER_CMD) run $(GALE_CMD) --workflow ci --job test --disable-checkout --export

.PHONY: build
build: $(DAGGER_CMD) $(GALE_CMD);
	@$(DAGGER_CMD) run $(GALE_CMD) --workflow ci --job build --disable-checkout --export

## ============================ Dependencies ================================= ##

$(DAGGER_CMD):
	@curl -L https://dl.dagger.io/dagger/install.sh | VERSION=$(DAGGER_VERSION) sh

$(GALE_CMD):
	@go install -v github.com/aweris/gale@$(GALE_VERSION)