package main
import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
	"net/http"
)

func ReadLinks() (List []LinksType){
	var Links []LinksType
	file, er := os.OpenFile("Links.json",os.O_RDONLY,777)
	if er != nil{
		
	}
	scanner := bufio.NewScanner(file)
	var a LinksType
	for scanner.Scan() {
		var s []byte
    	s = scanner.Bytes()
		json.Unmarshal(s, &a)
		Links = append(Links, a)
	}
	fmt.Println(Links)
	return Links
}
func TestLink(List []LinksType,url string)  []string{
	var SliceLinksFound []string
	for _, link := range List{
		response, err := http.Get(url+link.Link)
		if err != nil{
			fmt.Println(err)
		}
		if response.StatusCode == 404{
			fmt.Println(url+link.Link+" Not Found")
			response.Body.Close()
			continue 
		}
		SliceLinksFound = append(SliceLinksFound,url+link.Link)
		defer response.Body.Close()
	}
	return SliceLinksFound
}
type LinksType struct {
	Link string
}