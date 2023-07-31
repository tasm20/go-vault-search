package initVault

import (
	vault "github.com/hashicorp/vault/api"
	"github.com/tasm20/go-vault-search/prints"
	"os"
)

func GetClient() *vault.Client {
	clientVault, err := InitVault()
	if err != nil {
		prints.ErrorPrint(err)
		os.Exit(2)
	}
	return clientVault
}
