package main

import (
	"fmt"
	"sync"
)

func main(){
	r := ReadLinks()
	chanLinks := make(chan LinksType, 100000)
	chanOutputLinks := make(chan LinksType, 100000)
	var WG sync.WaitGroup
	for _,arr := range r {
		chanLinks <- arr
		WG.Add(1)
	}
	close(chanLinks)
	goroutines := 50
	url := "http://www.suagacollection.com"
	for i:= 0;i<goroutines;i++ {
		go TestLink(chanLinks,chanOutputLinks,url,&WG) 
	}
	WG.Wait()
	close(chanOutputLinks)
	var DoneLinks []LinksType
	for DoneLink := range chanOutputLinks{
		DoneLinks = append(DoneLinks,DoneLink)
	}
	results, count := socketMain(DoneLinks)
	for i:=0; i<count; i++ {
		fmt.Println("Client got: " + <-results)
	}
}

func checkerr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}