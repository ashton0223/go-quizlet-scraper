package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"github.com/ashton0223/go-quizlet-scraper/scraper"
	"github.com/ashton0223/go-quizlet-scraper/export"
)

func main() {
	var url string
	var filetype string
	filetypeGiven := false

	args := os.Args[1:]
	length := len(args)
	if length > 0 {
		if length > 2 {
			fmt.Println("Too many arguments")
			os.Exit(0)
		} else {
			url = args[0]
			if length > 1 {
				filetype = args[1]
				filetypeGiven = true 
			}
		} 
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the Quizlet URL")
		fmt.Print("--> ")
		url, _ = reader.ReadString('\n')
		url = strings.ReplaceAll(url, "\n", "")
		fmt.Println()
	}

	termArr, defArr, err := scraper.GetStudySet(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !filetypeGiven {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the desired filetype (csv or tsv)")
		fmt.Print("--> ")
		filetype, _ = reader.ReadString('\n')
		filetype = strings.ReplaceAll(filetype, "\n", "")
	}
	export.CreateSheet(termArr, defArr, filetype)
}
