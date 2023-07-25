package loops

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"github.com/tasm20/go-vault-search/listSecrets"
)

// func DirectoryLoop(clientVault *vault.Client, pathString string, dirsCh chan string) {
func DirectoryLoop(clientVault *vault.Client, pathString string) <-chan []string {
	list, err := listSecrets.ListVault(clientVault, pathString)
	if err != nil {
		fmt.Println(err)
	}

	dirsCh := make(chan []string)
	var pathSlice []string
	go func(dirsCh chan []string) {
		slice := innerNewPath(pathString, list)
		//for _, dir := range list {
		//	newPath := pathString + dir
		pathSlice = append(pathSlice, slice...)
		//}

		//chSlice := make(chan []string)
		//defer close(chSlice)
		//dirsCh <- pathSlice
		//dirsCh <- pathSlice
		tmpPathSlice := pathSlice
		defer close(dirsCh)
		pathCh := make(chan string)
		for {
			go innerPathSlice(tmpPathSlice, pathCh)
			tmpPathSlice, ok := innerVaultLoop(clientVault, pathCh)
			//slice = innerNewPath(pathString, tmpPathSlice)
			pathSlice = append(pathSlice, tmpPathSlice...)
			if !ok {
				close(pathCh)
				break
			}
		}
		dirsCh <- pathSlice

	}(dirsCh)

	return dirsCh
}

func innerPathSlice(pathSlice []string, pathCh chan string) {
	for _, path := range pathSlice {
		pathCh <- path
		//fmt.Println(path)
	}
}

func innerVaultLoop(clientVault *vault.Client, pathCh chan string) ([]string, bool) {
	path := <-pathCh
	list, _ := listSecrets.ListVault(clientVault, path)
	if len(list) == 0 {
		return nil, false
	}
	//var newPath []string
	//chSlice := make(chan []string)
	//newPath = append(newPath, list...)
	newPath := innerNewPath(path, list)
	//dirsCh <- path
	//chSlice <- newPath
	//close(chSlice)
	return newPath, true
}

func innerNewPath(pathString string, list []string) []string {
	var pathSlice []string
	for _, dir := range list {
		newPath := pathString + dir
		pathSlice = append(pathSlice, newPath)
	}

	return pathSlice
}
