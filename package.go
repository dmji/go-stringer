// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"log"
	"strings"

	"golang.org/x/tools/go/packages"
)

type Package struct {
	name         string
	defs         map[*ast.Ident]types.Object
	files        []*File
	hasTestFiles bool
}

// loadPackages analyzes the single package constructed from the patterns and tags.
// loadPackages exits if there is an error.
//
// Returns all variants (such as tests) of the package.
//
// logf is a test logging hook. It can be nil when not testing.
func loadPackages(
	patterns, tags []string,
	trimPrefix string, lineComment bool,
	logf func(format string, args ...interface{}),
) []*Package {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedFiles,
		// Tests are included, let the caller decide how to fold them in.
		Tests:      true,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
		Logf:       logf,
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) == 0 {
		log.Fatalf("error: no packages matching %v", strings.Join(patterns, " "))
	}

	out := make([]*Package, len(pkgs))
	for i, pkg := range pkgs {
		p := &Package{
			name:  pkg.Name,
			defs:  pkg.TypesInfo.Defs,
			files: make([]*File, len(pkg.Syntax)),
		}

		for j, file := range pkg.Syntax {
			p.files[j] = &File{
				file: file,
				pkg:  p,

				trimPrefix:  trimPrefix,
				lineComment: lineComment,
			}
		}

		// Keep track of test files, since we might want to generated
		// code that ends up in that kind of package.
		// Can be replaced once https://go.dev/issue/38445 lands.
		for _, f := range pkg.GoFiles {
			if strings.HasSuffix(f, "_test.go") {
				p.hasTestFiles = true
				break
			}
		}

		out[i] = p
	}
	return out
}

func findValues(typeName string, pkg *Package) []Value {
	values := make([]Value, 0, 100)
	for _, file := range pkg.files {
		// Set the state for this run of the walker.
		file.typeName = typeName
		file.values = nil
		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
			values = append(values, file.values...)
		}
	}
	return values
}
