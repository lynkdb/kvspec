package kvspec

import (
	"strings"
	"testing"
)

func Test_TableNameReg(t *testing.T) {
	for _, v := range []string{
		"aaaa-aaaa",
		"aaaa_aaaa",
		"aaaa/aaaa",
		"aaaa:aaaa",
		"aaaa.aaaa",
		strings.Repeat("a", 64),
	} {
		if !TableNameReg.MatchString(v) {
			t.Fatal("fail test")
		}
	}

	for _, v := range []string{
		"00",
		"a0",
		strings.Repeat("a", 65),
	} {
		if TableNameReg.MatchString(v) {
			t.Fatal("fail test")
		}
	}
}
