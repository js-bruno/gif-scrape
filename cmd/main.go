package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type WebSite struct {
	url  string
	name string
	gifs []string
}

func NewWebSite(url string) WebSite {
	// name := strings.Split(url, "//")
	reSearchName := regexp.MustCompile(`//([^.]+)\.`)
	name := reSearchName.FindStringSubmatch(url)

	return WebSite{
		url:  url,
		name: name[1],
		gifs: nil,
	}
}

func main() {
	url := "https://pixelsafari.neocities.org/buttons/"
	webSite := NewWebSite(url)

	response, err := http.Get(webSite.url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	htmlBodyText, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	reSearchTags := regexp.MustCompile(`<\s*([img]+)(\s[^>]*)?>`)
	matches := reSearchTags.FindAllStringSubmatch(string(htmlBodyText), -1)

	for _, match := range matches {

		gifName := strings.Split(match[2], "/")[2]
		gifName = strings.Trim(gifName, "\"")
		downloadImage(url, gifName)
	}
}

func downloadImage(baseURL, gifName string) {
	downlaodURL, err := url.JoinPath(baseURL, gifName)
	if err != nil {
		return
	}
	fmt.Println("DOWNLOAD URL: ", downlaodURL)
}
