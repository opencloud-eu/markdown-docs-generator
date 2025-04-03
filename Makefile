PHONY: all
all: gitcheckdirexist
	go run cmd/dochelpers/main.go all

PHONY: templates
templates:
	go run cmd/dochelpers/main.go templates

PHONY: rogue
rogue:
	go run cmd/dochelpers/main.go rogue

PHONY: globals
globals:
	go run cmd/dochelpers/main.go globals

PHONY: service-index
service-index:
	go run cmd/dochelpers/main.go service-index

PHONY: env-var-delta-table
env-var-delta-table:
	go run cmd/dochelpers/main.go env-var-delta-table

PHONY: gitcheckdirexist
gitcheckdirexist:
	@if [ ! -d "tmp" ]; then echo "Directory tmp does not exist. Please run `make gitclone` first.";exit 1;fi
	@if [ -z "$$(ls -A tmp)" ]; then echo "Directory tmp is empty. Please run `make gitclone` first.";exit 1;fi

PHONY: gitclean
gitclean:
	rm -rf tmp

PHONY: gitclone
gitclone:
	@if [ -d "tmp" ]; then echo "Directory tmp already exists. Please remove it before cloning.";exit 1;fi
	git clone https://github.com/opencloud-eu/opencloud.git tmp

PHONY: clean
clean: gitclean
	rm -Rfv output