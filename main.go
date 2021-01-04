package main

import (
	"fmt"
	"structures/list"
)

func main() {
	l := list.New()
	lTest := list.New()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	lTest.Prepend(1)
	lTest.Prepend(2)
	lTest.Prepend(3)

	l.Round(func(data interface{}) interface{} {
		return data.(int) * 2
	})

	l.Round(func(data interface{}) interface{} {
		fmt.Printf("%v", data)
		return data
	})

}
