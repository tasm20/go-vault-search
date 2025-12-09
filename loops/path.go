package loops

import (
	"reflect"
	"strings"

	"github.com/hashicorp/vault/api"
	"github.com/tasm20/go-vault-search/prints"
)

func GetList(pathString string) (PathStruct, map[string]*api.KVSecret) {
	list, err := clientVault.Logical().List(pathString)
	if err != nil {
		prints.ErrorPrint(err)
	}

	var secrets map[string]*api.KVSecret
	if list != nil {
		pathStruct = pathLoop(list, pathString)
		return pathStruct, secrets
	}

	secrets, _ = GetSecrets(pathString)
	return pathStruct, secrets
}

func pathLoop(list *api.Secret, pathString string) PathStruct {
	var dirsCount []string
	inPathsCh := make(chan []string)
	defer close(inPathsCh)
	for _, lisMap := range list.Data {
		paths := reflect.ValueOf(lisMap)
		go innerPathLoop(paths, inPathsCh)
		outPathsCh := make(chan []string)
		go NewPath(pathString, inPathsCh, outPathsCh)
		dirs := <-outPathsCh
		if dirs != nil {
			dirsCount = append(dirsCount, dirs...)
		}
		close(outPathsCh)
	}

	if len(dirsCount) > 0 {
		for _, path := range dirsCount {
			if !strings.Contains(path, "metadata") {
				path = pathString + path
			}
			GetList(path)
		}
	}
	return pathStruct
}

func innerPathLoop(pathsIn reflect.Value, newPathCh chan []string) {
	var pathOut []string
	for i := 0; i < pathsIn.Len(); i++ {
		pathOut = append(pathOut, pathsIn.Index(i).Interface().(string))
	}

	newPathCh <- pathOut
}
