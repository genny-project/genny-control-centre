// ASCII Colour Utilities
package main

// ASCII colour chars
var colourReset string = "\033[0m"
var colourRed string = "\033[31m"
var colourGreen string = "\033[32m"
var colourYellow string = "\033[33m"
var colourBlue string = "\033[34m"
var colourPurple string = "\033[35m"
var colourCyan string = "\033[36m"
var colourWhite string = "\033[37m"

// Return the string in ASCII Red
func Red(str string) string {
	return colourRed + str + colourReset
}

// Return the string in ASCII Green
func Green(str string) string {
	return colourGreen + str + colourReset
}

// Return the string in ASCII Yellow
func Yellow(str string) string {
	return colourYellow + str + colourReset
}

// Return the string in ASCII Blue
func Blue(str string) string {
	return colourBlue + str + colourReset
}

// Return the string in ASCII Purple
func Purple(str string) string {
	return colourPurple + str + colourReset
}

// Return the string in ASCII Cyan
func Cyan(str string) string {
	return colourCyan + str + colourReset
}

// Return the string in ASCII White
func White(str string) string {
	return colourWhite + str + colourReset
}
