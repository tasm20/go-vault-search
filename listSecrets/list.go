package listSecrets

import (
	"errors"
	"github.com/tasm20/go-vault-search/loops"
	"github.com/tasm20/go-vault-search/prints"
	"reflect"
)

func ListVault(path string) ([]string, error) {

	listSecrets, err := loops.GetSecrets(path)
	if err == nil {
		for key, val := range listSecrets {
			prints.PrintAllinPath(key, val.Data)
		}
		return nil, nil
	}

	path = AddMeta(path)
	listSecret, err := clientVault.Logical().List(path)
	if err != nil {
		return nil, err
	}

	if listSecret == nil {
		err := errors.New("something wrong in search address")
		return nil, err
	}

	var listMap []string
	for _, list := range listSecret.Data {
		switch reflect.TypeOf(list).Kind() {
		case reflect.Slice:
			searchPathMap := reflect.ValueOf(list)

			for i := 0; i < searchPathMap.Len(); i++ {
				pathString := searchPathMap.Index(i).Interface().(string)
				listMap = append(listMap, pathString)
			}
		default:
			panic("unhandled default case")
		}
	}

	return listMap, err
}
