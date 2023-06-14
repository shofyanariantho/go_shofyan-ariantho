package main

import "fmt"

func PairSum(arr []int, target int) []int {
	index := []int{}

	for i := range arr {
		sum := 0
		for j := range arr {
			sum = arr[i] + arr[j]
			if i != j && sum == target {
				index = append(index, i, j)	
				break	
			}
		}

		if sum == target {
			break
		}
	}

	return index
}

func main(){
	fmt.Println(PairSum([]int{1,2,3,4,6}, 6))
	fmt.Println(PairSum([]int{2,5,9,11}, 11))
	fmt.Println(PairSum([]int{1,3,5,7}, 12))
	fmt.Println(PairSum([]int{1,4,6,8}, 10))
	fmt.Println(PairSum([]int{1,5,6,7}, 6))
}