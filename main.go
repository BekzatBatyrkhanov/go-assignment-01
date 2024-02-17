package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 13, 4, 10, 3, 1}

	sortSlice(arr)

	printSlice(arr)

	incrementOdd(arr)

	printSlice(arr)

	reverseSlice(arr)

	printSlice(arr)

}

func sortSlice(slice []int) {

	n := len(slice)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {

				temp := slice[j+1]
				slice[j+1] = slice[j]
				slice[j] = temp

			}
		}
	}
}

func incrementOdd(slice []int) {
	for i := 0; i < len(slice); i += 2 {
		slice[i] += 1
	}
}

func printSlice(slice []int) {
	fmt.Print(slice)
	fmt.Println()
}

func reverseSlice(slice []int) {
	n := len(slice)

	for i := 0; i < len(slice)/2; i++ {
		temp := slice[i]
		slice[i] = slice[n-1-i]
		slice[n-1-i] = temp
	}
}

func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(a []int) {
		dst(a)
		for _, fn := range src {
			fn(a)
		}
	}
}
