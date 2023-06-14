package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	merge := arrayA
	for i := 0; i < len(arrayB); i++ {
		name := arrayB[i]
		found := false
		
		for j := 0; j < len(arrayA); j++ {
			if arrayA[j] == name {
				found = true
				break
			}
		}

		if !found {
			merge = append(merge, name)
		}
	} 
	return merge
}

func main(){
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa"}))
}