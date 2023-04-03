package main

import (
	"context"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"strings"
)

func searchInVaultSecret(client *vault.Client) (int, error) {
	var found []string
	ctx := context.Background()
	dataNotFound := fmt.Errorf("\x1b[%dm%s\x1b[0m", 31, "DATA NOT FOUND")

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
			secretKey := fmt.Sprintf("\u001B[%dm%s\u001B[0m", 31, k)
			secretValue := fmt.Sprintf("%v", v)

			if *searchKey {
				_, ok := searchInSlice(secretKey)
				if ok {
					result := searchPath + "/" + vaultSecret + " - " + secretKey + " = " + secretValue
					result = strings.Replace(result, "//", "/", 1)
					fmt.Println(result)
					found = append(found, result)
				}

				continue
			}

			searchItem, ok := searchInSlice(secretValue)
			if ok {
				coloredSecretValue := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, searchItem)
				k = fmt.Sprintf("\u001B[%dm%s\u001B[0m", 33, k)
				coloredResult := strings.Replace(secretValue, searchItem, coloredSecretValue, -1)
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

func searchInSlice(key string) (string, bool) {
	for _, v := range searchSlice {
		if strings.Contains(key, v) {
			return v, true
		}
	}
	return "no", false
}
