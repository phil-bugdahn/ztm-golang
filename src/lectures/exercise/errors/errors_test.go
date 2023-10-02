package timeparse

import ("testing"

	)

func TestParseTime(t *testing.T) {
	_, err := ParseTime("11:03:23")
	if err != nil {
		t.Errorf("parseTime error but should be OK")
	}

	_, err = ParseTime("11:03")
	if err == nil {
		t.Errorf("parseTime should fail on single colon")
	}

	_, err = ParseTime("1A:03:k3")
	if err == nil {
		t.Errorf("parseTime should fail on non-numeric characters")
	}
}