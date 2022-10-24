package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]*Vertex

func main() {
	m = make(map[string]*Vertex)
	m["Bell Labs"] = &Vertex{
		11.11, 33.33,
	}
	fmt.Println(m["aaaa"]) // return nil
}
