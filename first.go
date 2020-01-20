//умножение через сложение
package main

import (
	"fmt"
)

func multiply(a, b int) int {
	var c int
	if b < a {
		return multiply(b, a)
	} else if b < 0 {
		b = -b
		a = -a
	}
	for i := 0; i < b; i++ {
		c += a
	}
	return c
}

func main() {
	res := multiply(46, 2)
	fmt.Println(res)
}
