package main

import (
	"fmt"
//	"sync"
)

//func main(){
//	r := ReadLinks()
//	chanLinks := make(chan LinksType, 5000)
//	chanOutputLinks := make(chan LinksType, 5000)
//	var WG sync.WaitGroup
//	for _,arr := range r {
//		chanLinks <- arr
//	}
//	close(chanLinks)
//	goroutines := 50
//	url := "https://google.com"
//	for i:= 0;i<goroutines;i++  {
//		go TestLink(chanLinks,chanOutputLinks,url,&WG) 
//	}
//	WG.Wait()
//	close(chanOutputLinks)
//	var DoneLinks []LinksType
//	for DoneLink := range chanOutputLinks{
//		DoneLinks = append(DoneLinks,DoneLink)
//	}
//	fmt.Println(DoneLinks)
//}

func checkerr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}