package stemp

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var findWords = regexp.MustCompile(`{\w+}`)

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
		target, exists := maps[sourcetxt]
		if exists {
			template = strings.ReplaceAll(template, source, fmt.Sprintf("%v", target))
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
		target := reflect.ValueOf(maps).FieldByName(sourcetxt)
		if target.IsValid() {
			template = strings.ReplaceAll(template, source, fmt.Sprintf("%v", target))
		} else {
			hasError = true
		}
	}
	if hasError {
		return template, fmt.Errorf("There are some invalid fields")
	}
	return template, nil
}
