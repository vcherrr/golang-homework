//фибоначчи через рекурсию
package main

import "fmt"

func fib2(n int) int {
	if n < 2 {
		return 1
	}
	return fib2(n-2) + fib2(n-1)
}

func main() {
	res := fib2(6)
	fmt.Println("n-ный элемент: ", res)
}
