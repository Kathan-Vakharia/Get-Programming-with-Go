package main

import "fmt"

const msg = "Hola Estación Espacial Internacional"

func main() {
	for _, char := range msg {
		char += 13
		switch {
		case char >= 'a' && char <= 'z': //if lower-case

			if char > 'z' {
				char -= 26
			}
		case char >= 'A' && char <= 'Z': // if upper-case

			if char > 'Z' {
				char -= 26
			}

		}
		fmt.Printf("%c", char)

	}
}
