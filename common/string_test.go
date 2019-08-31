package common

import (
	"testing"
)

func TestGetContinueCharLen(t *testing.T) {
	base := "--++++++++++--"
	start := 2
	if res := GetContinueCharLen(base, start); res != 10 {
		t.Fail()
	}
}
