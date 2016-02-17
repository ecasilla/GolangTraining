package main

import("fmt")

func main() {
 // this is only possible inside a fn
	message := "hello"
	a,b,c := 1,false,3
	d := 4
	e := true
	fmt.Println(message,a,b,c,d,e)
}