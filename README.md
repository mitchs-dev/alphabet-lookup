# Alphabet Lookup

This is a Go program that takes a string input from the user and converts each alphabetical character to its corresponding position in the alphabet. For example, 'a' or 'A' would be '1', 'b' or 'B' would be '2', and so on.

## How to Run

You can run this program in two ways:

1. Provide the string as a command line argument:

   ```bash
   go run alphabet-lookup.go "your string here"
   ```

2. Run the program without any arguments, and you will be prompted to enter a string:

   ```bash
   go run alphabet-lookup.go
   Enter a string: <your string here>
   ```

In both cases, the program will print the string with each alphabetical character replaced by its position in the alphabet. Non-alphabetical characters will be ignored.

## Importing and Using Convert Function

You can also import this package into your own Go program and use the `Convert` function. Here's an example:

```go
package main

import (
    "fmt"
    "alphabet-lookup" // replace with the actual package name
)

func main() {
    result := alphabet-lookup.Convert("your string here")
    fmt.Println(result)
}
```

In this example, replace `alphabet-lookup` with the actual name of the package, and `your string here` with the string you want to convert.

## Code Structure

The main function reads the user input, either from the command line arguments or from the standard input if no arguments are provided. It then calls the `Convert` function to convert the input string.

The Convert function takes a string as input and returns a string. It iterates over each character in the input string. If the character is a letter, it converts it to its corresponding number in the alphabet and adds it to the output string. If the character is not a letter, it is ignored.
