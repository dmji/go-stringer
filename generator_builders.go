// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

// buildOneRun generates the variables and String method for a single run of contiguous values.
func (g *Generator) buildOneRun(runs [][]Value, typeName string) {
	values := runs[0]
	g.Printf("\n")
	g.declareIndexAndNameVar(values, typeName)

	// Stringer implementation
	g.buildOneRunStringer(values, typeName)

	// Type from string implementation
	if g.genFromStringFn {
		g.buildOneRunFromStringParser(values, typeName)
	}
}

// buildMultipleRuns generates the variables and String method for multiple runs of contiguous values.
// For this pattern, a single Printf format won't do.
func (g *Generator) buildMultipleRuns(runs [][]Value, typeName string) {
	g.Printf("\n")
	g.declareIndexAndNameVars(runs, typeName)

	// Stringer implementation
	g.buildMultipleRunsStringer(runs, typeName)

	// Type from string implementation
	if g.genFromStringFn {
		g.buildMultipleRunsFromStringParser(runs, typeName)
	}
}

// buildMap handles the case where the space is so sparse a map is a reasonable fallback.
// It's a rare situation but has simple code.
func (g *Generator) buildMap(runs [][]Value, typeName string) {
	g.Printf("\n")
	g.declareNameVars(runs, typeName, "")
	g.Printf("\nvar _%s_map = map[%s]string{\n", typeName, typeName)
	n := 0
	for _, values := range runs {
		for _, value := range values {
			g.Printf("\t%s: _%s_name[%d:%d],\n", &value, typeName, n, n+len(value.name))
			n += len(value.name)
		}
	}
	g.Printf("}\n\n")

	// Stringer implementation
	g.buildMapStringer(typeName)

	// Type from string implementation
	if g.genFromStringFn {
		g.buildMapFromStringParser(typeName)
	}
}

// declareNameVars declares the concatenated names string representing all the values in the runs.
func (g *Generator) declareNameVars(runs [][]Value, typeName string, suffix string) {
	g.Printf("const _%s_name%s = \"", typeName, suffix)
	for _, run := range runs {
		for i := range run {
			g.Printf("%s", run[i].name)
		}
	}
	g.Printf("\"\n")
}
