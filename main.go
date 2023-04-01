package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	version string = "0.1.0"
)

func main() {
	startTime := time.Now()
	duration := time.Since(startTime)
	var path []string

	showVersion := flag.Bool("v", false, "version")
	vaultPath := flag.String("p", "kv/metadata/", "path to vault secret start searching")
	searchItem := flag.String("s", "", "what to search")
	searchKey := flag.Bool("k", false, "search secret key instead secret value")

	flag.Parse()

	if len(os.Args[1:]) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if *searchItem == "" {
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

	err = getListVault(client, path)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	found, err := searchInVaultSecret(client, *searchItem, *searchKey)
	if err != nil {
		fmt.Println(err)
	}

	foundCount := len(found)

	fmt.Println()

	for _, v := range found {
		fmt.Println(v)
	}

	fmt.Printf("\nfound %d in %s\n", foundCount, duration)
}
