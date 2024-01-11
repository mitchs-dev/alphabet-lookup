// This program takes an alphabetic string and converts each character to its corresponding number
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the user inline input
	userInput := strings.Join(os.Args[1:], " ")

	// If no inline input, read from stdin
	if len(userInput) == 0 {
		fmt.Print("Enter a string: ")
		reader := bufio.NewReader(os.Stdin)
		userInput, _ = reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput) // remove the newline at the end
	}

	// Convert the input string
	outputString := Convert(userInput)

	// Print the output
	fmt.Println(outputString)
}

// Convert takes an alphabetic string and converts each character to its corresponding number
func Convert(inputString string) string {
	// Provide a variable to store the output
	var outputString strings.Builder

	// Iterate over the string and convert each character to its corresponding number
	for _, char := range inputString {
		// If the character is a letter, convert it to its corresponding number
		// If the character is not a letter, ignore it
		// Provide a space between each character
		if 'a' <= char && char <= 'z' {
			outputString.WriteString(fmt.Sprintf(" %d", char-'a'+1))
		} else if 'A' <= char && char <= 'Z' {
			outputString.WriteString(fmt.Sprintf(" %d", char-'A'+1))
		}
	}

	// Return the output string
	return outputString.String()
}
