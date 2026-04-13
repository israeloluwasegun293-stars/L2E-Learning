package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertHexBin(text string) string {
	//we split the text into words and store it in a variable
	words := strings.Fields(text)

	//then, we create a slice to store the results
	result := []string{}

	//looping through each wordf
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			// grab the previous word and convert from hex to decimal
			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				result[len(result)-1] = fmt.Sprintf("%d", num)
			}
			// we don't append "(hex)"  we skip it

			//same process for (bin)
		} else if words[i] == "(bin)" && i > 0 {

			// grab the previous word and convert from binary to decimal
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				result[len(result)-1] = fmt.Sprintf("%d", num)
			}
			// we don't append "(bin)" — we skip it

			//if the word is not (bin) or (hex), we just keep it the way it is
		} else {
			result = append(result, words[i])
		}
	}

	return strings.Join(result, " ")
}

func upCase(text string) string {

	//we split the text into words and store it in a variable
	words := strings.Fields(text)

	//then, we create a slice to store the results
	result := []string{}

	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" && i > 0 {
			result[len(result)-1] = strings.ToUpper(result[len(result)-1])
		} else {
			result = append(result, words[i])
		}
	}

	return strings.Join(result, " ")
}
func caseLow(text string) string {
	words := strings.Fields(text)
	result := []string{}

	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" && i > 0 {
			result[len(result)-1] = strings.ToLower(result[len(result)-1])
		} else {
			result = append(result, words[i])
		}
	}
	return strings.Join(result, " ")
}

func convertCap(text string) string {
	words := strings.Fields(text)
	result := []string{}
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" && i > 0 {
			word := result[len(result)-1]
			result[len(result)-1] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		} else {
			result = append(result, words[i])
		}
	}
	return strings.Join(result, " ")
}

func convertUpNum(text string) string {

	//we split the text here
	words := strings.Fields(text)

	//then we create a result slice
	result := []string{}

	//we loop through the words
	for i := 0; i < len(words); i++ {

		//we try to Detect (up,
		if words[i] == "(up," {

			//here we remove the bracket )
			numWord := strings.TrimRight(words[i+1], ")")

			//then we convert the string to number
			n, _ := strconv.Atoi(numWord)

			for j := 0; j < n; j++ {
				result[len(result)-1-j] = strings.ToUpper(result[len(result)-1-j])
			}

			i++

		} else {

			result = append(result, words[i])

		}
	}

	return strings.Join(result, " ")
}

func convertLowNum(text string) string {
	//we split the text first just like the other ones

	words := strings.Fields(text)

	result := []string{}

	for i := 0; i < len(words); i++ {

		//we try to Detect (low,
		if words[i] == "(low," {

			//here we remove the bracket )
			numWord := strings.TrimRight(words[i+1], ")")

			//then we convert the string to number
			n, _ := strconv.Atoi(numWord)

			for j := 0; j < n; j++ {
				result[len(result)-1-j] = strings.ToLower(result[len(result)-1-j])
			}

			i++

		} else {

			result = append(result, words[i])

		}

	}

	return strings.Join(result, " ")
}

func convertCapnum(text string) string {
	//we split the text first just like the other ones

	words := strings.Fields(text)

	result := []string{}

	for i := 0; i < len(words); i++ {

		//we try to Detect (cap,
		if words[i] == "(cap," {

			//here we remove the bracket )
			numWord := strings.TrimRight(words[i+1], ")")

			//then we convert the string to number
			n, _ := strconv.Atoi(numWord)

			for j := 0; j < n; j++ {
				word := result[len(result)-1-j]
				result[len(result)-1-j] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
			}

			i++

		} else {

			result = append(result, words[i])

		}

	}

	return strings.Join(result, " ")
}

func fixPunctuation(text string) string {
    // step 1: remove space BEFORE punctuation
    for _, p := range []string{",", ".", "!", "?", ":", ";"} {
        text = strings.ReplaceAll(text, " "+p, p)
    }
    // step 2: remove space BETWEEN punctuation marks
    for _, p := range []string{",", ".", "!", "?", ":", ";"} {
        for _, p2 := range []string{",", ".", "!", "?", ":", ";"} {
            text = strings.ReplaceAll(text, p+" "+p2, p+p2)
        }
    }
    // step 3: add space AFTER punctuation
    for _, p := range []string{",", ".", "!", "?", ":", ";"} {
        text = strings.ReplaceAll(text, p, p+" ")
    }
    // step 4: remove space BETWEEN punctuation marks again
    for _, p := range []string{",", ".", "!", "?", ":", ";"} {
        for _, p2 := range []string{",", ".", "!", "?", ":", ";"} {
            text = strings.ReplaceAll(text, p+" "+p2, p+p2)
        }
    }
    // step 5: clean up double spaces
    for strings.Contains(text, "  ") {
        text = strings.ReplaceAll(text, "  ", " ")
    }
    // step 6: trim trailing space
    text = strings.TrimSpace(text)
    return text
}
func fixArticle(text string) string {
    words := strings.Fields(text)
    for i := 0; i < len(words)-1; i++ {
        if words[i] == "a" || words[i] == "A" {
            // check if next word starts with a vowel or h
            nextWord := strings.ToLower(words[i+1])
            if strings.ContainsRune("aeiuoh", rune(nextWord[0])) {
                // preserve the case
                if words[i] == "A" {
                    words[i] = "An"
                } else {
                    words[i] = "an"
                }
            }
        }
    }
    return strings.Join(words, " ")
}
