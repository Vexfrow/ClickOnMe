package checker

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	FileURLs   string
	RateLimit  int
	URL        string
	OutputFile string
)

type Result int

var resName = map[Result]string{
	OK:    "OK",
	NOK:   "NOK",
	MAYBE: "MAYBE",
	ERROR: "ERROR",
}

// Result that can be returned from "CheckURL"
const (
	result Result = iota
	OK
	NOK
	MAYBE
	ERROR
)

func (res Result) toString() string {
	return resName[res]
}

func fillListFromFile() []string {
	file, err := os.Open(FileURLs)
	if err != nil {
		fmt.Printf("[ERROR] An error occurred while trying to open the file : %s\nPlease check the name of the file\n", FileURLs)
		os.Exit(1)
	}
	var listURLs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		listURLs = append(listURLs, scanner.Text())
	}

	return listURLs
}

func testURL(url string) Result {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("[Error] Error while testing \"%s\" : %s\n", url, err)
		return ERROR
	}

	// Check X-Frame-Option to see if it accepts Iframe from web pages of remote origin
	xFrameOption := resp.Header["X-Frame-Options"]
	if xFrameOption != nil && (strings.Contains(xFrameOption[0], "deny") || strings.Contains(xFrameOption[0], "sameorigin")) {
		return NOK
	}

	if resp.Header["Content-Security-Policy"] != nil {
		return MAYBE
	}

	return OK
}

func checkListURL(checkListURL []string) {

	//Determine either the result should be printed on stdout or written in a file
	writeFD := os.Stdout

	if OutputFile != "" {
		file, err := os.OpenFile(OutputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		if err == nil {
			writeFD = file
		} else {
			fmt.Printf("[ERROR] The file \"%s\" cannot be opened : %s\n", OutputFile, err)
		}
	}

	fmt.Print("\nStart checking URL\n\n")

	for _, url := range checkListURL {
		answer := testURL(url)
		fmt.Fprintf(writeFD, "|  %s\t|\t%s\t|\n", url, answer.toString())
	}

}

func StartChecking() {
	var listURLs []string

	if (FileURLs != "" && URL != "") || (FileURLs == "" && URL == "") {
		fmt.Printf("You should use \"--inputFile\" or \"--url\" at least, but not both at the same time\n")
		os.Exit(1)
	} else if FileURLs != "" {
		listURLs = fillListFromFile()
	} else {
		listURLs = append(listURLs, URL)
	}

	checkListURL(listURLs)

	fmt.Print("\nEvery URL has been verified !\n")

}
