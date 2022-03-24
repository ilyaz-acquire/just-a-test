package main

import "fmt"

func Foo(x int, y int) int {
	if x == 42 {
		return 19
	}
	return x + y
}

func Bar(x int, y int) int {
	return x * y
}

func main() {
	fmt.Println("vim-go")
}
