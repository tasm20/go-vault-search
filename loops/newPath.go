package loops

import "strings"

func NewPath(pathString string, list []string) []string {
	var dirs []string
	for _, dir := range list {
		newPath := pathString + dir

		if strings.HasSuffix(newPath, "/") {
			pathStruct.dirs = append(pathStruct.dirs, newPath)
			dirs = append(dirs, newPath)
		} else {
			pathStruct.files = append(pathStruct.files, newPath)
		}
	}
	return dirs
}
