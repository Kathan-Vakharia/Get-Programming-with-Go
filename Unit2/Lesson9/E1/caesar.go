package main

import "fmt"

const msg = "L fdph, L vdz, L frqtxhuhg."

func main() {
	for _, char := range msg {
		//subtracting 3 to decipher
		char = char - 3

		switch {
		case char >= 'a' && char <= 'z': //if lower-case

			if char < 'a' {
				char += 26
			}
		case char >= 'A' && char <= 'Z':// if upper-case

			if char < 'A' {
				char += 26
			}

		}
		fmt.Printf("%c", char)

	}

}
