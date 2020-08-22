package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	section := `As far as eye could reach he saw nothing but the stems of the great plants about him
receding in the violet shade, and far overhead the multiple transparency of huge leaves
filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever
he felt able he ran again; the ground continued soft and springy, covered with the same
resilient weed which was the first thing his hands had touched in Malacandra. Once or
twice a small red creature scuttled across his path, but otherwise there seemed to be no
life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned
and alone in a forest of unknown vegetation thousands or millions of miles beyond the
reach or knowledge of man.

—C.S. Lewis, Out of the Silent Planet,`
	//converting text to lowercase
	section = strings.ToLower(section)

	//declaring map
	wordCount := make(map[string]int)

	//Anonymous function to check if a  codepoint is alphanumeric
	isAlphanumeric := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	//splitting the section on sequence of non-alphanumeric letters
	wordSlice := strings.FieldsFunc(section, isAlphanumeric)

	//finding frequency of every words
	for _, word := range wordSlice {
		wordCount[word]++
	}

	//displaying words having more than one count
	for word, freq := range wordCount {
		if freq > 1 {
			fmt.Printf("Word :%s , Count:%d\n", word, freq)
		}
	}

}
