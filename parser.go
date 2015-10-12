package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func checkerr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	Parser("https://exploit-db.com/search/?order_by=date&order=desc&pg=1&action=search&description=SQL")
}

func Parser(url string) {
	doc, err := goquery.NewDocument(url)
	checkerr(err)
	doc.Find(".description").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		link = url + link
		fmt.Printf("%s\n", link)
	})
}
