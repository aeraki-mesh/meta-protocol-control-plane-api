.DEFAULT_GOAL	:= build

#------------------------------------------------------------------------------
# Variables
#------------------------------------------------------------------------------

SHELL 	:= /bin/bash
BINDIR	:= bin
PKG 	:= github.com/envoyproxy/go-control-plane

.PHONY: build
build:
	@go build ./pkg/... ./api/...

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