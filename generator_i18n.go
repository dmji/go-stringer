// Copyright 2014 daimonji. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

func (g *Generator) buildRunsConsti18nListing(runs [][]Value) {
	/* 	g.Printf("func %[1]sFromString(s string) (%[1]s, error) {\n", typeName)
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
	*/
	g.Printf("\nconst (\n")

	for _, values := range runs {

		// single value case
		for _, v := range values {
			id := transformTextTo(g.i18nTransform, g.pkg.name+v.originalName+g.i18nIdSuffix)
			g.Printf("\t%s = \"%s\"\n", g.i18nIdNamePrefix+v.originalName+"_i18n_ID", id)
		}
	}

	g.Printf(")\n")
}
