.DEFAULT_GOAL	:= build

#------------------------------------------------------------------------------
# Variables
#------------------------------------------------------------------------------

SHELL 	:= /bin/bash
BINDIR	:= bin
PKG 	:= github.com/envoyproxy/meta-protocol-control-plane-api

.PHONY: build
build:
	@go build ./pkg/... ./meta_protocol_proxy/...

.PHONY: format
format:
	@goimports -local $(PKG) -w -l pkg

#--------------------------------------
#-- example xDS control plane server
#--------------------------------------
.PHONY: $(BINDIR)/example example

$(BINDIR)/example:
	@go build -race -o $@ example/main/main.go

example: $(BINDIR)/example