package main

import (
	"fmt"
	"os"
	"strings"

	ascii "github.com/israeloluwasegun293-stars/Internal/Ascii"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println(`Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored>"something"`)
		return
	}

	if !(strings.HasPrefix(os.Args[1], "--color=")) {
		fmt.Println(`Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored>"something"`)
		return
	}

	color := strings.TrimPrefix(os.Args[1], "--color=")

	substring := os.Args[2]

	text := os.Args[3]

	fmt.Println(color, substring, text)

	art := ascii.NewAsciiArt("Banner-file/shadow.txt",)
	err := art.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
}
