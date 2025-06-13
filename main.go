package main

import (
	"fmt"
)

func main() {
	digits := []int{1, 2, 3}
	fmt.Println("Digits:", digits)
	fmt.Println("Digits plus one:", digitsPlusOne(digits))
}

func digitsPlusOne(d []int) []int {
	for i := len(d) - 1; i >= 0; i-- {

		if d[i] < 9 {
			d[i]++
			return d
		}
		d[i] = 0
	}

	return append([]int{1}, d...)
}
