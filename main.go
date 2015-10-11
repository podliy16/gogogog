package main

import (
	"fmt"
)
func main(){
	r := ReadLinks()
	//fmt.Println(r)
	url := "https://google.com"
	t := TestLink(r,url) 
	fmt.Println(t)

}
