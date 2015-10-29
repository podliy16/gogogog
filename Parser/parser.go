package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"regexp"
)

func checkerr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

type LinksType struct {
	Link string
	//	Cve         string
	//	Description string
	//	PatchLink   string
	//	Danger      int64
}

var url string

func main() {
	parseStartLink()
	parseTextFromStartLink()
	searchAndWriteExploit()
	err := os.Remove("temp.txt")
	checkerr(err)
	err = os.Remove("waste.txt")
	checkerr(err)
	err = os.Remove("textTemp.txt")
	checkerr(err)
	os.Exit(0)
}

func parseStartLink() {
	fmt.Println("Input url: ")
	fmt.Scanf("%s", &url)
	firstDoc, err := goquery.NewDocument(url)
	checkerr(err)
	firstDoc.Find("tbody").Each(func(i int, tbody *goquery.Selection) {
		tbody.Find(".description").Each(func(j int, s *goquery.Selection) {
			link, _ := s.Find("a").Attr("href")
			x, _ := regexp.MatchString(`https://www.exploit-db.com/exploits/.....`, link)
			if x == true {
				file, err := os.OpenFile("temp.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
				checkerr(err)
				_, err = file.WriteString(link + "\n")
				checkerr(err)
				file.Close()
			}
			y, _ := regexp.MatchString(`/docs/......pdf`, link)
			if y == true {
				wasteUrl, err := os.OpenFile("waste.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
				checkerr(err)
				_, err = wasteUrl.WriteString(link)
				checkerr(err)
				wasteUrl.Close()
			}
		})
	})
}

func parseTextFromStartLink() {
	file, err := os.Open("temp.txt")
	checkerr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		j := scanner.Text()
		secondDoc, err := goquery.NewDocument(j)
		checkerr(err)
		text := secondDoc.Find("pre").Text()
		textFile, err := os.OpenFile("textTemp.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		checkerr(err)
		_, err = textFile.WriteString(text + "\n")
		checkerr(err)
	}
}

func searchAndWriteExploit() {
	text, err := os.OpenFile("textTemp.txt", os.O_RDWR, 0666)
	checkerr(err)
	defer text.Close()
	scanner := bufio.NewScanner(text)
	for scanner.Scan() {
		scan := scanner.Text()
		x, _ := regexp.MatchString(`option=com_rpl.*[d1t']`, scan)
		if x == true {
			re := regexp.MustCompile(`option=com_rpl.*[d1t']`)
			inj := re.FindString(scan)
			links := LinksType{Link: inj}
			marshToJson(links, "Links.json")
		}
	}
}

func marshToJson(links LinksType, filename string) {
	j, err := json.Marshal(links)
	checkerr(err)
	j = append(j, "\n"...)
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if _, err = f.Write(j); err != nil {
		checkerr(err)
	}
}
