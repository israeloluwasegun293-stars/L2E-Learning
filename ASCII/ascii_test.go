package main

// import (
// 	"os"
// 	"strings"
// 	"testing"
// )

// // helper function to load banner
// func loadBanner(t *testing.T) []string {
// 	file, err := os.ReadFile("standard.txt")
// 	if err != nil {
// 		t.Fatalf("failed to read standard.txt: %v", err)
// 	}
// 	return strings.Split(string(file), "\n")
// }

// func TestSingleCharacter(t *testing.T) {
// 	lines := loadBanner(t)

// 	result := GetASCIIArt("A", lines)

// 	if result == "" {
// 		t.Errorf("Expected ASCII art for 'A', got empty string")
// 	}
// }

// func TestWord(t *testing.T) {
// 	lines := loadBanner(t)

// 	result := GetASCIIArt("Hi", lines)

// 	if !strings.Contains(result, "\n") {
// 		t.Errorf("Expected multi-line ASCII output, got: %q", result)
// 	}
// }

// func TestNewLineInput(t *testing.T) {
// 	lines := loadBanner(t)

// 	result := GetASCIIArt("Hello\\nWorld", lines)

// 	linesOut := strings.Split(result, "\n")

// 	if len(linesOut) < 16 {
// 		t.Errorf("Expected multiple ASCII lines for multi-line input, got: %d lines", len(linesOut))
// 	}
// }

// func TestEmptyInput(t *testing.T) {
// 	lines := loadBanner(t)

// 	result := GetASCIIArt("", lines)

// 	if result != "\n" {
// 		t.Errorf("Expected empty string, got: %q", result)
// 	}
// }

// func TestInvalidCharacter(t *testing.T) {
// 	lines := loadBanner(t)

// 	result := GetASCIIArt("\x01", lines)

// 	if !strings.Contains(result, "nul") {
// 		t.Errorf("Expected 'nul' for invalid character, got: %q", result)
// 	}
// }
