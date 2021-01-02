package list

import (
	"structures/test"
	"testing"
)

func TestNew(t *testing.T) {
	l := New()

	test.AssertNodeEqual(t, l.head, nil)
	test.AssertNodeEqual(t, l.tail, nil)
	test.AssertIntEqual(t, l.length, 0)
}

func TestGetLength(t *testing.T) {
	l := New()

	test.AssertIntEqual(t, l.getLength(), 0)

	l.Append(1)
	l.Append(1)
	l.Append(1)

	test.AssertIntEqual(t, l.getLength(), 3)
}
