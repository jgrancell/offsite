package main

import "testing"

func TestRun(t *testing.T) {
	out := Run(nil)
	if out != 0 {
		t.Errorf("error code: %d", out)
	}
}
