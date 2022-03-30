package main

// text colours
var colourReset string = "\033[0m"
var colourRed string = "\033[31m"
var colourGreen string = "\033[32m"
var colourYellow string = "\033[33m"
var colourBlue string = "\033[34m"
// colourPurple := "\033[35m"
// colourCyan := "\033[36m"
// colourWhite := "\033[37m"

func red(str string) string {
	return colourRed + str + colourReset
}

func green(str string) string {
	return colourGreen + str + colourReset
}

func yellow(str string) string {
	return colourYellow + str + colourReset
}

func blue(str string) string {
	return colourBlue + str + colourReset
}
