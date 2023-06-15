package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	arrA := strings.Split(a, "")
	arrB := strings.Split(b, "")
	arrResult := []string{}
	for i := range arrA {
		for j := range arrB {
			if arrA[i] == arrB[j] {
				arrResult = append(arrResult, arrA[i])
				break
			}
		}
	}
	
	result := strings.Join(arrResult, "")	
	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
}