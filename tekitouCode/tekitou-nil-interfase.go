package main

import "fmt"

type Interface interface {
	Methods()
}

type Type struct {
	String string
}

func (t *Type) Methods() {
	fmt.Println(t.String)
}

func main() {
	var i Interface
	var t *Type
	i = t
	describe(i)
}

func describe(i Interface) {
	fmt.Printf("(%v, %T)\n", i, i)
}
