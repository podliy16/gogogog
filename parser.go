package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"sort"
	"strings"
)

type LinksForParsing struct {
	link string
}

var template1 string = "https://www.exploit-db.com/exploits/*****"

func checkerr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func SortString(w string) string {
	sorting := strings.Split(w, "")
	sort.Strings(sorting)
	return strings.Join(sorting, "")
}

func writeToJson(url LinksForParsing, filename string) {
	j, err := json.Marshal(url)
	checkerr(err)
	j = append(j, "\n"...)
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if _, err = f.Write(j); err != nil {
		checkerr(err)
	}
}

func main() {
	if _, err := os.Stat("Links.json"); os.IsNotExist(err) {
		fmt.Println("Links.json not found")
		return
	}
	foundedUrls := make(map[string]bool)
	seedUrls := os.Args[1:]
	channelUrls := make(chan string)
	channelFinish := make(chan bool)
	for _, url := range seedUrls {
		go crawler(url, channelUrls, channelFinish)
	}
	for i := 0; i < len(seedUrls); {
		select {
		case url := <-channelUrls:
			foundedUrls[url] = true
		case <-channelFinish:
			i++
		}
	}
	url := SortString(template1)
	fmt.Println(url)
	writeToJson(url, "Links.json")

	for url, _ := range foundedUrls {
		fmt.Println(" - " + url)
	}

	close(channelUrls)
}

func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	return
}

func crawler(url string, channel chan string, channelFinish chan bool) {
	responce, err := http.Get(url)
	defer func() {
		channelFinish <- true
	}()
	checkerr(err)
	body := responce.Body
	defer body.Close()
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()
			aTag := t.Data == "a"
			if !aTag {
				continue
			}
			ok, url := getHref(t)
			if !ok {
				continue
			}

			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				channel <- url
			}
		}
	}
}
