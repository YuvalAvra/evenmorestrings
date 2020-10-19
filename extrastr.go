// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package evenmorestrings

import "github.com/yuvalavra/evenmorestrings/another_package"

// ReverseRunes returns its argument string reversed rune-wise left to right.
func DoubleString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return ((string)(r))
}


func CallAnotherPackTripleStr(s string) string {
    return DoubleString(another_package.TripleString(s))
}

