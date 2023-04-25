package main

import (
	"flag"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"os"
	"strings"
)

const (
	version string = "0.3.3"
)

// TODO: do search by folder/file

var (
	searchKey   *bool
	searchSlice []string
	foundCount  int
)

func checkFolder(client *vault.Client) bool {
	err := searchInVaultSecret(client)
	if err != nil {
		return false
	}

	return true
}

func main() {
	var path []string

	showVersion := flag.Bool("v", false, "version")
	vaultPath := flag.String("p", "kv/", "path to vault secret start searching")
	searchItem := flag.String("s", "", "what to search")
	searchKey = flag.Bool("k", false, "search secret key instead secret value")
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

		err = searchInVaultSecret(client)
		if err != nil {
			fmt.Println(err)
		}
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

	fmt.Printf("\nfound %d\n", foundCount)
}
