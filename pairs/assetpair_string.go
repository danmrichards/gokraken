// Code generated by "stringer -type AssetPair"; DO NOT EDIT.

package pairs

import "fmt"

const _AssetPair_name = "XETHZEURXREPZUSDXXMRZEUREOSETHXICNXETHXXRPZUSDXZECZJPYXZECZUSDDASHXBTXETCXETHXREPXXBTEOSUSDXETHXXBTXETHZCADXXDGXXBTXXRPZJPYBCHUSDEOSEURGNOXBTXXBTZJPYEOSXBTGNOEURXXRPZCADXETCZUSDXXLMXXBTXXBTZEURXXMRXXBTDASHEURXETCXXBTXXBTZUSDUSDTZUSDXREPZEURXREPXETHXXLMZUSDXLTCXXBTXMLNXXBTXXBTZCADXXBTZGBPXXLMZEURXETCZEURXLTCZUSDXETHZJPYXXMRZUSDXZECXXBTGNOUSDBCHEURGNOETHXETHZGBPXICNXXBTXZECZEURBCHXBTXETHZUSDXMLNXETHDASHUSDXXRPXXBTXXRPZEURXLTCZEUR"

var _AssetPair_index = [...]uint16{0, 8, 16, 24, 30, 38, 46, 54, 62, 69, 77, 85, 91, 99, 107, 115, 123, 129, 135, 141, 149, 155, 161, 169, 177, 185, 193, 201, 208, 216, 224, 232, 240, 248, 256, 264, 272, 280, 288, 296, 304, 312, 320, 328, 336, 342, 348, 354, 362, 370, 378, 384, 392, 400, 407, 415, 423, 431}

func (i AssetPair) String() string {
	if i < 0 || i >= AssetPair(len(_AssetPair_index)-1) {
		return fmt.Sprintf("AssetPair(%d)", i)
	}
	return _AssetPair_name[_AssetPair_index[i]:_AssetPair_index[i+1]]
}
