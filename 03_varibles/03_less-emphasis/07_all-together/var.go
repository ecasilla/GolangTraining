package main

import (
	"fmt"
)
var a = "this is stored a the package level scope " //package level
var b, c string = "hello b package","hello c package" //package level
var d string //package level

func main() {
	d = "hello d package" // still package level because the declaration happened
	// outside the function
	var e = 2 //fn level scoped variable
	f := 43
	g := "hello G"
	h, i := "hello h","hello I"


	j, k, l, m := 44.7, true, false, 'm' // single quotes
	n := "n"                             // double quotes
	o := `o`                             // back ticks

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)

}
