package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//so, we check for command line arguements
	if len(os.Args) != 3 {
		fmt.Println("Invalid length of arguements")
		return
	}

	//we get the file names and store it in a variable
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	//we read the input file
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading File:", err)
		return
	}

	//the ReadFile returns what it read through in Bytes format in default, so we convert it to string for readability
	data := string(file)

	//so we need the string to be a slice of string

	arrayData := strings.Fields(data)

	//we need to make the slice of string be in a string form, i.e seperated by a space(" ")

	joinedArray := strings.Join(arrayData, " ")

	//we bring in the processed Text and assign it to the variable result
	result := processText(joinedArray)

	//here, we write to the output file, the 0644 is the permission code
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error Writing file:", err)
		return
	}
	//sucess message
	fmt.Println("Successfully written to:", outputFile)
}

func processText(text string) string {
	text = convertHexBin(text)
	text = upCase(text)
	text = caseLow(text)
	text = convertCap(text)
	text = convertUpNum(text)
	text = convertLowNum(text)
	text = convertCapnum(text)
	text = fixPunctuation(text)
	text = fixArticle(text)
	// Each transformation will be added here one by one
	return text
}
