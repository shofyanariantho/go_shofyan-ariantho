package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {
	arrAngka := strings.Split(angka, "")
	var arrCheck []int

	for i := range arrAngka {
		unique := true

		for j := range arrAngka {
			if i != j && arrAngka[i] == arrAngka[j] {
				unique = false
				break
			}
		}
		
		if unique {
			number, _ := strconv.Atoi(arrAngka[i])
			arrCheck = append(arrCheck, number)
		}
	}

	return arrCheck
}

func main(){
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}