package tests

import (
	"fmt"
	"testing"
)

type (
	testsRes[T any] struct {
		Out  T
		Want T
		Test string
	}
)

func check[T any](res []testsRes[T], t *testing.T) {
	for _, r := range res {
		if fmt.Sprint(r.Out) != fmt.Sprint(r.Want) {
			t.Errorf("%s: %v, want: %v.", r.Test, r.Out, r.Want)
		}
	}
}
