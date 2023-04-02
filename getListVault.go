package main

import (
	"errors"
	vault "github.com/hashicorp/vault/api"
	"reflect"
	"strings"
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
						pathString := path + "/" + searchPathMap.Index(i).Interface().(string)
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
