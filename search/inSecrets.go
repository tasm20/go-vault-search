package search

import (
	"fmt"
	"strings"
)

func InSecrets(secretsDataCh chan map[string]map[string][]byte, searchSlice []string, keySearch bool) map[string]map[string]string {
	isKeySearch = keySearch
	searchItems = searchSlice
	secretsDATA := <-secretsDataCh
	foundMap := make(map[string]map[string]string)
	innerCh := make(chan map[string]string)
	for secretKey := range secretsDATA {
		go innerSecretValueLoop(secretsDATA[secretKey], innerCh)
		innerMap := <-innerCh
		if len(innerMap) != 0 {
			foundMap[secretKey] = innerMap
		}
	}
	close(innerCh)
	return foundMap
}

func innerSecretValueLoop(secretsMap map[string][]byte, innerCh chan map[string]string) {
	outMassive := make(map[string]string)
	for key, value := range secretsMap {
		valueStr := string(value)
		if isKeySearch {
			key = searchSecrets(key)
			if key != "" {
				outMassive[key] = valueStr
			}
		} else {
			valueStr = searchSecrets(valueStr)
			if valueStr != "" {
				outMassive[key] = valueStr
			}
		}
	}
	innerCh <- outMassive
}

func searchSecrets(item string) string {
	var found string
	for _, searchItem := range searchItems {
		if strings.Contains(item, searchItem) {
			colorFoundItem := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, searchItem)
			colorResult := strings.Replace(item, searchItem, colorFoundItem, -1)
			found = colorResult
		}
	}
	return found
}