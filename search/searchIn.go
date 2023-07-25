package search

import (
	"fmt"
	"strings"
)

func SearchInPath(searchSlice []string, dirCh chan string) {
	dir := <-dirCh
	for _, searchItem := range searchSlice {
		if strings.Contains(dir, searchItem) {
			colorFoundItem := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, searchItem)
			colorResult := strings.Replace(dir, searchItem, colorFoundItem, -1)
			fmt.Println(colorResult)
		}
	}
}
