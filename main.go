package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	version string = "0.2.1"
)

func main() {
	var path []string

	showVersion := flag.Bool("v", false, "version")
	vaultPath := flag.String("p", "kv/", "path to vault secret start searching")
	searchItem := flag.String("s", "", "what to search")
	searchKey := flag.Bool("k", false, "search secret key instead secret value")
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
	if !strings.Contains(*vaultPath, "metadata") {
		pathString = strings.Replace(*vaultPath, "kv", "kv/metadata", 1)
	}

	path = append(path, pathString)

	err = getListVault(client, path, *listVaults)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println()

	if *listVaults {
		fmt.Printf("found dirs in %s:\n", *vaultPath)
		for _, secret := range secrets {
			fmt.Printf("\t%s\n", secret)
		}
		fmt.Printf("\nfound %d keys\n", len(secrets))
		os.Exit(0)
	}

	found, err := searchInVaultSecret(client, *searchItem, *searchKey)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nfound %d\n", found)
}
