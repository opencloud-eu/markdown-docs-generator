OC_GIT_BRANCH ?= main
DOC_GIT_BRANCH ?= main

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
	@if [ ! -d "tmpdocs" ]; then echo "Directory tmpdocs does not exist. Please run `make git-clone` first.";exit 1;fi
	@if [ -z "$$(ls -A tmp)" ]; then echo "Directory tmp is empty. Please run `make git-clone` first.";exit 1;fi

.PHONY: git-clean
git-clean:
	rm -rf tmp
	rm -rf tmpdocs

.PHONY: git-clone
git-clone: git-clean
	@if [ -d "tmp" ]; then echo "Directory tmp already exists. Please remove it before cloning.";exit 1;fi
	git clone -b "${OC_GIT_BRANCH}" https://github.com/opencloud-eu/opencloud.git tmp
	@if [ -d "tmpdocs" ]; then echo "Directory tmpdocs already exists. Please remove it before cloning.";exit 1;fi
	gh repo clone opencloud-eu/docs tmpdocs -- -b "${OC_GIT_BRANCH}"; cd tmpdocs && git checkout -b docs-update-$$(uuidgen -r) && cd ..

.PHONY: clean
clean: gitclean output-clean

.PHONY: output-clean
output-clean:
	rm -Rfv output/*

.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy

.PHONY: create-docs-pullrequest
create-docs-pullrequest:
	cp -Rfv output/docs/* tmpdocs/static/env-vars/
	pushd tmpdocs && \
	git config --add --bool push.autoSetupRemote true && \
	git add * && \
	git commit -m "Update docs with latest env vars" && \
	git push && \
	gh pr create --title "Update docs" --body "This PR updates the documentation." --draft --label "Docs:Build&Tools"