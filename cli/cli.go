package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ashton0223/go-quizlet-scraper/export"
	"github.com/ashton0223/go-quizlet-scraper/scraper"
)

// Errors are in red, success in green
var (
	red   = "\033[1;31m%s\033[0m"
	green = "\033[1;32m%s\033[0m"
)

func main() {
	var url string
	var filetype string
	var name string

	args := os.Args[1:]
	length := len(args)

	// Check length of command line arguments, and set values given
	switch length {
	case 3:
		url = args[0]
		filetype = args[1]
		name = args[2]
		break
	case 2:
		url = args[0]
		filetype = args[1]
		name = getName()
		break
	case 1:
		url = args[0]
		filetype = getFiletype()
		name = getName()
		break
	case 0:
		url = getUrl()
		filetype = getFiletype()
		name = getName()
		break
	default:
		fmt.Println("Too many arguments")
		os.Exit(0)
	}

	termArr, defArr, err := scraper.GetStudySet(url)
	if err != nil {
		fmt.Println(fmt.Sprintf(red, err))
		os.Exit(1)
	}

	err = export.CreateSheet(termArr, defArr, filetype, name)
	if err != nil {
		fmt.Println(fmt.Sprintf(red, err))
		os.Exit(1)
	} else {
		fmt.Println(fmt.Sprintf(green, "Successfully created spreadsheet"))
	}
}

func getInput(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(message)
	fmt.Print("--> ")
	input, _ := reader.ReadString('\n')
	input = strings.ReplaceAll(input, "\n", "")
	return input
}

func getUrl() string {
	return getInput("Enter the Quizlet URL")
}

func getFiletype() string {
	return getInput("Enter the desired filetype (csv or tsv)")
}

func getName() string {
	return getInput("Enter the name of the spreadsheet")
}
