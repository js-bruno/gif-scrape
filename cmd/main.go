package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
)

type WebSite struct {
	url  string
	name string
	gifs []Gif
}

type Gif struct {
	url  string
	name string
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
	urlWeb := "https://pixelsafari.neocities.org/buttons/"
	webSite := NewWebSite(urlWeb)
	response, err := http.Get(webSite.url)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer response.Body.Close()
	htmlBodyText, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	reSearchTags := regexp.MustCompile(`<\s*([img]+)(\s[^>]*)?>`)
	matches := reSearchTags.FindAllStringSubmatch(string(htmlBodyText), -1)

	for _, match := range matches {
		fmt.Println(match[2])
		gifName := strings.Split(match[2], "/")[2]
		gifName = strings.Trim(gifName, "\"")

		gifURL, err := url.JoinPath(webSite.url, gifName)

		fmt.Println(gifURL)
		if err != nil {
			fmt.Println(err.Error())
		}
		webSite.gifs = append(webSite.gifs, Gif{
			url:  gifURL,
			name: gifName,
		})
	}

	var wg sync.WaitGroup
	for i, gif := range webSite.gifs {
		wg.Add(1)
		go downloadWorker(i, &wg, webSite.name, gif)
	}
	wg.Wait()
}

func downloadImage(siteName string, gif Gif) (err error) {
	resp, err := http.Get(gif.url)
	if err != nil {
		return err
	}

	os.Mkdir(siteName, 0755)
	if err != nil {
		return err
	}
	gifAddrs := siteName + "/" + gif.name
	file, err := os.Create(gifAddrs)
	fmt.Println(gifAddrs)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("Downloading: " + gifAddrs)

	return nil

}

func downloadWorker(id int, wg *sync.WaitGroup, siteName string, gif Gif) {
	defer wg.Done()
	fmt.Printf("Worker %d start:\n", id)
	downloadImage(siteName, gif)
	fmt.Printf("Worker %d over:\n", id)
}
