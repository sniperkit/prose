/*
Sniperkit-Bot
- Status: analyzed
*/

/*
Package model contains internals used by prose/tag.
*/
package model

import (
	"bytes"
	"encoding/gob"

	"github.com/sniperkit/snk.fork.jdkato-prose/internal/util"
)

// GetAsset returns the named Asset.
func GetAsset(name string) *gob.Decoder {
	b, err := Asset("internal/model/" + name)
	util.CheckError(err)
	return gob.NewDecoder(bytes.NewReader(b))
}
