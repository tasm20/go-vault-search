package search

import (
	"fmt"
	"strings"
)

func InPath(searchSlice []string, pathItem string, found chan string) {
	for _, searchItem := range searchSlice {
		if strings.Contains(pathItem, searchItem) {
			res := strings.Replace(pathItem, "kv/metadata", "kv/", 1)
			res = innerCutDoubleSlash(res)
			colorFoundItem := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, searchItem)
			colorResult := strings.Replace(res, searchItem, colorFoundItem, -1)
			found <- colorResult
		} else {
			found <- ""
		}
	}
}

func innerCutDoubleSlash(str string) string {
	res := strings.Replace(str, "//", "/", -1)
	if strings.Contains(res, "//") {
		return innerCutDoubleSlash(res)
	}
	return res
}
