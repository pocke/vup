package main

import "testing"

func TestVersionUp(t *testing.T) {
	type Expectation struct {
		prevVersion string
		target      string
		expect      string
	}

	ex := []Expectation{
		{"0.0.1", "patch", "0.0.2"},
		{"0.1.0", "patch", "0.1.1"},
		{"0.1.0", "minor", "0.2.0"},
		{"0.1.1", "minor", "0.2.0"},
		{"0.1.1", "major", "1.0.0"},
	}
	for _, e := range ex {
		got, err := versionUp(e.prevVersion, e.target)
		if err != nil {
			t.Error(err)
		}
		if got != e.expect {
			t.Errorf("got: %s, expect: %s", got, e.expect)
		}
	}
}
