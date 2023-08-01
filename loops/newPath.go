package loops

import (
	"strings"
)

func NewPath(pathString string, inPathsCh, outPathsCh chan []string) {
	var dirs []string
	for _, dir := range <-inPathsCh {
		if !strings.HasSuffix(pathString, "/") {
			pathString += "/"
		}
		newPath := pathString + dir

		if strings.HasSuffix(newPath, "/") {
			pathStruct.dirs = append(pathStruct.dirs, newPath)
			dirs = append(dirs, newPath)
		} else {
			pathStruct.files = append(pathStruct.files, newPath)
		}
	}
	outPathsCh <- dirs
}
