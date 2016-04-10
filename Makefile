PROJECT=tinker
ORGANIZATION=crisidev

SOURCE := $(shell find . -name '*.go')
VERSION := $(shell cat VERSION)
COMMIT := $(shell git rev-parse --short HEAD)
GOPATH := $(shell pwd)/.gobuild
PROJECT_PATH := $(GOPATH)/src/github.com/$(ORGANIZATION)
GOVERSION := "1.5.3"

ifndef GOOS
	GOOS := $(shell go env GOOS)
endif
ifndef GOARCH
	GOARCH := $(shell go env GOARCH)
endif

.PHONY: all clean run-tests deps install fmt

all: deps $(PROJECT)

ci: clean all run-tests

clean:
		rm -rf $(GOPATH) $(PROJECT)

docker-run-tests:
	@echo Testing in Docker for $(GOOS)/$(GOARCH)
	docker run \
	    --rm \
	    -v $(shell pwd):/usr/code \
	    -e GOPATH=/usr/code/.gobuild \
	    -e GOOS=$(GOOS) \
	    -e GOARCH=$(GOARCH) \
	    -w /usr/code \
	    golang:${GOVERSION} \
	    go test

run-tests:
	@echo Testing for $(GOOS)/$(GOARCH)
	go test -v -cover

deps:
	mkdir -p $(PROJECT_PATH)
	cd $(PROJECT_PATH) && ln -nsf ../../../.. $(PROJECT)
	GOPATH=$(GOPATH) go get github.com/${ORGANIZATION}/$(PROJECT)

# build
docker-$(PROJECT): $(SOURCE) VERSION
	@echo Building in Docker for $(GOOS)/$(GOARCH)
	docker run \
	    --rm \
	    -v $(shell pwd):/usr/code \
	    -e GOPATH=/usr/code/.gobuild \
	    -e GOOS=$(GOOS) \
	    -e GOARCH=$(GOARCH) \
	    -w /usr/code \
	    golang:$(GOVERSION) \
	    go build -a -ldflags "-X main.projectVersion=$(VERSION) -X main.projectBuild=$(COMMIT)" -o $(PROJECT)

$(PROJECT): $(SOURCE) VERSION
	@echo Building for $(GOOS)/$(GOARCH)
	go build -a -ldflags "-X main.projectVersion=$(VERSION) -X main.projectBuild=$(COMMIT)" -o $(PROJECT)


install: $(PROJECT)
	install -m 0755 $(PROJECT) /usr/local/bin/

fmt:
	gofmt -l -w .
