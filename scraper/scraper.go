package scraper

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// Returns the Quizlet study set from a given URL in the form of two slices.

func GetStudySet(url string) ([]string, []string, error) {
	termArr := make([]string, 0)
	defArr := make([]string, 0)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	// Avoid the CAPTCHA
	req.Header.Set("User-Agent", "Mozilla/5.0 (Android 4.4; Mobile; rv:41.0) Gecko/41.0 Firefox/41.0")

	// Do the request
	resp, err := client.Do(req)
	if err != nil {
		funcErr := errors.New(reqErr(err))
		return nil, nil, funcErr
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	// Search for terms in body of website
	r, _ := regexp.Compile("<\\s*span class=\"TermText[^>]*>(?P<Text>(.*?))<\\s*/\\s*span>")
	for i, text := range r.FindAllStringSubmatch(string(body), -1) {
		text[1] = strings.ReplaceAll(text[1], "<br>", "\r")
		if i%2 == 0 {
			termArr = append(termArr, text[1])
		} else {
			defArr = append(defArr, text[1])
		}
	}

	// Check to make sure that it was actually a study set URL
	if len(termArr) == 0 {
		err = errors.New("Not a valid Quizlet study set URL")
	}
	return termArr, defArr, err
}

// Handle errors from the HTTP request
func reqErr(err error) string {
	returnString := err.Error()
	if strings.Contains(err.Error(), "unsupported protocol scheme") {
		returnString = "Not a valid URL. Is there a 'https://' at the beginning?"
	} else if strings.Contains(err.Error(), "connection refused") {
		returnString = "Connection refused. Are you connected to the network?"
	} else if strings.Contains(err.Error(), "no such host") {
		returnString = "Not a valid URL. Did you type in the address correctly?"
	}
	return returnString
}
