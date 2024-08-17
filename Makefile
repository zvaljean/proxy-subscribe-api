MAKEFLAGS += --no-builtin-rules
BINS = proxy
BUILD_PATH := $(shell basename $$(pwd))
BIN_EXTENSION :=
ifeq ($(OS), windows)
  BIN_EXTENSION := .exe
endif

OUTBINS = $(foreach bin,$(BINS), $(OS)_$(ARCH)/$(bin)$(BIN_EXTENSION))

GO_CMD := $(shell go env GOROOT)/bin/go

VERSION ?= $(shell git describe --tags --always --dirty)
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))
TAG := $(VERSION)__$(OS)_$(ARCH)
CGO_ENABLED = 0

SHELL := /usr/bin/env bash -o errexit -o pipefail -o nounset
GOFLAGS = $(GOFLAGS) -modcacherw
PHONY_TARGET = clean
.PHONY: $(PHONY_TARGET)

all: install

build: $(addprefix build-, $(BINS) )
install: $(addprefix install-, $(BINS) )
clean: $(addprefix clean-, $(BINS) )

build-%:
	$(info --------begin build--------)
	GOARCH=$(ARCH) GOOS=$(OS) CGO_ENABLED=$(CGO_ENABLED) \
	$(GO_CMD) build -o ./bin/$* ./cmd/$*

install-%: build
	$(info --------begin intall--------)
	$(GO_CMD) install ./cmd/$*

clean-%:
	rm -rf ./bin/$(OS)_$(ARCH)/*
	$(GO_CMD) clean -i  ./cmd/$*
