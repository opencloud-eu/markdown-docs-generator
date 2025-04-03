module github.com/opencloud-eu/markdown-docs-generator

go 1.24.1

require (
	github.com/opencloud-eu/opencloud v1.1.0
	github.com/rogpeppe/go-internal v1.14.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/kr/pretty v0.3.1 // indirect
	golang.org/x/mod v0.24.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

// DO NOT CHANGE THIS REPLACE
// This ensures that the local cloned version is used instead of the one in the go module cache
replace github.com/opencloud-eu/opencloud => ./tmp/
