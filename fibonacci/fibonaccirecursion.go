package main

import "fmt"


func fibonaccirecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonaccirecursion(n - 1) + fibonaccirecursion(n - 2)
}

func main() {
	for i := 0; i <= 6; i++ {
		fmt.Println(fibonaccirecursion(i))
	}
}