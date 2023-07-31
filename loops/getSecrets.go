package loops

import (
	"context"
	"github.com/hashicorp/vault/api"
	"github.com/tasm20/go-vault-search/prints"
	"strings"
)

func GetSecrets(val string) map[string]*api.KVSecret {

	pathSlice := strings.Split(val, "/")
	path := pathSlice[0]
	path = strings.Replace(path, pathSlice[len(pathSlice)-1], "", 1)
	secret := strings.Replace(val, "kv/metadata/", "", 1)
	secret = strings.Replace(secret, "//", "/", -1)

	ctx := context.Background()
	secretsDATA, err := clientVault.KVv2(path).Get(ctx, secret)
	if err != nil {
		prints.ErrorPrint(err)
	}
	secrets := map[string]*api.KVSecret{
		secret: secretsDATA,
	}
	return secrets
}
