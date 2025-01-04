// Copyright 2014 daimonji. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

type ConstGenerationParams struct {
	ValueSuffix         string
	NamePrefix          string
	NameSuffix          string
	ValueTransformClass _TextConvertTo
}

func (g *ConstGenerationParams) IsInit() bool {
	return len(g.NameSuffix) > 0 || len(g.NamePrefix) > 0
}

func buildRunsConstsListing(p *ConstGenerationParams, runs [][]Value, pgkName string, gPrintf func(string, ...interface{})) {
	gPrintf("\nconst (\n")

	for _, values := range runs {

		// single value case
		for _, v := range values {
			id := transformTextTo(p.ValueTransformClass, pgkName+v.originalName+p.ValueSuffix)
			gPrintf("\t%s = \"%s\"\n", p.NamePrefix+v.originalName+p.NameSuffix, id)
		}
	}

	gPrintf(")\n")
}
