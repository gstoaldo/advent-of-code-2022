package utils

import "testing"

func Assert(t *testing.T, want interface{}, got interface{}) {
	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
