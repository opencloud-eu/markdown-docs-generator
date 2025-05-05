OC_GIT_BRANCH ?= main

.PHONY: all
all: git-checkdirexist go-mod-tidy
	go run cmd/dochelpers/main.go all

.PHONY: templates
templates: go-mod-tidy
	go run cmd/dochelpers/main.go templates

.PHONY: rogue
rogue: go-mod-tidy
	go run cmd/dochelpers/main.go rogue

.PHONY: globals
globals: go-mod-tidy
	go run cmd/dochelpers/main.go globals

.PHONY: service-index
service-index: go-mod-tidy
	go run cmd/dochelpers/main.go service-index

.PHONY: env-var-delta-table
env-var-delta-table: go-mod-tidy
	go run cmd/dochelpers/main.go env-var-delta-table

.PHONY: git-checkdirexist
git-checkdirexist:
	@if [ ! -d "tmp" ]; then echo "Directory tmp does not exist. Please run `make git-clone` first.";exit 1;fi
	@if [ -z "$$(ls -A tmp)" ]; then echo "Directory tmp is empty. Please run `make git-clone` first.";exit 1;fi

.PHONY: git-clean
git-clean:
	rm -rf tmp

.PHONY: git-clone
git-clone: git-clean
	@if [ -d "tmp" ]; then echo "Directory tmp already exists. Please remove it before cloning.";exit 1;fi
	git clone -b "${OC_GIT_BRANCH}" https://github.com/opencloud-eu/opencloud.git tmp

.PHONY: clean
clean: gitclean output-clean

.PHONY: output-clean
output-clean:
	rm -Rfv output/*

.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy