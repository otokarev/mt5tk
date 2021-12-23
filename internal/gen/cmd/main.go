package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

type data struct {
	Type      string
	Name      string
	ShortName string
	IdName    string
	GetPath   string
	CmdNames  []string
}

func main() {
	var d data
	var outName string
	var cmdNames string
	var tmplFileName string
	flag.StringVar(&d.Type, "type", "", "class name")
	flag.StringVar(&d.Name, "name", "", "package name")
	flag.StringVar(&d.IdName, "id-name", "", "get request id name")
	flag.StringVar(&d.ShortName, "short-name", "", "short name")
	flag.StringVar(&d.GetPath, "get-path", "", "get path, like: /api/group/get")
	flag.StringVar(&cmdNames, "cmd-names", "", "command names")
	flag.StringVar(&outName, "out", "", "output file name")
	flag.StringVar(&tmplFileName, "template", "", "template file")
	flag.Parse()

	for _, v := range strings.Split(cmdNames, ",") {
		if v != "" {
			d.CmdNames = append(d.CmdNames, strings.Title(v))
		}
	}

	if d.IdName == "" {
		d.IdName = d.Name
	}

	if d.GetPath == "" {
		d.GetPath = "/api/" + d.Name + "/get"
	}

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
