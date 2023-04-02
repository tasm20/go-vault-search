package main

import (
	"context"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"strings"
)

func searchInVaultSecret(client *vault.Client, searchItem string, searchKey bool) (int, error) {
	var found []string
	ctx := context.Background()
	dataNotFound := fmt.Errorf("\x1b[%dm%s\x1b[0m", 31, "DATA NOT FOUND")
	coloredSearchItem := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, searchItem)

	for _, secret := range secrets {

		vaultSecretPath := strings.Split(secret, "/")
		searchPath := vaultSecretPath[0]
		vaultSecret := strings.Replace(secret, "kv/metadata/", "", 1)
		vaultSecret = strings.Replace(vaultSecret, "//", "/", -1)

		check, err := client.KVv2(searchPath).Get(ctx, vaultSecret)
		if err != nil {
			return 0, err
		}

		for k, v := range check.Data {
			if searchKey {
				if strings.Contains(k, searchItem) {
					k = fmt.Sprintf("\u001B[%dm%s\u001B[0m", 33, k)
					result := searchPath + "/" + vaultSecret + " - " + k
					result = strings.Replace(result, "//", "/", -1)
					fmt.Println(result)
					found = append(found, result)

					continue
				}
			}

			secretString := fmt.Sprintf("%v", v)
			if strings.Contains(secretString, searchItem) {
				k = fmt.Sprintf("\u001B[%dm%s\u001B[0m", 33, k)
				coloredResult := strings.Replace(secretString, searchItem, coloredSearchItem, -1)
				result := searchPath + "/" + vaultSecret + " - " + k + " = " + coloredResult
				result = strings.Replace(result, "//", "/", -1)
				fmt.Println(result)
				found = append(found, result)
			}
		}

	}

	if len(found) > 0 {
		return len(found), nil
	}

	return 0, dataNotFound
}
