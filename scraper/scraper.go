package scraper

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"errors"
	"strings"
)

// Returns the Quizlet study set from a given URL in the form of two slices.
// This will return empty slices if the URL is not a Quizlet URL.

func GetStudySet(url string) ([]string, []string, error) {
	termArr := make([]string, 0)
	defArr := make([]string, 0)

	client := &http.Client{}
	req, err := http.NewRequest("GET",url, nil)
	if err != nil {
		panic(err)
	}

	// Avoid the CAPTCHA
	req.Header.Set("User-Agent", "Mozilla/5.0 (Android 4.4; Mobile; rv:41.0) Gecko/41.0 Firefox/41.0")

	resp, err := client.Do(req)
	if err != nil {
		returnString := err.Error()
		if strings.Contains(err.Error(), "unsupported protocol scheme") {
			returnString = "Not a valid URL. Is there a 'https://' at the beginning?"
		} else if strings.Contains(err.Error(), "connection refused") {
			returnString = "Connection refused. Are you connected to the network?"
		} else if strings.Contains(err.Error(), "no such host") {
			returnString = "Not a valid URL. Did you type in the address correctly?"
		}
		funcErr := errors.New(returnString)
		return nil, nil, funcErr
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	r, _ := regexp.Compile("<\\s*span class=\"TermText[^>]*>(?P<Text>(.*?))<\\s*/\\s*span>")
	for i, text := range r.FindAllStringSubmatch(string(body), -1) {
		if i % 2 == 0 {
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
