# stemp
Go simple string templates

## Installation
```bash
go get github.com/nokusukun/stemp
```

## Usage
```go
package main

import (
	"fmt"
	"strings"

	"github.com/nokusukun/stemp"
)

type Client struct {
	Name    string
	BotName string
}

func main() {
	myClient := Client{Name: "Noku", BotName: "RegBot"}
	resultStruct := stemp.CompileStruct("Hello [{Name:j=c,w=10}], I'm {BotName}. Nice to meet you {asdasds}. I'm accessing the values from a struct!", myClient)
    fmt.Println(resultStruct)
    

	data := []string{
		"Mays,123 st. 456 Lane",
		"Harold,523 st. Sunset Overdrive",
	}
    result := stemp.CompileJSON("   |{col1:j=c,w=10}|{col2:j=c,w=40}|\n---------------------------------------------------\n", `{"col1": "Name", "col2": "Address"}`)
	for idx, d := range data {
		s := strings.Split(d, ",")

		result += stemp.Compile("{idx:w=3}|{name:j=c,w=10}|{add:j=c,w=40}|\n", map[string]interface{}{"idx": idx, "name": s[0], "add": s[1]})
	}
	fmt.Println(result)
}

```

### Result
```
Hello [   Noku   ], I'm RegBot. Nice to meet you {asdasds}. I'm accessing the values from a struct!
   |   Name   |                Address                 |
---------------------------------------------------
0  |   Mays   |            123 st. 456 Lane            |
1  |  Harold  |        523 st. Sunset Overdrive        |
```

## Parameters
Parameters can be passed to the template which modifies the final output.
These are the current available parameters
* `j=[l/c/r]`
    * Justifies the templated string to `l`eft, `c`enter or `r`ight.
* `w=int`
    * Sets the minimum width of the string.

## Docs
https://godoc.org/github.com/nokusukun/stemp
