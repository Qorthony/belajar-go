package main

import (
	"fmt"
	"strings"
)

func main() {
	var kata string
	fmt.Print("Masukkan Kata : ")
	fmt.Scan(&kata)
	reversedSplit := make([]string, len(kata))
	for i := len(kata) - 1; i >= 0; i-- {
		reversedSplit = append(reversedSplit, strings.Split(kata, "")[i])
	}
	reversedKata := strings.Join(reversedSplit, "")
	fmt.Println(reversedKata)
}
