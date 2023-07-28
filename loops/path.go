package loops

import (
	vault "github.com/hashicorp/vault/api"
	"github.com/tasm20/go-vault-search/prints"
	"reflect"
)

func PathLoop(clientVault *vault.Client, pathString string) PathStruct {
	var dirsCount []string
	list, err := clientVault.Logical().List(pathString)
	if err != nil {
		prints.ErrorPrint(err)
	}
	for _, lisMap := range list.Data {
		paths := reflect.ValueOf(lisMap)
		newPaths := innerPathLoop(paths)
		dirs := NewPath(pathString, newPaths)
		dirsCount = append(dirsCount, dirs...)
	}

	if len(dirsCount) > 0 {
		for _, t := range dirsCount {
			PathLoop(clientVault, t)
		}
	}
	return pathStruct
}

func innerPathLoop(pathsIn reflect.Value) []string {
	var pathOut []string
	for i := 0; i < pathsIn.Len(); i++ {
		pathOut = append(pathOut, pathsIn.Index(i).Interface().(string))
	}

	return pathOut
}
