package main

import "fmt"

func Foo(x int, y int) int {
	if x == 42 {
		return 19
	}
	return x + y
}

func bar(x int, y int) int {
	return x * y
}

func Buzz(x int) string {
	switch x {
	case 1:
		return "boom"
	case 2:
		return "two"
	default:
		return "hm"
	}
}

func main() {
	fmt.Println("vim-go")
}
