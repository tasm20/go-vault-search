package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	vault "github.com/hashicorp/vault/api"
)

func searchInVaultSecret(client *vault.Client) error {
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
			return err
		}

		for k, v := range check.Data {
			if rec, ok := v.(map[string]interface{}); ok {
				for kk, vv := range rec {
					secretKey := fmt.Sprintf("\u001B[%dm%s\u001B[0m", 31, kk)
					jStr, err := json.MarshalIndent(vv, "", " ")
					if err != nil {
						fmt.Println(err)
					}

					_, ok := searchInSlice(k, secretKey)
					if ok {
						result := searchPath + "/" + vaultSecret + " - " + secretKey + " = " + string(jStr)
						fmt.Println(result)
						found = append(found, result)
					}
				}
				continue
			}

			jStr, err := json.MarshalIndent(v, "", " ")
			if err != nil {
				fmt.Println(err)
			}

			secretKey := fmt.Sprintf("\u001B[%dm%s\u001B[0m", 31, k)
			secretValue := string(jStr)

			if *searchKey {
				_, ok := searchInSlice(k, secretKey)
				if ok {
					result := searchPath + "/" + vaultSecret + " - " + secretKey + " = " + secretValue
					fmt.Println(result)
					found = append(found, result)
				}

				continue
			}

			searchItem, ok := searchInSlice(k, secretValue)
			if ok {
				for _, val := range searchItem {
					coloredSecretValue := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, val)
					k = fmt.Sprintf("\u001B[%dm%s\u001B[0m", 33, k)
					coloredResult := strings.Replace(secretValue, val, coloredSecretValue, -1)
					result := searchPath + "/" + vaultSecret + " - " + k + " = " + coloredResult
					fmt.Println(result)
					found = append(found, result)
				}
			}
		}

	}

	if len(found) > 0 {
		foundCount = len(found)
		return nil
	}

	return dataNotFound
}

func searchInSlice(key, value string) (map[string]string, bool) {
	var found = make(map[string]string)

	for _, v := range searchSlice {
		searchItemColor := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, v)
		if strings.Contains(value, v) {
			_, ok := found[key]
			if ok {
				found[key] = strings.Replace(found[key], v, searchItemColor, -1)
			} else {
				value = strings.Replace(value, v, searchItemColor, -1)
				found[key] = value
			}
		}
	}

	if len(found) > 0 {
		return found, true
	}

	return nil, false
}
