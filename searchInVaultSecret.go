package main

import (
	"context"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"strings"
)

func searchInVaultSecret(client *vault.Client, searchItem string, searchKey bool) ([]string, error) {
	var found []string
	ctx := context.Background()
	dataNotFound := fmt.Errorf("\x1b[%dm%s\x1b[0m", 31, "DATA NOT FOUND")
	coloredSearchItem := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, searchItem)

	for _, secret := range secrets {

		vaultSecretPath := strings.Split(secret, "/")
		searchPath := vaultSecretPath[0]
		vaultSecret := strings.Replace(secret, "kv/metadata//", "", 1)
		vaultSecret = strings.Replace(vaultSecret, "//", "/", -1)

		if searchKey {
			if strings.Contains(secret, searchItem) {
				secret = strings.Replace(secret, "//", "/", -1)
				coloredResult := strings.Replace(secret, searchItem, coloredSearchItem, -1)
				found = append(found, coloredResult)
			}
		} else {
			check, err := client.KVv2(searchPath).Get(ctx, vaultSecret)
			if err != nil {
				return nil, err
			}

			for k, v := range check.Data {
				secretString := v.(string)
				if strings.Contains(secretString, searchItem) {
					coloredResult := strings.Replace(secretString, searchItem, coloredSearchItem, -1)
					result := searchPath + "/" + vaultSecret + " - " + k + " = " + coloredResult
					found = append(found, result)
				}
			}
		}
	}

	if len(found) > 0 {
		return found, nil
	}

	return nil, dataNotFound
}
