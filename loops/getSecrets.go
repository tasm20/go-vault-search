package loops

import (
	"context"
	"github.com/hashicorp/vault/api"
	"strings"
)

func GetSecrets(val string) (map[string]*api.KVSecret, error) {

	pathSlice := strings.Split(val, "/")
	path := pathSlice[0]
	path = strings.Replace(path, pathSlice[len(pathSlice)-1], "", 1)
	secret := strings.Replace(val, "kv/metadata/", "", 1)
	if strings.HasPrefix(secret, "kv") {
		secret = strings.Replace(secret, "kv", "", 1)
	}
	secret = strings.Replace(secret, "//", "/", -1)

	ctx := context.Background()
	secretsDATA, err := clientVault.KVv2(path).Get(ctx, secret)
	if err != nil {
		//prints.ErrorPrint(err)
		return nil, err
	}
	secrets := map[string]*api.KVSecret{
		secret: secretsDATA,
	}
	return secrets, nil
}
