package initVault

import (
	"errors"
	"os"

	vault "github.com/hashicorp/vault/api"
)

func InitVault() (*vault.Client, error) {
	config := vault.DefaultConfig()
	vaultAddress := os.Getenv("VAULT_ADDR")
	vaultToken := os.Getenv("VAULT_TOKEN")

	if vaultAddress != "" {
		config.Address = vaultAddress
	} else {
		errorString := errors.New("export VAULT_ADDR='vault address'" +
			"export VAULT_TOKEN=<VAULT_TOKEN>")
		return nil, errorString
	}

	client, err := vault.NewClient(config)
	if err != nil {
		return nil, err
	}

	client.SetToken(vaultToken)

	return client, nil
}
