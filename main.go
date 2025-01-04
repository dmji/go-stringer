// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func work(args, tags, types []string, dir string) {

	// For each type, generate code in the first package where the type is declared.
	// The order of packages is as follows:
	// package x
	// package x compiled for tests
	// package x_test
	//
	// Each package pass could result in a separate generated file.
	// These files must have the same package and test/not-test nature as the types
	// from which they were generated.
	//
	// Types will be excluded when generated, to avoid repetitions.

	extraConstValueTransformClass, err := _TextConvertToFromString(*extraConstValueTransformClass)
	if err != nil {
		panic(err)
	}
	nameTransform, err := _TextConvertToFromString(*nameTransformClass)
	if err != nil {
		panic(err)
	}

	pkgs := loadPackages(args, tags, *trimprefix, *linecomment, nameTransform, nil /* logf */)
	sort.Slice(pkgs, func(i, j int) bool {
		// Put x_test packages last.
		iTest := strings.HasSuffix(pkgs[i].name, "_test")
		jTest := strings.HasSuffix(pkgs[j].name, "_test")
		if iTest != jTest {
			return !iTest
		}

		return len(pkgs[i].files) < len(pkgs[j].files)
	})

	generateConstParams := &ConstGenerationParams{
		ValueTransformClass: extraConstValueTransformClass,
		ValueSuffix:         *extraConstValueSuffix,
		NamePrefix:          *extraConstNamePrefix,
		NameSuffix:          *extraConstNameSuffix,
	}

	marshallingParams := &MarshallingGenerationParams{
		QS:        *marshalQS,
		PackageQS: *marshalQsPkg,

		JSON:        *marshalJson,
		PackageJSON: *marshalJsonPkg,
	}

	for _, pkg := range pkgs {
		g := Generator{
			pkg:                 pkg,
			fromStringGenFn:     *fromStringGenFn,
			generateConstParams: generateConstParams,
			marshallingParams:   marshallingParams,
		}

		// Print the header and package clause.
		g.buildHeader(strings.Join(os.Args[1:], " "))

		// Run generate for types that can be found. Keep the rest for the remainingTypes iteration.
		var foundTypes, remainingTypes []string
		for _, typeName := range types {
			values := findValues(typeName, pkg)
			if len(values) > 0 {
				g.generate(typeName, values)
				foundTypes = append(foundTypes, typeName)
			} else {
				remainingTypes = append(remainingTypes, typeName)
			}
		}
		if len(foundTypes) == 0 {
			// This package didn't have any of the relevant types, skip writing a file.
			continue
		}
		if len(remainingTypes) > 0 && output != nil && *output != "" {
			log.Fatalf("cannot write to single file (-output=%q) when matching types are found in multiple packages", *output)
		}
		types = remainingTypes

		// Format the output.
		src := g.format()

		// Write to file.
		outputName := *output
		if outputName == "" {
			// Type names will be unique across packages since only the first
			// match is picked.
			// So there won't be collisions between a package compiled for tests
			// and the separate package of tests (package foo_test).
			outputName = filepath.Join(dir, baseName(pkg, foundTypes[0]))
		}
		err := os.WriteFile(outputName, src, 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	}

	if len(types) > 0 {
		log.Fatalf("no values defined for types: %s", strings.Join(types, ","))
	}
}

// baseName that will put the generated code together with pkg.
func baseName(pkg *Package, typename string) string {
	suffix := "string.go"
	if pkg.hasTestFiles {
		suffix = "string_test.go"
	}

	transform, err := _TextConvertToFromString(*outputTransformClass)
	if err != nil {
		panic(err)
	}
	name := transformTextTo(transform, typename)

	return fmt.Sprintf("%s_%s", name, suffix)
}
