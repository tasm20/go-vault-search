package main

import (
	"flag"
	"fmt"
	"github.com/tasm20/go-vault-search/listSecrets"
	"github.com/tasm20/go-vault-search/loops"
	"github.com/tasm20/go-vault-search/prints"
	"github.com/tasm20/go-vault-search/search"
	"os"
	"strings"
)

const (
	version string = "1.0.0"
)

func main() {
	showVersion := flag.Bool("v", false, "version")
	vaultPath := flag.String("p", "kv/", "path to vault secret start searching")
	searchItems := flag.String("s", "", "what to search")
	searchKey := flag.Bool("k", false, "search secret key instead secret value")
	folderSearch := flag.Bool("cat", false, "search folder or file")
	listVaults := flag.Bool("l", false, "show only listSecrets of vaults in path")

	flag.Parse()

	if len(os.Args[1:]) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *showVersion {
		fmt.Println(version)
		return
	}

	if *searchItems == "" && !*listVaults {
		fmt.Println("Use", os.Args[0], "-s string to search")
		return
	}

	pathString := *vaultPath

	if !strings.Contains(pathString, "metadata") {
		pathString = strings.Replace(pathString, "kv", "kv/metadata", 1)
	}

	var searchSlice []string
	searchSlice = append(searchSlice, *searchItems)
	searchArgs := flag.Args()
	searchSlice = append(searchSlice, searchArgs...)

	if *listVaults {
		list, err := listSecrets.ListVault(pathString)
		if err != nil {
			prints.ErrorPrint(err)
			return
		}

		prints.PrintList(list, pathString)

		return
	}

	fmt.Println()
	paths, secrets := loops.GetList(pathString)

	if *folderSearch && paths.GetDirs() != nil {
		search.ByFolder(paths, searchSlice)
		return
	}

	var wasNotFound bool

	if secrets != nil {
		wasNotFound = search.InSecretsMap(secrets, searchSlice, *searchKey)
	} else {
		wasNotFound = search.InSecretsPath(paths, searchSlice, *searchKey)
	}

	if wasNotFound {
		prints.NotFound()
	}
	return
}
