package search

import (
	"fmt"
	"github.com/tasm20/go-vault-search/loops"
	"strings"
)

func InPath(searchSlice []string, pathItem string, found chan string) {
	for _, searchItem := range searchSlice {
		if strings.Contains(pathItem, searchItem) {
			res := loops.RemoveExtrasInPath(pathItem)
			colorFoundItem := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, searchItem)
			colorResult := strings.Replace(res, searchItem, colorFoundItem, -1)
			found <- colorResult
		} else {
			found <- ""
		}
	}
}
