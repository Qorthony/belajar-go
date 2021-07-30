package main

import (
	"fmt"
	"strings"
)

func main() {
	var email string

	fmt.Print("Email :")
	fmt.Scan(&email)

	hasil := validasi_email(email)

	if hasil {
		fmt.Println("Email valid")
	} else {
		fmt.Println("Email tidak valid")
	}

}

func validasi_email(email string) bool {
	/*
		Rule :
		1. harus terdapat '@'
		2. harus ada titik setelah '@'
		3. maksimal panjang karakter sebelum '@' adalah 20
		4. domain yang diperbolehkan hanya '.co.id' dan '.id'
	*/
	var result bool
	if exist := existAtSign(email); exist == false {
		result = false
	} else if dotAfterAtSign(email) == false {
		result = false
	} else if maxCharBeforAtSign(email) == false {
		result = false
	} else if restrictedDomain(email) == false {
		result = false
	} else {
		result = true
	}
	return result
}

func existAtSign(email string) bool {
	if strings.Contains(email, "@") {
		return true
	}
	return false
}

func indexOfAtSign(email string) int {
	return strings.Index(email, "@")
}

func dotAfterAtSign(email string) bool {
	var containAtSign = existAtSign(email)
	if containAtSign {
		var splitEmail = strings.Split(email, "")

		for i := indexOfAtSign(email); i < len(splitEmail)-1; i++ {
			if splitEmail[i] == "." {
				return true
			}
		}
	}
	return false
}

func maxCharBeforAtSign(email string) bool {
	var containAtSign = existAtSign(email)

	if containAtSign {

		if indexOfAtSign(email) <= 20 {
			return true
		}
	}
	return false
}

func restrictedDomain(email string) bool {
	return strings.HasSuffix(email, ".co.id") || strings.HasSuffix(email, ".id")
}
