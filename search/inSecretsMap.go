package search

import (
	"github.com/hashicorp/vault/api"
	"github.com/tasm20/go-vault-search/loops"
	"github.com/tasm20/go-vault-search/prints"
)

func InSecretsMap(secrets map[string]*api.KVSecret, searchSlice []string, searchKey bool) {
	secretsDataCh := make(chan map[string]map[string][]byte)

	go loops.SecretsLoop(secrets, secretsDataCh)
	found := InSecrets(secretsDataCh, searchSlice, searchKey)

	if FoundCount > 0 {
		prints.MapsOfFoundSecrets(found)
	}
	defer close(secretsDataCh)
}
