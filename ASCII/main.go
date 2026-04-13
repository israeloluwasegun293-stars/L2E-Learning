package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	arg := "Invalid number of arguements!"
	if len(os.Args) != 2 {
		fmt.Println(arg)
		return
	}

	input := os.Args[1]

	file, _ := os.ReadFile("standard.txt")

	line := strings.Split(string(file), "\n")

	art := GetASCIIArt(input, line)

	fmt.Print(art)
}
