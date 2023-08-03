package search

import (
	"github.com/tasm20/go-vault-search/loops"
	"github.com/tasm20/go-vault-search/prints"
)

func ByFolder(paths loops.PathStruct, searchSlice []string) {
	foundCh := make(chan string)

	defer close(foundCh)
	for _, path := range paths.GetDirs() {
		go InPath(searchSlice, path, foundCh)
		prints.PrintFound(foundCh)
	}

	return
}
