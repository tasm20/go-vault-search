package listSecrets

import (
	"errors"
	"reflect"
)

func ListVault(path string) ([]string, error) {

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
		}
	}

	return listMap, err
}
