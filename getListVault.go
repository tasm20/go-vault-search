package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	vault "github.com/hashicorp/vault/api"
)

var secrets []string

func getListVault(client *vault.Client, vaultPath []string, listVaults bool) error {
	var dirs []string

	for _, path := range vaultPath {
		vaultList, err := client.Logical().List(path)
		if err != nil {
			return err
		}

		if vaultList == nil {
			err := errors.New("something wrong in search address")
			return err
		}

		if listVaults {
			for _, listMap := range vaultList.Data {
				switch reflect.TypeOf(listMap).Kind() {
				case reflect.Slice:
					searchPathMap := reflect.ValueOf(listMap)

					for i := 0; i < searchPathMap.Len(); i++ {
						pathString := searchPathMap.Index(i).Interface().(string)
						secrets = append(secrets, pathString)
					}
				}
			}
			return nil
		}

		secretsList, dirsList := loop(path, vaultList)

		secrets = append(secrets, secretsList...)
		dirs = append(dirs, dirsList...)
	}

	if len(folderFound) > 0 {
		return nil
	}

	if len(dirs) > 0 {
		_ = getListVault(client, dirs, false)
	}

	return nil
}

func loop(vaultPath string, vaultList *vault.Secret) ([]string, []string) {
	var secrets, dirs []string
	for _, listMap := range vaultList.Data {

		switch reflect.TypeOf(listMap).Kind() {
		case reflect.Slice:
			searchPathMap := reflect.ValueOf(listMap)

			for i := 0; i < searchPathMap.Len(); i++ {
				currentSearchPath := vaultPath + "/" + searchPathMap.Index(i).Interface().(string)

				if *folderSearch {
					for _, searchItem := range searchSlice {
						if strings.Contains(currentSearchPath, searchItem) {
							color := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, searchItem)
							coloredRes := strings.Replace(currentSearchPath, "metadata/", "", 1)
							coloredRes = strings.Replace(coloredRes, "//", "/", -1)
							coloredRes = strings.Replace(coloredRes, searchItem, color, -1)
							folderFound = append(folderFound, coloredRes)
						}
					}
				}

				if strings.HasSuffix(currentSearchPath, "/") {
					dirs = append(dirs, currentSearchPath)
				} else {
					secrets = append(secrets, currentSearchPath)
				}
			}
		}
	}

	return secrets, dirs
}
