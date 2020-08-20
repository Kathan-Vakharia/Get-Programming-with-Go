package main

import "fmt"

func main() {
	yesNo := "true"

	var launch bool
	switch yesNo {
	case "true", "1", "yes":
		launch = true
	case "false", "0", "no":
		launch = false
	default:
		fmt.Println(yesNo, "is not a valid argument to convert into boolean")
	}
	fmt.Println("Ready for launch? ", launch)
}
