## Go Reloaded 

Go reloaded is a simple text processing, and editor. It reads a text file, applies corrections to it, then writes the corrected text to a new file.

### Features 

1. **Hexadecimal to Decimal Conversion:** every instance of (hex) converts the previous word to their decimal equivalents.

2. **Binary Conversion:** every instance of (bin) converts the previous word to their decimal equivalents

3. **Uppercase Conversion:** every instance of (upper) converts the previous word to uppercase.

4. **Lowercase Conversion:** every instance of (lower) converts the previous word to lowercase.

5. **Capitalization:** Every instance of (cap) capitalizes the previous word

6. **Multiple Word Transformations:** Every instance of(up, 2)(low, 3) (cap, 4)This means the transformation should be applied to the previous N words.

7. **Punctuation Formatting:** The program fixes spacing around punctuation (. , ! ? : ;) Punctuation should be attached to the previous word. A space must appear after punctuation.
   
8. **Quote Formatting:** Single quotes ' should surround the text inside them without spaces.
   
9.  **Grammar Correction (a → an):** If the article a appears before a word starting with a vowel or h, it should become an.

### My project structure

go-reloaded/
├── go.mod
├── main.go
|--my helper-functions.go
├── README.md
|--result.txt
|--sample.txt

