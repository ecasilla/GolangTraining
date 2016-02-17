package main

import (
	"fmt"
)

func main() {
	var x = 42
	fmt.Println(x)
	foo()
}

func foo()  {
	fmt.Println(x)
}
