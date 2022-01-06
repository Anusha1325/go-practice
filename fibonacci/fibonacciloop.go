package main

import "fmt"

func fibonacciloop(n int) int {
	var a, b, c int = 0, 1, 0
	for i := 1; i <= n; i++ {
		c = a + b
		fmt.Println(a)
		a = b
		b = c
	}
	return a
}
func main() {
	fibonacciloop(8)
}
