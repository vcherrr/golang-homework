//фибоначчи через цикл
package main

import (
	"fmt"
)

func fib1(n int) int {
	var a, b int = 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a

}

func main() {
	res := fib1(5)
	fmt.Println("n-ный элемент: ", res)
}
