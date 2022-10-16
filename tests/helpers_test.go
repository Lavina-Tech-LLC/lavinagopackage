package tests

import (
	"testing"

	lvn "github.com/Lavina-Tech-LLC/lavinagopackage/v2"
)

func TestTernary(t *testing.T) {

	k := lvn.Ternary(true, 0, 1)
	if k != 0 {
		t.Errorf("Res was incorrect, got: %d, want: %d.", k, 0)
	}
	k = lvn.Ternary(0, 0, 1)
	if k != 1 {
		t.Errorf("Res was incorrect, got: %d, want: %d.", k, 1)
	}

	k = lvn.Ternary("", 0, 1)
	if k != 1 {
		t.Errorf("Res was incorrect, got: %d, want: %d.", k, 1)
	}

}
