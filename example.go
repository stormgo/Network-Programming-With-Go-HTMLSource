package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var servers = []string{
	"http://www.google.com/search?client=ubuntu&channel=fs&q=go+language&ie=utf-8&oe=utf-8",
	"http://golang.org/",
	"http://golang.org/doc/go_tutorial.html"}

var maxOpenRequests int = 3

var rpmRegexp = regexp.MustCompile("href=\"[^\"]+\"")

func main() {
	search := getSearchTerm()
	out := make(chan string)
	done := make(chan int)
	go printer(out)
	for _, url := range servers {
		go searchURL(search, url, out, done)
	}
	for i := 0; i < len(servers); i++ {
		<-done
	}
}

func getSearchTerm() string {
	args := os.Args
	if len(args) != 2 {
		die("Please enter one search term")
	}
	return args[1]
}

func die(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func printer(out chan string) {
	for {
		fmt.Println(<-out)
	}
}

var requestSemaphore = make(chan int, maxOpenRequests) // Integer chanel with a maximum queue size

func searchURL(search string, url string, out chan string, done chan int) {
	requestSemaphore <- 1 // Block until put in the semaphore queue
	response, realURL, err := http.Get(url)
	if err == nil {
		bufferedReader := bufio.NewReader(response.Body)
		err = searchAll(search, bufferedReader, url, out)
		response.Body.Close()
	}
	if err != nil {
		die("Could not read from " + realURL + ":" + err.Error())
	}
	<-requestSemaphore // Dequeue from the semaphore
	done <- 1          // Signal that function is done
}

func searchAll(search string, reader *bufio.Reader, httpRoot string, out chan string) error {
	var error error = nil
	for {
		var line, err = reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return error
		}
		var hrefs = rpmRegexp.AllMatchesStringIter(line, 0)
		for href := range hrefs {
			start := strings.Index(href, "\"") + 1
			end := len(href) - 1
			packageFile := href[start:end]
			if strings.Index(packageFile, search) != -1 {
				out <- httpRoot + packageFile
			}
		}
	}
	return error
}
