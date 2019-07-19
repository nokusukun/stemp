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
	"github.com/nokusukun/stemp"
)

type Client struct {
	Name string
	BotName string
}


func main() {
    myClient := stemp.Client{Name: "Noku", BotName: "RegBot"}
    myMap := map[string]interface{}{"name": "Nokui", "botname": "Somalia"}

    resultStruct := stemp.CompileStruct("Hello {Name}, I'm {BotName}. Nice to meet you {asdasds}. I'm accessing the values from a struct!", myClient)
    
    result := stemp.Compile("Hello {name}, I'm {botname}. I'm accessing the values from a map!", myMap)
    
    fmt.Println(resultStruct)
    fmt.Println(result)
}
```

### Result
```
Hello Noku, I'm RegBot. Nice to meet you {asdasds}. I'm accessing the values from a struct!
Hello Nokui, I'm Somalia. I'm accessing the values from a map!
```

## Docs
https://godoc.org/github.com/nokusukun/stemp