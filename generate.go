package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

func main() {
	var flagPkg string
	var flagImport string
	var flagTypeName string
	var flagOutputDir string
	var flagOutputFile string

	flag.StringVar(&flagPkg, "package", "godash", "set the package name on generated files")
	flag.StringVar(&flagImport, "import", "", "add an import to generated files (needed for custom types)")
	flag.StringVar(&flagTypeName, "type", "string", "the type of slice to generate for")
	flag.StringVar(&flagOutputDir, "dir", "go-dash-slice", "output directory (created if needed)")
	flag.StringVar(&flagOutputFile, "out", "", "output filename")

	if flagOutputFile == "" {
		flagOutputFile = flagTypeName + ".go"
	}

	flag.Parse()

	t, err := template.ParseFiles("./templates/simple.go.tpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load template: %s", err.Error())
		os.Exit(1)
	}

	err = os.MkdirAll(flagOutputDir, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mkdir: %s", err.Error())
		os.Exit(1)
	}

	f, err := os.OpenFile(path.Join(flagOutputDir, flagOutputFile), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open: %s", err.Error())
		os.Exit(1)
	}

	typeNameCapitalised := strings.ToUpper(flagTypeName[0:1]) + flagTypeName[1:]

	err = t.Execute(f, map[string]interface{}{
		"TypeNameCapitalised": typeNameCapitalised,
		"TypeName":            flagTypeName,
		"Package":             flagPkg,
		"Import":              flagImport,
		"NewFuncName":         "New" + typeNameCapitalised + "Slice",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "template execute: %s", err.Error())
		os.Exit(1)
	}
}
