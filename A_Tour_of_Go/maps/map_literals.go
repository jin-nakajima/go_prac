package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		11.11, 22.22,
	},
	"Google": Vertex{
		33.33, 44.44,
	},
}

func main() {
	fmt.Println(m)
}
