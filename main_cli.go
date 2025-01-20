// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	typeNames            = flag.String("type", "", "comma-separated list of type names\nmust be set")
	output               = flag.String("output", "", "output file name\ndefault srcdir/<type>_string.go")
	outputTransformClass = flag.String("outputtransform", "lower", "text-style name of default output name\navailable options: "+availableValuesForFlagsDefault())

	trimprefix         = flag.String("trimprefix", "", "trim the `prefix` from the generated constant names (@me if same as type name)")
	linecomment        = flag.Bool("linecomment", false, "use line comment text as printed text when present")
	nameTransformClass = flag.String("nametransform", "none", "text-style name of String() data\navailable options: "+availableValuesForFlagsDefault())

	fromStringGenFn = flag.Bool("fromstringgenfn", false, "use to generate ${TypeName}FromString() function")

	buildTags  = flag.String("tags", "", "comma-separated list of build tags to apply")
	targetFile = flag.String("target", ".", "path to file or folder with desired types")
)

// parameters of extra consts generation (optional)
var (
	extraConstValueSuffix         = flag.String("extraconstsvaluesuffix", "", "suffix to generate consts with id for i18n")
	extraConstNamePrefix          = flag.String("extraconstsnameprefix", "", "prefix to generate consts names with id for i18n")
	extraConstNameSuffix          = flag.String("extraconstsnamesuffix", "", "suffix to generate consts names with id for i18n")
	extraConstValueTransformClass = flag.String("extraconstsvaluetransform", "pascal_case", "text-style name of i18n ID\navailable options: "+availableValuesForFlagsDefault())
)

// parameters of extra marshaler generation (optional)
var (
	marshalJsonPkg = flag.String("marshaljsonpkg", "encoding/json", "use to specify import path of package json")
	marshalJson    = flag.Bool("marshaljson", false, "use to generate implementation for json.Marshaler and json.Unmarshaler")

	marshalQsPkg = flag.String("marshalqspkg", "github.com/pasztorpisti/qs", "use to specify import path of package qs")
	marshalQS    = flag.Bool("marshalqs", false, "use to generate implementation for qs.MarshalQS and qs.UnmarshalQS")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of go-stringer (stringer fork):\n")
	fmt.Fprintf(os.Stderr, "\tgo-stringer [flags] -type T [directory]\n")
	fmt.Fprintf(os.Stderr, "\tgo-stringer [flags] -type T files... # Must be a single package\n")
	fmt.Fprintf(os.Stderr, "For more information, see:\n")
	fmt.Fprintf(os.Stderr, "\thttps://pkg.go.dev/golang.org/x/tools/cmd/stringer\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("go-stringer: ")
	flag.Usage = Usage
	flag.Parse()
	_ = targetFile
	if len(*typeNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	types := strings.Split(*typeNames, ",")
	var tags []string
	if len(*buildTags) > 0 {
		tags = strings.Split(*buildTags, ",")
	}

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{*targetFile}
	}

	// Parse the package once.
	var dir string
	// TODO(suzmue): accept other patterns for packages (directories, list of files, import paths, etc).
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		if len(tags) != 0 {
			log.Fatal("-tags option applies only to directories, not when files are specified")
		}
		dir = filepath.Dir(args[0])
	}

	work(args, tags, types, dir)
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}
