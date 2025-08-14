package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

func main() {
	url := "https://pixelsafari.neocities.org/buttons/"
	response, err := http.Get(url)
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

		// gifName := strings.Split(match[2], "/")[2]
		fmt.Println("Gif Found:", match)
		// downloadImage(url, gifName)
	}
}

func downloadImage(baseURL, gifName string) {
	downlaodURL, err := url.JoinPath(baseURL, gifName)
	if err != nil {
		return
	}
	fmt.Println("DOWNLOAD URL: ", downlaodURL)
}
