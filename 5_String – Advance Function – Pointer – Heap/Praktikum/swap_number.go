package main

import "fmt"

func swap(a, b *int) {
	tampung := *a
	*a = *b
	*b = tampung
}

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}