package main

import (
	"fmt"
	"github.com/jeffail/gabs"
)

type Main struct{
	items []Item
	parametr Parametr
}

type Item struct{
	parametr string
	typeItem string
	title string
}

type Parametr struct{
	typeParametr string
	parametr string
}

func parseAnswear(answear string){
	m := Main{}
	p := Parametr{}
	mainItem := make([]Item,0)
	jsonParsed,err := gabs.ParseJSON([]byte(answear))
	if err != nil {
		return 	
	}
	p.parametr = jsonParsed.Search("main","param","param").Data().([]interface{})[0].(string)
	p.typeParametr = jsonParsed.Search("main","param","t").Data().([]interface{})[0].(string)
	
	//childreItem - items for key "i"
	
	childrenItem,_ := jsonParsed.Search("main","i").Children()

	//children - array with main key "i"

	children := childrenItem[0]

	//resultChildren - final array with our keys "p","t","ti"
	
	resultChildren,_ := children.Children()
	var item Item
	for _,child := range resultChildren{
		item.parametr = child.Data().(map[string]interface{})["p"].(string)	
		item.typeItem = child.Data().(map[string]interface{})["t"].(string)
		item.title = child.Data().(map[string]interface{})["ti"].(string)
		mainItem = append(mainItem, item)
	}
	m.items = mainItem
	m.parametr = p
	fmt.Println(m)
}