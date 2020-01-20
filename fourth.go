//сортировка методом пузырька
package main

import (
	"fmt"
)

func bubble_sort(x []int) {
	for i := 0; i < len(x); i++ {
		for j := len(x) - 1; j > i; j-- {
			if x[j-1] > x[j] {
				swap(x, j-1, j)
			}
		}
	}
}

func swap(x []int, i, j int) {
	y := x[i]
	x[i] = x[j]
	x[j] = y
}

func main() {
	x := []int{24, 0, -5, -2, 8, 5}
	bubble_sort(x)
	fmt.Println(x)
}
