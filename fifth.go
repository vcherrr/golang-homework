//колво различных значений
package main

import (
	"fmt"
)

func main() {
	a := []float64{1, 1, 1, 1, 2, 3, 4, 5, 9, 3, 3, 3, 3}
	c := 0
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			c += 0
		} else {
			c += 1
		}
	}
	fmt.Println(c)
}
