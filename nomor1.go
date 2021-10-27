package main

import (
	"fmt"
	"math"
)

func primeX(number int) (prime int) {

	var tampung int
	i := 2

	for i >= 2 {
		cek := true
		akar := int(math.Sqrt(float64(i)))
		j := 2
		for j <= akar {
			if i%j == 0 {
				cek = false
				j = i
			}
			j++
		}

		if cek {
			tampung++
		}

		if tampung == number {
			return i
		}

		i++
	}
	return i
}

func main() {

	fmt.Println(primeX(1)) // 2

	fmt.Println(primeX(5)) // 11

	fmt.Println(primeX(8)) // 19

	fmt.Println(primeX(9)) // 23

	fmt.Println(primeX(10)) // 29

}