// Copyright 2014 daimonji. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

func (g *Generator) buildOneRunFromStringParser(values []Value, typeName string) {
	// The generated code is simple enough to write as a Printf format.
	g.Printf(stringFromStringOneRun, typeName, len(values), values[0].String())
}

// Arguments to format are:
//
//	[1]: type name
//	[2]: size of element array
//	[3]: lowest defined value for type, as a string
const stringFromStringOneRun = `func FromString%[1]s(s string) (%[1]s, error) {
	for i:=0; i<%[2]d; i++ {
		if e:= %[1]s(i+%[3]s); s == e.String() {
			return e, nil
		}
	}
	return %[1]s(0), errors.New("cannot deternime %[1]s from string")
}
`

func (g *Generator) buildMultipleRunsFromStringParser(runs [][]Value, typeName string) {
	g.Printf("func FromString%[1]s(s string) (%[1]s, error) {\n", typeName)
	g.Printf("\tswitch s {\n")
	for _, values := range runs {

		// single value case
		for _, v := range values {
			g.Printf("\tcase %s.String():\n", v.originalName)
			g.Printf("\t\treturn %s, nil\n", v.originalName)
		}
	}
	g.Printf("\tdefault:\n")
	g.Printf("\t\treturn %[1]s(0), errors.New(\"cannot deternime %[1]s from string\")\n", typeName)
	g.Printf("\t}\n")
	g.Printf("}\n")
}

func (g *Generator) buildMapFromStringParser(typeName string) {
	g.Printf(stringFromStringMap, typeName)
}

// Argument to format is the type name.
const stringFromStringMap = `func FromString%[1]s(s string) (%[1]s, error) {
	for k, v := range _%[1]s_map {
		if v == s {
            return k, nil
        }
	}
	return %[1]s(0), errors.New("cannot deternime %[1]s from string")
}
`
