NAME   := terraform-provider-slack
OS     := $(shell go env GOOS)
ARCH   := $(shell go env GOARCH)
DIST   := dist/${OS}_${ARCH}
TARGET := ${HOME}/.terraform.d/plugins/${OS}_${ARCH}
FILES  := $(shell find . -type f -name '*.go')

.DEFAULT_GOAL := build
.PHONY: build clean install uninstall

${DIST}/${NAME}: ${FILES}
	go build -o $(DIST)/${NAME}

build: ${DIST}/${NAME}

clean:
	rm -rf ${DIST}/*

install: build
	mkdir -p $(TARGET); \
	install -m 0755 $(DIST)/${NAME} $(TARGET)/${NAME}

uninstall:
	rm -rf $(TARGET)/${NAME}
