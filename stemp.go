package stemp

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Parameters struct {
	Source  string
	Justify string
	Width   int
}

func (p *Parameters) BuildString(target string) string {
	if p.Width <= len(target) {
		return target
	}
	wsl := p.Width - len(target)

	if p.Justify == "" || p.Justify == "l" {
		target = target + whitespace(wsl)
	} else if p.Justify == "c" {
		target = whitespace(wsl/2) + target + whitespace(wsl/2)
	} else if p.Justify == "r" {
		target = whitespace(wsl) + target
	}

	if len(target) < p.Width {
		target += whitespace(p.Width - len(target))
	}

	return target
}

func whitespace(length int) string {
	ws := ""
	for i := 0; i < length; i++ {
		ws += " "
	}
	return ws
}

func getParameters(stxt string) *Parameters {
	finder := regexp.MustCompile(`[a-z]=[a-zA-Z0-9]+`)
	elems := strings.Split(stxt, ":")
	params := &Parameters{}

	if len(elems) == 0 {
		return params
	}

	params.Source = elems[0]
	if len(elems) > 1 {
		stringparams := finder.FindAllString(elems[1], -1)
		for _, param := range stringparams {
			paramelem := strings.Split(param, "=")
			if len(paramelem) != 2 {
				continue
			}
			prop := paramelem[0]
			val := paramelem[1]
			switch prop {
			case "j":
				params.Justify = val
			case "w":
				params.Width, _ = strconv.Atoi(val)
			}
		}
	}
	return params
}

var findWords = regexp.MustCompile(`{\w+(:[a-zA-Z0-9,= ]+)?}`)

// CompileJSON compiles a template string given a JSON string.
func CompileJSON(template, jsonstring string) string {
	template, _ = CompileJSONStrict(template, jsonstring)
	return template
}

// CompileJSONStrict compiles a template string given a JSON string.
// Returns an error and the partial compiled string when it finds an invalid target.
func CompileJSONStrict(template, jsonstring string) (string, error) {
	maps := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonstring), &maps)
	if err != nil {
		return template, err
	}

	template, err = CompileStrict(template, maps)
	if err != nil {
		return template, err
	}

	return template, nil
}

// Compile - Compiles a string template given a map of [string]interface{}.
func Compile(template string, maps map[string]interface{}) string {
	template, _ = CompileStrict(template, maps)
	return template
}

// CompileStrict - Compiles a string template given a map of [string]interface{}.
// Returns an error and the partial compiled string when it finds an invalid target.
func CompileStrict(template string, maps map[string]interface{}) (string, error) {
	words := findWords.FindAllString(template, -1)
	hasError := false
	for _, source := range words {
		sourcetxt := source[1 : len(source)-1]
		params := getParameters(sourcetxt)
		target, exists := maps[params.Source]
		if exists {
			template = strings.ReplaceAll(template, source, params.BuildString(fmt.Sprintf("%v", target)))
		}
	}
	if hasError {
		return template, fmt.Errorf("There are some invalid fields")
	}
	return template, nil
}

// CompileStruct - Compiles a string template given a struct.
func CompileStruct(template string, maps interface{}) string {
	template, _ = CompileStructStrict(template, maps)
	return template
}

// CompileStructStrict - Compiles a string template given a struct.
// Returns an error and the partial compiled string when it finds an invalid target.
func CompileStructStrict(template string, maps interface{}) (string, error) {
	words := findWords.FindAllString(template, -1)
	hasError := false
	for _, source := range words {
		sourcetxt := source[1 : len(source)-1]
		params := getParameters(sourcetxt)
		target := reflect.ValueOf(maps).FieldByName(params.Source)
		if target.IsValid() {
			template = strings.ReplaceAll(template, source, params.BuildString(fmt.Sprintf("%v", target)))
		} else {
			hasError = true
		}
	}
	if hasError {
		return template, fmt.Errorf("There are some invalid fields")
	}
	return template, nil
}
