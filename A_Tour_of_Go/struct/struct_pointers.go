package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v // 構造体のアドレスのはずが、&{1, 2}を返す
	fmt.Println(p)
	fmt.Println(*p)
	p.X = 1e9
	fmt.Println(v)
}
