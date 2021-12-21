package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"text/template"
)

type data struct {
	Type string
	Name string
}

func main() {
	var d data
	var outName string
	var tmplFileName string
	flag.StringVar(&d.Type, "type", "", "class name")
	flag.StringVar(&d.Name, "name", "", "package name")
	flag.StringVar(&outName, "out", "", "output file name")
	flag.StringVar(&tmplFileName, "template", "", "template file")
	flag.Parse()

	t := template.Must(template.New(path.Base(tmplFileName)).ParseFiles(tmplFileName))
	outFile, err := os.Create(outName)
	if err != nil {
		exitErrorf("failed to open out file, %v.", err)
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			exitErrorf("failed to successfully write %q file, %v", outName, err)
		}
	}()
	err = t.Execute(outFile, d)
	if err != nil {
		exitErrorf("cannot process template %v", err)
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
