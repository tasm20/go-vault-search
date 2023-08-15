package loops

import (
	"encoding/json"
	"github.com/hashicorp/vault/api"
	"github.com/tasm20/go-vault-search/prints"
)

func kvSecretsToMap(secretsList *api.KVSecret) map[string]interface{} {
	res := make(map[string]interface{})
	for keyList, value := range secretsList.Data {
		if mapSecrets, ok := value.(map[string]interface{}); ok {
			for keySecrets, valueSecrets := range mapSecrets {
				res[keySecrets] = valueSecrets
			}
		} else {
			res[keyList] = value
		}
	}
	return res
}

func SecretsLoop(secretsList map[string]*api.KVSecret, secretsDataCh chan map[string]map[string][]byte) {
	var secretDATA = make(map[string]map[string][]byte)
	for key, values := range secretsList {
		valueMapCh := make(chan map[string][]byte)
		valuesMap := kvSecretsToMap(values)
		go innerSecretValueLoop(valuesMap, valueMapCh)

		secretDATA[key] = <-valueMapCh
	}

	secretsDataCh <- secretDATA
}

func innerSecretValueLoop(values map[string]interface{}, valueMapCh chan map[string][]byte) {
	var newMap = make(map[string][]byte)

	for key, value := range values {
		var vJson []byte
		valueCh := make(chan interface{})
		valuesInterface, ok := value.(map[string]interface{})
		if ok {
			go innerInterfaceLoop(valuesInterface, valueCh)
			vJson = innerToJson(<-valueCh)
		} else {
			vJson = innerToJson(value)
		}
		newMap[key] = vJson
	}

	valueMapCh <- newMap
}

func innerInterfaceLoop(values map[string]interface{}, valueCh chan interface{}) {
	for _, value := range values {
		valueCh <- value
	}
}

func innerToJson(value interface{}) []byte {
	vJson, err := json.MarshalIndent(value, "", " ")
	if err != nil {
		prints.ErrorPrint(err)
	}
	return vJson
}
