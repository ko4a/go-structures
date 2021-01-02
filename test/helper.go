package test

import (
	"structures/list"
	"testing"
)

func AssertNodeEqual(t *testing.T, a *list.Node, b *list.Node) {
	if a != b {
		t.Fatalf("%v != %v", a, b)
	}
}

func AssertIntEqual(t *testing.T, a int, b int) {
	if a != b {
		t.Fatalf("expected %d, but got %d", a, b)
	}
}
