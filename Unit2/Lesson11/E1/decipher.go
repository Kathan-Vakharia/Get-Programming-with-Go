package main

import "fmt"

func main() {
	cipheredMsg := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0

	for i := 0; i < len(cipheredMsg); i++ {

		//calculating the offset from letter 'A'
		offset := cipheredMsg[i] - keyword[keyIndex]

		//Adding offset to 'A'
		decipheredC := (offset+26)%26 + 'A'

		//converting the byte to string
		message += string(decipheredC)

		//handling out of of keyIndex after iteration
		keyIndex++
		keyIndex %= len(keyword)
	}
	fmt.Printf("Original Message : %s\n", cipheredMsg)
	fmt.Printf("Deciphered Message : %s\n", message)

}
