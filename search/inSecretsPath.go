package search

import (
	"github.com/tasm20/go-vault-search/loops"
	"github.com/tasm20/go-vault-search/prints"
)

func InSecretsPath(paths loops.PathStruct, searchSlice []string, searchKey bool) bool {
	secretsDataCh := make(chan map[string]map[string][]byte)
	var wasNotFound bool

	for _, path := range paths.GetFiles() {
		secrets := loops.GetSecrets(path)
		go loops.SecretsLoop(secrets, secretsDataCh)
		found := InSecrets(secretsDataCh, searchSlice, searchKey)
		if len(found) > 0 {
			prints.MapsOfFoundSecrets(found)
			wasNotFound = false
		} else {
			wasNotFound = true
		}
	}
	defer close(secretsDataCh)
	return wasNotFound
}
