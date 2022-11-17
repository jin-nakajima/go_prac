// str sum method
package main

import "fmt"

type User struct {
	Name string
	Age  string
}

type Test struct {
	hoge string
}

// method syntax
// func (v T) funcName() returnType {}
func (playUser User) StrSum() string {
	result := playUser.Name + playUser.Age
	return result
}

func (playUser Test) StrSum() string {
	result := playUser.hoge
	return result
}

func main() {
	// playUser := User{"nishimura", ":30"}
	playUser := Test{"fuga"}
	// ↑playUser.StrSum undefined
	// (type Test has no field or method StrSum)
	fmt.Println("Hello", playUser.StrSum())
}
