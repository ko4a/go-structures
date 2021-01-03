package test

import (
	"structures/list"
	"testing"
)

func TestNew(t *testing.T) {
	l := list.New()

	AssertIntEqual(t, l.GetLength(), 0)
	AssertNodeEqual(t, l.GetHead(), l.GetTail())
}

func TestGetLength(t *testing.T) {
	l := list.New()

	AssertIntEqual(t, l.GetLength(), 0)

	l.Append(1)
	l.Append(1)
	l.Append(1)

	AssertIntEqual(t, l.GetLength(), 3)
}

func TestAppend(t *testing.T) {
	l := list.New()

	l.Append(1)

	AssertIntEqual(t, l.GetLength(), 1)
	AssertNodeEqual(t, l.GetHead(), l.GetTail())
	AssertIntEqual(t, l.GetTail().GetData().(int), 1)

	l.Append(2)

	AssertIntEqual(t, l.GetLength(), 2)
	AssertNodeNotEqual(t, l.GetHead(), l.GetTail())
	AssertIntEqual(t, l.GetTail().GetData().(int), 2)
}

func TestPrepend(t *testing.T) {
	l := list.New()

	l.Prepend(1)

	AssertIntEqual(t, l.GetLength(), 1)
	AssertNodeEqual(t, l.GetHead(), l.GetTail())
	AssertIntEqual(t, l.GetTail().GetData().(int), 1)

	l.Prepend(2)

	AssertIntEqual(t, l.GetLength(), 2)
	AssertNodeNotEqual(t, l.GetHead(), l.GetTail())
	AssertIntEqual(t, l.GetHead().GetData().(int), 2)
}
