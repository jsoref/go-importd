// Copyright 2015 The Go Authors.  All rights reserved.
// Copyright 2016 Christian Höltje.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	addr = flag.String("addr", ":http", "serve http on `address`")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: go-import-redirector\n")
	fmt.Fprintf(os.Stderr, "options:\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("go-import-redirector: ")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 0 {
		flag.Usage()
	}
	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var tmpl = template.Must(template.New("main").Parse(`<!DOCTYPE html>
<html lang="en">
<meta charset="utf-8">
<meta name="go-import" content="{{.ImportPrefix}} git {{.VcsUrl}}">
<meta name="go-source" content="{{.ImportPrefix}} {{.VcsUrl}} {{.VcsUrl}}/tree/master{/dir} {{.VcsUrl}}/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url={{.VcsUrl}}">
<p>
The source is at <a href="{{.VcsUrl}}">{{.VcsUrl}}</a>; you should be redirected shortly.
`))

type templateData struct {
	ImportPrefix string
	VcsUrl       string
}

func redirect(resp http.ResponseWriter, req *http.Request) {
	data := &templateData{}

	s := strings.SplitN(strings.Trim(req.URL.Path, "/"), "/", 2)
	repoName := s[0]
	data.ImportPrefix = "docwhat.org/" + repoName
	data.VcsUrl = "https://github.com/docwhat/" + repoName

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		http.Error(resp, err.Error(), 500)
		return
	}

	resp.Write(buf.Bytes())
}
