package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	if number < 2 {
		return false
	}

	sqrt := int(math.Sqrt(float64(number)))

	for i := 2; i <= sqrt; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
