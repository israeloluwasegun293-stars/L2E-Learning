package main

import "strings"

func GetASCIIArt(input string, lines []string) string {

	// convert escaped \n into real newlines
	input = strings.ReplaceAll(input, "\\n", "\n")

	const letterHeight = 8
	var output strings.Builder

	linesInput := strings.Split(input, "\n")

	// loop through each line of input
	for _, line := range linesInput {

		// ✅ FIX: handle empty lines properly
		if line == "" {
			output.WriteString("\n")
			continue
		}

		// loop through each row of ASCII art
		for i := 0; i < letterHeight; i++ {

			// loop through each character
			for _, char := range line {

				// ignore invalid ASCII characters
				if char < 32 || char > 126 {
					output.WriteString("nul")
					continue
				}

				// locate character in banner
				index := (int(char) - 32) * 9

				// append the correct row
				output.WriteString(lines[index+i+1])

				// optional spacing between characters
				// output.WriteString(" ")
			}

			// move to next ASCII row
			output.WriteString("\n")
		}
	}

	return output.String()
}