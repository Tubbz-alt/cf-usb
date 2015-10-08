NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
DEPS=$(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
GOPKGS=$(shell go list -f '{{.Dir}}' ./...)

include version.mk

BUILD:=$(shell echo `whoami`-`git rev-parse --short HEAD`-`date -u +%Y%m%d%H%M%S`)
APP_VERSION=$(VERSION)-$(BUILD)
DIST_FIND_BUILDS=find * -type d -not -path "forpatches" -exec

.PHONY: all dist format lint vet build test tools bench clean generate cleangeneratedfiles
.SILENT: all dist format lint vet build test tools bench clean generate cleangeneratedfiles

all: clean format lint vet build test dist

format:
	@echo "$(OK_COLOR)==> Checking format$(ERROR_COLOR)"
	@echo $(PKGSDIRS) | tr ' ' '\n' | xargs -I '{p}' -n1 goimports -e -l {p} | sed "s/^/Failed: /"
	@echo "$(NO_COLOR)\c"

lint:
	@echo "$(OK_COLOR)==> Linting$(ERROR_COLOR)"
	@echo $(PKGSDIRS) | tr ' ' '\n' | xargs -I '{p}' -n1 golint {p} | grep -v "mock_.*\.go" | sed "s/^/Failed: /"
	@echo "$(NO_COLOR)\c"

vet:
	@echo "$(OK_COLOR)==> Vetting$(ERROR_COLOR)"
# Blame https://code.google.com/p/go/issues/detail?id=6820 for the -composites=false
	@echo $(GOPKGS) | tr ' ' '\n' | xargs -I '{p}' -n1 go tool vet -composites=false {p} | sed "s/^/Failed: /"
	@echo "$(NO_COLOR)\c"

build: generate #deps
	@echo "$(OK_COLOR)==> Building$(NO_COLOR)"
	export GOPATH=$(shell godep path):$(shell echo $$GOPATH) &&\
	gox -verbose \
	-ldflags="-X main.version $(APP_VERSION)" \
	-os="windows linux darwin " \
	-arch="amd64" \
	-output="build/{{.OS}}-{{.Arch}}/{{.Dir}}" ./...

patch: cleangeneratedfiles
	@echo "$(OK_COLOR)==> Generating Update Patches$(NO_COLOR)"
	export CGOENABLED=1 && \
	export GOPATH=$(shell godep path):$(shell echo $$GOPATH) &&\
	gox -verbose \
	-ldflags="-X main.version $(APP_VERSION) -X main.updateserver $(PATCH_SERVER) -X main.branch $(PATCH_CHANNEL)" \
	-os="windows linux darwin " \
	-arch="amd64" \
	-output="build/forpatches/{{.OS}}-{{.Arch}}" ./...
	mv build/forpatches/windows-amd64.exe build/forpatches/windows-amd64

	go-selfupdate "build/forpatches" $(APP_VERSION)
	ln -s public cf-mgmt_$(MAJOR_MINOR)_$(PATCH_CHANNEL)

cleangeneratedfiles:
	rm -rf build/forpatches
	rm -rf public

test:
	@echo "$(OK_COLOR)==> Testing$(NO_COLOR)"
	export GOPATH=$(shell godep path):$(shell echo $$GOPATH) &&\
	gocov test ./... | gocov-xml > coverage.xml
	@echo "$(NO_COLOR)\c"

tools:
	@echo "$(OK_COLOR)==> Installing tools$(NO_COLOR)"
	#Great tools to have, and used in the build file
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/tools/cmd/vet
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/golang/lint/golint
	#Tools for the project
	go get -u github.com/codegangsta/cli
	go get -u github.com/tools/godep
	go get -u github.com/mitchellh/gox
	go get -u github.com/vektra/mockery/cmd/mockery
	go get -u github.com/axw/gocov/gocov

	# gox -build-toolchain
	#dependencies for project
	go get gopkg.in/yaml.v2
	go get github.com/hpcloud/loom
	go get github.com/stretchr/testify/assert
	go get github.com/stretchr/testify/mock
	#Tools for code coverage reporting
	go get github.com/axw/gocov/...
	go get github.com/AlekSi/gocov-xml
	go get gopkg.in/matm/v1/gocov-html

clean: cleangeneratedfiles
	@echo "$(OK_COLOR)==> Cleaning$(NO_COLOR)"
	rm -rf build
	rm -rf $(GOPATH)/pkg/*
	rm -f $(GOPATH)/bin/als