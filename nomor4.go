package main

import "fmt"

func sumSubArray(arr []int, lindex int, rindex int) int {
	var tampung int
	for i := lindex; i <= rindex; i++ {
		tampung += arr[i]
	}
	return tampung
}

func MaxSequence(arr []int) int {

	var subArray int
	var maksimal int = arr[0]
	
	for i := range arr {
		for j := len(arr) - 1; j > i; j-- {
			subArray = sumSubArray(arr, i, j)
			if maksimal < subArray {
				maksimal = subArray
			}
		}
	}

	return maksimal
}

func main() {

	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

	fmt.Println(MaxSequence([]int{-2, -5, -3, -2, -3, -1, -6, -6}))

}