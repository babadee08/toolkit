package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools

	s := testTools.RandomString(64)
	if len(s) != 64 {
		t.Error("Wrong length random string returned")
	}
}
