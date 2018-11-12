TARGET_CLI = ns

GO      	= go
GOLINT  	= $(GOPATH)/bin/golint
GO_SUBPKGS 	= $(shell $(GO) list ./... | grep -v /vendor/ | sed -e "s!$$($(GO) list)!.!")

BINDIR = ./bin/

GLIDE_VERSION := $(shell glide --version 2>/dev/null)
DEP_VERSION := $(shell dep version 2>/dev/null)

UPX := $(shell upx --version 2>/dev/null)

all: $(TARGET_CLI)

$(TARGET_CLI): build
ifdef UPX
	upx --brute $(BINDIR)$@
endif

build: vendor clean $(BINDIR)
	$(GO) build -ldflags="-s -w" -o $(BINDIR)$(TARGET_CLI) ./cmd/cli/*.go

vendor:
ifdef GLIDE_VERSION
	glide install
else
	go get .
endif

clean:
	rm -f $(BINDIR)*

$(BINDIR):
	mkdir -p $(BINDIR)

test: vendor
	$(GO) test -race $$($(GO) list ./...)

$(GOLINT):
	$(GO) get -u golang.org/x/lint/golint

lint: $(GOLINT)
	@for f in $(GO_SUBPKGS) ; do $(GOLINT) $$f ; done

.PHONY:test lint vendor
