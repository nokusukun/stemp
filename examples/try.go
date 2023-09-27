package main

import (
	"fmt"
	"github.com/nokusukun/stemp"
	"strings"
)

type Client struct {
	Name    string
	BotName string
}

func main() {
	myClient := Client{Name: "Noku", BotName: "RegBot"}
	//myMap := map[string]interface{}{"col1": "Name", "col2": "Address"}
	data := []string{
		"Mays,123 st. 456 Lane",
		"Harold,523 st. Sunset Overdrive",
	}

	resultStruct := stemp.CompileStruct("{_:w=3}Hello [{Name:j=c,w=10,f=-}], I'm {BotName}. Nice to meet you {asdasds}. I'm accessing the values from a struct!", myClient)

	result := stemp.CompileJSON("{_:w=3}|{col1:j=c,w=10}|{col2:j=c,w=40}|\n{_:w=60,f=-}\n", `{"col1": "Name", "col2": "Address"}`)
	for idx, d := range data {
		s := strings.Split(d, ",")

		result += stemp.Compile("{idx:w=3}|{name:j=c,w=10}|{add:j=c,w=40}|\n", map[string]interface{}{"idx": idx, "name": s[0], "add": s[1]})
	}
	fmt.Println(resultStruct)
	fmt.Println(result)

	fmt.Println(stemp.Inline("Hello {0}, I'm {1}. Nice to meet you {2}. I'm accessing the values from a struct! {0:w=10,j=c,f=*}", "Noku", "RegBot", "asdasds"))
}
