package search

import (
	"github.com/tasm20/go-vault-search/loops"
	"github.com/tasm20/go-vault-search/prints"
)

func ByFolder(paths loops.PathStruct, rootPath string, searchSlice []string) {
	foundCh := make(chan string)

	defer close(foundCh)
	for _, path := range folderSearchPaths(paths, rootPath) {
		go InPath(searchSlice, path, foundCh)
		prints.PrintFound(foundCh)
	}

	return
}

func folderSearchPaths(paths loops.PathStruct, rootPath string) []string {
	seen := make(map[string]bool)
	var res []string

	for _, path := range append([]string{rootPath}, append(paths.GetDirs(), paths.GetFiles()...)...) {
		if path == "" || seen[path] {
			continue
		}
		seen[path] = true
		res = append(res, path)
	}

	return res
}
