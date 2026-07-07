package search

import (
	"fmt"
	"strings"

	"github.com/tasm20/go-vault-search/loops"
)

func InPath(searchSlice []string, pathItem string, found chan string) {
	res := loops.RemoveExtrasInPath(pathItem)
	colorResult := res
	matched := false

	for _, searchItem := range searchSlice {
		if strings.Contains(res, searchItem) {
			FoundCount++
			colorFoundItem := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, searchItem)
			colorResult = strings.Replace(colorResult, searchItem, colorFoundItem, -1)
			matched = true
		}
	}

	if !matched {
		found <- ""
		return
	}

	found <- colorResult
}
