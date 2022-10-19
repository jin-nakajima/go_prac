package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	rtn := make([][]uint8, dy)
	for y := range rtn {
		rtn[y] = make([]uint8, dx)
	}

	for y, s := range rtn {
		for x := range s {
			s[x] = uint8(x + y)
		}
	}
	return rtn
}

func main() {
	pic.Show(Pic)
}
