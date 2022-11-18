package main

import (
	"fmt"
)

type Calclate interface {
	Add() float64
	Sub() float64
}

type Input struct {
	value1, value2 float64
}

func (i *Input) Add() float64 {
	result := i.value1 + i.value2
	return result
}

func (i *Input) Sub() float64 {
	result := i.value1 - i.value2
	return result
}

func main() {
	var c Calclate
	f := Input{3, 4}

	c = &f
	fmt.Println(c.Add())
	fmt.Println(c.Sub())
}
