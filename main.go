package main

import (
	"structures/list"
)

func main()  {
	l := list.New()

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)

/*	l.Print()
*/	l.Reverse()
	l.Print()
}