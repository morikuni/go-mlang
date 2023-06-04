package mlangfailure

import (
	"testing"

	"github.com/morikuni/failure/v2"

	"github.com/morikuni/go-mlang"
)

func TestMessageOf(t *testing.T) {
	err := failure.Unexpected("unexpected", mlang.Dict[string]{1: "one"})
	m := MessageOf(err)
	if m == nil {
		t.Fatal("m is nil")
	}
	if m.MustGet(1) != "one" {
		t.Errorf("got %v, want %v", m.MustGet(1), "one")
	}
}
