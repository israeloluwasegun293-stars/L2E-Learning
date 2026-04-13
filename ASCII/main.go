package main

import (
	"fmt"
	"os"
)

func main() {

	arg := "Invalid number of arguements!"
	if len(os.Args) != 3 {
		fmt.Println(arg)
		return
	}

	input := os.Args[1]
	font := os.Args[2]

	art := GetASCIIArt(input, font)

	fmt.Print(art)
}
