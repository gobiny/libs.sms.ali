package utils

import "testing"

func TestGetUid(t *testing.T) {
	a := GetUid()
	b := GetUid()
	if a == b {
		t.Error(a)
	}
}
