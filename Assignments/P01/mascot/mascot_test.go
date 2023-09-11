package mascot_test

import (
	"example/P01/mascot"
	"testing"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}
