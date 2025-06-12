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
	k := 0
	var copyD = d

	for i := 1; len(copyD) > 0; {
		k++
		copyD = copyD[i:]
	}

	powerTen := 1

	for i := 0; i < k-1; i++ {
		powerTen *= 10
	}

	number := 0

	for i := 0; i < len(d); i++ {
		number += d[i] * powerTen
		powerTen /= 10
	}

	number++
	nNumber := number
	newK := 0

	for nNumber > 0 {
		newK++
		nNumber /= 10
	}

	powerTen = 1

	for i := 0; i < newK-1; i++ {
		powerTen *= 10
	}

	var result []int

	for i := 0; i < newK; i++ {
		result = append(result, number/powerTen)
		number -= powerTen * (number / powerTen)
		powerTen /= 10
	}

	return result
}
