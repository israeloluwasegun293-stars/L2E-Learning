package main

import (
	"fmt"
	"os"
	"strings"
)

func GetASCIIArt(input string, font string) string {

	if font != "standard" && font != "shadow" && font != "thinkertoy" {
		fmt.Println("Invalid font")
		return ""
	}

	file, err := os.ReadFile("fonts/" + font + ".txt")

	if err != nil {
		fmt.Println("Error reading banner file")
		return ""
	}
	lines := strings.Split(string(file), "\n")

	// convert escaped \n into real newlines
	input = strings.ReplaceAll(input, "\\n", "\n")

	const letterHeight = 8
	var output strings.Builder

	linesInput := strings.Split(input, "\n")

	// loop through each line of input
	for _, line := range linesInput {

		// handle empty lines properly
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

/*
what this function actually does in real life terms.

1. CHECK IF FONT IS VALID

2. READ THE FONT FILE

4. SPLIT FILE INTO LINES

5.HANDLE "\n" FROM INPUT using strings.RepalaceAll replacing all \\n with actual \n

6. SET HEIGHT

7. PREPARE OUTPUT

8. SPLIT INPUT INTO LINES

9. LOOP EACH LINE

10. LOOP ROWS (VERY IMPORTANT)

11. LOOP EACH CHARACTER

12. FIND CHARACTER IN FILE
index := (int(char) - 32) * 9

13. GET THE CORRECT ROW
output.WriteString(lines[index+i+1])

14. AFTER EACH ROW
output.WriteString("\n")

15. FINAL STEP
return output.String()
*/
