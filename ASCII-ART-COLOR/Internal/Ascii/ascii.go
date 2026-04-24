package ascii

import (
	"bufio"
	"fmt"
	"os"
)

type Asciiart struct {
	chars map[rune][]string

	banner string
}

func NewAsciiArt(banner string) Asciiart {
	return Asciiart{
		chars:  make(map[rune][]string),
		banner: banner,
	}
}
func (c Asciiart) Load() error {

	file, err := os.Open(c.banner)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	return nil
}
