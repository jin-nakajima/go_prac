// str sum method
package main

import "fmt"

type User struct {
	Name string
	Age  string
}

// method syntax
// func (v T) funcName() returnType {}
func (playUser User) StrSum() string {
	result := playUser.Name + playUser.Age
	return result
}

func main() {
	playUser := User{"nishimura", ":30"}
	fmt.Println("Hello", playUser.StrSum())
}
