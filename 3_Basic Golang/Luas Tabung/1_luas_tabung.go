package main

import "fmt"

func main(){
	// input 
	// T := 20.0
	// r := 4.0

	// luas permukaan tabung
	// output := 2 * 3.14 * r * (r + T)
	// fmt.Println("Luas Permukaan Tabung :", output)

	var T float64
    var r float64

    fmt.Print("Masukkan tinggi tabung: ")
    fmt.Scanf("%f", &T)

    fmt.Print("Masukkan jari-jari tabung: ")
    fmt.Scanf("%f", &r)

    output := 2 * 3.14 * r * (r + T)
    fmt.Println("Luas Permukaan Tabung:", output)

}