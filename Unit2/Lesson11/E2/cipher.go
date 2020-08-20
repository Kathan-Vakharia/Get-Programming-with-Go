package main

import (
	"fmt"
	"strings"
	
)

func main() {
	message := "My name is Kathan Vakharia"
	cipheredMessage := ""
	keyword := "POLYGLOTPROGRAMMER"
	keyIndex := 0

	convertedMessage := strings.ToUpper(strings.ReplaceAll(message, " ", ""))

	fmt.Println(convertedMessage, keyword)
	for i := 0; i < len(convertedMessage); i++ {
		//convertin character to their offset from 'A'
		c := convertedMessage[i] - 'A'
		key := keyword[keyIndex] - 'A'

		//calculating offset for ciphered text
		offset := (c+key)%26 + 'A'

		cipheredMessage += string(offset)

		//handling out of bounds for keyword
		keyIndex++
		keyIndex %= len(keyword)
	}
	fmt.Println("The original message is", message)
	fmt.Println("The ciphered message is", cipheredMessage)

}
