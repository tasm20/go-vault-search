package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	vault "github.com/hashicorp/vault/api"
)

const (
	version string = "0.3.5"
)

// TODO: do search by folder/file - done, but still need to work on it

var (
	searchKey    *bool
	searchSlice  []string
	foundCount   int
	folderSearch *bool
	folderFound  []string
)

func checkFolder(client *vault.Client) bool {
	err := searchInVaultSecret(client)
	return err == nil
}

func main() {
	var path []string
	var dataNotFound error

	showVersion := flag.Bool("v", false, "version")
	vaultPath := flag.String("p", "kv/", "path to vault secret start searching")
	searchItem := flag.String("s", "", "what to search")
	searchKey = flag.Bool("k", false, "search secret key instead secret value")
	folderSearch = flag.Bool("cat", false, "search folder or file")
	listVaults := flag.Bool("l", false, "show only list of vaults in path")

	flag.Parse()

	if len(os.Args[1:]) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if *searchItem == "" && !*listVaults {
		fmt.Println("Use", os.Args[0], "-s string to search")
		os.Exit(1)
	}

	client, err := initVault()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	pathString := *vaultPath

	if !strings.Contains(pathString, "metadata") {
		pathString = strings.Replace(pathString, "kv", "kv/metadata", 1)
	}

	searchSlice = append(searchSlice, *searchItem)
	searchArgs := flag.Args()
	searchSlice = append(searchSlice, searchArgs...)
	secrets = append(secrets, pathString)

	fmt.Println()

	pathToSecret := checkFolder(client)

	if !pathToSecret {
		secrets = nil
		pathString = pathString + "/"
		path = append(path, pathString)

		err = getListVault(client, path, *listVaults)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}

	if !*folderSearch && !*listVaults {
		dataNotFound = searchInVaultSecret(client)
	}

	if *folderSearch {
		if len(folderFound) > 0 {
			fmt.Printf("folder/file %s was found in:\n", searchSlice)

			for _, path := range folderFound {
				fmt.Println(path)
			}
		} else {
			fmt.Printf("folder/file %s was not found in %s\n", searchSlice, pathString)
		}

		foundCount = len(folderFound)
	}

	if *listVaults {
		pathString = strings.Replace(pathString, "metadata/", "", 1)
		fmt.Printf("found dirs in %s:\n", pathString)
		for _, secret := range secrets {
			fmt.Printf("\t%s\n", secret)
		}
		fmt.Printf("\nfound %d \n", len(secrets))
		os.Exit(0)
	}

	if dataNotFound != nil {
		fmt.Println(dataNotFound)
	}

	fmt.Printf("\nfound %d\n", foundCount)
}
