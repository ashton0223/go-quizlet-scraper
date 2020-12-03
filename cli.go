package main

import (
	"fmt"
//	"net/http"
//	"io/ioutil"
//	"regexp"
	"bufio"
	"os"
	"strings"

	"github.com/ashton0223/go-quizlet-scraper/scraper"
)

func main() {
	var url string
	args := os.Args[1:]
	if len(args) > 0 {
		if len(args) > 1 {
			fmt.Println("Too many arguments")
			os.Exit(0)
		} else {
			url = args[0]
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the Quizlet URL")
		fmt.Print("--> ")
		url, _ = reader.ReadString('\n')
		url = strings.ReplaceAll(url, "\n", "")
	}

	termArr, defArr, err := scraper.GetStudySet(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(termArr,defArr)
}
