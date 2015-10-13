package main
import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
	"net/http"
	"sync"
)

func ReadLinks() (List []LinksType){
	var Links []LinksType
	file, er := os.OpenFile("Links.json",os.O_RDONLY,777)
	checkerr(er)
	scanner := bufio.NewScanner(file)
	var a LinksType
	for scanner.Scan() {
		var s []byte
    	s = scanner.Bytes()
		json.Unmarshal(s, &a)
		Links = append(Links, a)
	}
	return Links
}
func TestLink(List ,Output chan LinksType,url string,WG *sync.WaitGroup)  {
	for {
		link,more := <- List
		if more == false {
			break
			
		}
		fmt.Println(url+link.Link)
		response, err := http.Get(url+link.Link)
		checkerr(err)
		if response.StatusCode == 200{
			if response.Body != nil{
				link.Link = url+link.Link
				Output <- link
				fmt.Println(link)
				response.Body.Close()
			}
		}
		WG.Done()
	}
}
type LinksType struct {
	Link string
	Cve string
	
	Description string
	
	PatchLink string
	
	Danger int64 //from 1 to 10k for exmpl
}