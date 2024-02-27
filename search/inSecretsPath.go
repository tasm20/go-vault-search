package search

import (
	"github.com/tasm20/go-vault-search/loops"
	"github.com/tasm20/go-vault-search/prints"
)

func InSecretsPath(paths loops.PathStruct, searchSlice []string, searchKey bool) {
	secretsDataCh := make(chan map[string]map[string][]byte)

	for _, path := range paths.GetFiles() {
		secrets, _ := loops.GetSecrets(path)
		go loops.SecretsLoop(secrets, secretsDataCh)
		found := InSecrets(secretsDataCh, searchSlice, searchKey)

		if FoundCount > 0 {
			prints.MapsOfFoundSecrets(found)
		}
	}
	defer close(secretsDataCh)
}
