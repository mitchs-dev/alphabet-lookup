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

## Reverse conversion

You can also convert a string of numbers back to its original string. To do this, run the program with the `-r` flag:

```bash
go run alphabet-lookup.go -r "your number string here"
```

## Importing and Using Convert Function

You can also import this package into your own Go program and use the `Convert` function. Here's an example:

```go
package main

import (
    "fmt"
    "github.com/mitchs-dev/alphabet-lookup"
)

func main() {
    result := alphabet-lookup.Convert("your string here")
    fmt.Println(result)
}
```

In this example, replace `your string here` with the string you want to convert.

### Reverse conversion function

You can also import the `ReverseConvert` function to convert a string of numbers back to its original string. Here's an example:

```go
package main

import (
    "fmt"
    "github.com/mitchs-dev/alphabet-lookup"
)

func main() {
    result := alphabet-lookup.ReverseConvert("your number string here")
    fmt.Println(result)
}
```

In this example, replace `your number string here` with the string you want to convert.

## Code Structure

The main function reads the user input, either from the command line arguments or from the standard input if no arguments are provided. It then calls the `Convert` function to convert the input string. Or, if the `-r` flag is provided, it calls the `ReverseConvert` function to convert the input string.

The Convert function takes a string as input and returns a string. It iterates over each character in the input string. If the character is a letter, it converts it to its corresponding number in the alphabet and adds it to the output string. If the character is not a letter, it is ignored.

The ReverseConvert function takes a string of numbers as input and returns a string. It iterates over each character in the input string. If the character is a number, it converts it to its corresponding letter in the alphabet and adds it to the output string. If the character is not a number, it is ignored.
