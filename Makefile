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

plan: install
	@cd examples && terraform init && terraform plan

apply: install
	@cd examples && terraform init && terraform apply

destroy:
	@cd examples && terraform destroy

fmt:
	@gofmt -l -w -s .
	@terraform fmt -recursive .

install: build
	@mkdir -p $(TARGET); \
	install -m 0755 -p $(DIST)/${NAME} $(TARGET)/${NAME}
	@echo "installed"

uninstall:
	rm -rf $(TARGET)/${NAME}
