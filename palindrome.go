package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Kata : ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	fmt.Println(isPalindrome(input))
}

func isPalindrome(input string) bool {
	inputSplit := strings.Split(input, " ")

	//Jika Sebuah Kalimat
	if len(inputSplit) > 1 {
		firstWord := inputSplit[0]
		lastWord := reverseKata(inputSplit[len(inputSplit)-1])

		if firstWord == lastWord {
			return true
		}
		return false
	}

	//Jika Sebuah Kata
	return input == reverseKata(input)

}

func reverseKata(kata string) string {
	reversedSplit := make([]string, len(kata))
	for i := len(kata) - 1; i >= 0; i-- {
		reversedSplit = append(reversedSplit, strings.Split(kata, "")[i])
	}
	reversedKata := strings.Join(reversedSplit, "")
	return reversedKata
}
