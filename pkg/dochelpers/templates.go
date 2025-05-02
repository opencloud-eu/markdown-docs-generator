package dochelpers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

var targets = map[string]string{
	"templates/go/markdown-generator.go.tmpl":                  "output/markdown/markdown-generator.go",
	"templates/go/example-config-generator.go.tmpl":            "output/exampleconfig/example-config-generator.go",
	"templates/go/environment-variable-docs-generator.go.tmpl": "output/env/environment-variable-docs-generator.go",
	"templates/go/envar-delta-table.go.tmpl":                   "output/env/envvar-delta-table.go",
}

// RenderTemplates does something with templates
func RenderTemplates() {
	baseDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Getting relevant packages")
	paths, err := filepath.Glob("tmp/services/*/pkg/config/defaults/defaultconfig.go")
	if err != nil {
		log.Fatal(err)
	}
	replacer := strings.NewReplacer(
		"tmp/", "github.com/opencloud-eu/opencloud/",
		"/defaultconfig.go", "",
	)
	for i := range paths {
		paths[i] = replacer.Replace(paths[i])
	}

	for tmpl, output := range targets {
		generateIntermediateCode(baseDir+"/"+tmpl, output, paths)
		runIntermediateCode(output)
	}
	fmt.Println("Cleaning up")
	err = os.RemoveAll("output")
	if err != nil {
		fmt.Println(err)
	}
}

func generateIntermediateCode(templatePath string, intermediateCodePath string, paths []string) {
	content, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generating intermediate go code for " + intermediateCodePath + " using template " + templatePath)
	tpl := template.Must(template.New("").Parse(string(content)))
	err = os.MkdirAll(path.Dir(intermediateCodePath), 0700)
	if err != nil {
		log.Fatal(err)
	}
	runner, err := os.Create(intermediateCodePath)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(runner, paths)
	if err != nil {
		log.Fatal(err)
	}
}

func runIntermediateCode(intermediateCodePath string) {
	defaultConfigPath := "/etc/opencloud"
	defaultDataPath := "/var/lib/opencloud"
	os.Setenv("OCIS_BASE_DATA_PATH", defaultDataPath)
	os.Setenv("OCIS_CONFIG_DIR", defaultConfigPath)
	fmt.Println("Running go mod tidy")
	out, err := exec.Command("go", "mod", "tidy").CombinedOutput()
	if err != nil {
		log.Fatal(string(out), err)
	}
	fmt.Println("Running intermediate go code for " + intermediateCodePath)
	out, err = exec.Command("go", "run", intermediateCodePath).CombinedOutput()
	if err != nil {
		log.Fatal(string(out), err)
	}
	fmt.Println(string(out))
}
