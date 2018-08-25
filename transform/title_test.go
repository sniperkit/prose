/*
Sniperkit-Bot
- Status: analyzed
*/

package transform

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sniperkit/snk.fork.jdkato-prose/internal/util"
)

type testCase struct {
	Input  string
	Expect string
}

func TestTitle(t *testing.T) {
	tests := make([]testCase, 0)
	cases := util.ReadDataFile(filepath.Join(testdata, "title.json"))

	util.CheckError(json.Unmarshal(cases, &tests))
	tc := NewTitleConverter(APStyle)
	for _, test := range tests {
		assert.Equal(t, test.Expect, tc.Title(test.Input))
	}
}

func BenchmarkTitle(b *testing.B) {
	tests := make([]testCase, 0)
	cases := util.ReadDataFile(filepath.Join(testdata, "title.json"))

	util.CheckError(json.Unmarshal(cases, &tests))
	tc := NewTitleConverter(APStyle)
	for n := 0; n < b.N; n++ {
		for _, test := range tests {
			_ = tc.Title(test.Input)
		}
	}
}
