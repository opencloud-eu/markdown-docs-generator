package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/opencloud-eu/markdown-docs-generator/pkg/dochelpers"
)

const devMode = false

func main() {

	if len(os.Args) > 1 {
		doClone := true
		if devMode {
			if _, err := os.Stat("tmp/services"); os.IsNotExist(err) {
				doClone = false
			}
		}
		if doClone {
			// Check if the tmp directory exists, if so, remove it
			os.RemoveAll("tmp/")
			// Clone the opencloud repository into the tmp directory
			_, err := git.PlainClone("tmp/", false, &git.CloneOptions{
				URL:      "https://github.com/opencloud-eu/opencloud.git",
				Progress: os.Stdout,
			})
			if err != nil {
				fmt.Println("Error cloning repo:", err)
				return
			}
		}
		switch os.Args[1] {
		case "templates":
			dochelpers.RenderTemplates()
		case "rogue":
			dochelpers.GetRogueEnvs()
		case "globals":
			dochelpers.RenderGlobalVarsTemplate()
		case "service-index":
			dochelpers.GenerateServiceIndexMarkdowns()
		case "env-var-delta-table":
			// This step is not covered by the all or default case, because it needs explicit arguments
			if len(os.Args) != 4 {
				fmt.Println("Needs two arguments: env-var-delta-table <first-version> <second-version>")
				fmt.Println("Example: env-var-delta-table v5.0.0 v6.0.0")
				fmt.Println("Will not generate usable results for versions Prior to v5.0.0")
			} else {
				dochelpers.RenderEnvVarDeltaTable(os.Args)
			}
		case "all":
			dochelpers.RenderTemplates()
			dochelpers.GetRogueEnvs()
			dochelpers.RenderGlobalVarsTemplate()
			dochelpers.GenerateServiceIndexMarkdowns()
		case "help":
			fallthrough
		default:
			printHelp()
		}
	} else {
		// No arguments given, print help
		printHelp()
	}
}

func printHelp() {
	fmt.Printf("Usage: %s [templates|rogue|globals|service-index|env-var-delta-table|all|help]\n", os.Args[0])
}
