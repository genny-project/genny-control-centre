package main

// text colours
var colourReset string = "\033[0m"
var colourRed string = "\033[31m"
var colourGreen string = "\033[32m"
var colourYellow string = "\033[33m"
var colourBlue string = "\033[34m"
var colourPurple string = "\033[35m"
var colourCyan string = "\033[36m"
var colourWhite string = "\033[37m"

func Red(str string) string {
	return colourRed + str + colourReset
}

func Green(str string) string {
	return colourGreen + str + colourReset
}

func Yellow(str string) string {
	return colourYellow + str + colourReset
}

func Blue(str string) string {
	return colourBlue + str + colourReset
}

func Purple(str string) string {
	return colourPurple + str + colourReset
}

func Cyan(str string) string {
	return colourCyan + str + colourReset
}

func White(str string) string {
	return colourWhite + str + colourReset
}
