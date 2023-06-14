## =========================== Variables =================================== ##

SHELL := /bin/bash -o pipefail

export GOBIN  := $(CURDIR)/bin

## ============================== Commands =================================== ##


DAGGER_CMD     := $(CURDIR)/bin/dagger
DAGGER_VERSION := 0.6.1

GALE_CMD     := $(CURDIR)/bin/gale
GALE_VERSION := main

MAGE_CMD     := $(CURDIR)/bin/mage
MAGE_VERSION := latest

## ============================== Targets ==================================== ##

.PHONY: lint
lint: $(DAGGER_CMD) $(GALE_CMD);
	@$(DAGGER_CMD) run $(GALE_CMD) run ci lint --disable-checkout --export

.PHONY: test
test: $(DAGGER_CMD) $(GALE_CMD);
	@$(DAGGER_CMD) run $(GALE_CMD) run ci test --disable-checkout --export

.PHONY: build
build: $(DAGGER_CMD) $(GALE_CMD);
	@$(DAGGER_CMD) run $(GALE_CMD) run ci build --disable-checkout --export

.PHONY: mage.lint
mage.lint: $(DAGGER_CMD) $(MAGE_CMD)
	@$(DAGGER_CMD) run $(MAGE_CMD) lint

.PHONY: mage.test
mage.test:$(DAGGER_CMD) $(MAGE_CMD)
	@$(DAGGER_CMD) run $(MAGE_CMD) test

.PHONY: mage.build
mage.build: $(DAGGER_CMD) $(MAGE_CMD)
	@$(DAGGER_CMD) run $(MAGE_CMD) build

.PHONY: mage.all
mage.all: $(DAGGER_CMD) $(MAGE_CMD)
	@$(DAGGER_CMD) run $(MAGE_CMD) go:all

## ============================ Dependencies ================================= ##

$(DAGGER_CMD):
	@curl -L https://dl.dagger.io/dagger/install.sh | VERSION=$(DAGGER_VERSION) sh

$(GALE_CMD):
	@go install -v github.com/aweris/gale@$(GALE_VERSION)

$(MAGE_CMD):
	@go install -v github.com/magefile/mage@$(MAGE_VERSION)