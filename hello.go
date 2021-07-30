package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan angka :")
	fmt.Scan(&angka)
	
	if angka%3==0 && angka%5==0 {
		fmt.Println("Hello World")	
	} else if angka%3==0 {
		fmt.Println("Hello")
	} else if angka%5==0 {
		fmt.Println("World")
	} else {
		fmt.Println("Tidak habis dibagi 3 dan 5")
	}
	
}