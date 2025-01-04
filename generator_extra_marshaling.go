// Copyright 2014 daimonji. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

type MarshallingGenerationParams struct {
	QS        bool
	PackageQS string

	JSON        bool
	PackageJSON string
}

func (g *MarshallingGenerationParams) IsInit() bool {
	return g.QS || g.JSON
}

func buildRunsMarshalling(p *MarshallingGenerationParams, typeName string, gPrintf func(string, ...interface{})) {
	if p.QS {
		buildRunsMarshallingQS(typeName, gPrintf)
	}
	if p.JSON {
		buildRunsMarshallingJSON(typeName, gPrintf)
	}
}

func buildRunsMarshallingQS(typeName string, gPrintf func(string, ...interface{})) {
	gPrintf(`
	func (e %[1]s) MarshalQS(opts *qs.MarshalOptions) ([]string, error) {
		return []string{e.String()}, nil
	}

	func (e *%[1]s) UnmarshalQS(a []string, opts *qs.UnmarshalOptions) error {
		s, err := opts.SliceToString(a)
		if err != nil {
			return err
		}
		*e, err = %[1]sFromString(s)
		return err
	}`, typeName)
}

func buildRunsMarshallingJSON(typeName string, gPrintf func(string, ...interface{})) {
	gPrintf(`
	func (e %[1]s) MarshalJSON() ([]byte, error) {
		return json.Marshal(e.String())
	}

	func (e *%[1]s) UnmarshalJSON(data []byte) error {
		var s string
		err := json.Unmarshal(data, &s);
		if err != nil {
			return err
		}

		*e, err = %[1]sFromString(s)
		return err
	}`, typeName)
}
