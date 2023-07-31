package loops

import (
	"encoding/json"
	"github.com/hashicorp/vault/api"
	"github.com/tasm20/go-vault-search/prints"
)

func SecretsLoop(secretsList map[string]*api.KVSecret, secretsDataCh chan map[string]map[string][]byte) {
	var secretDATA = make(map[string]map[string][]byte)
	for key, values := range secretsList {
		valueMapCh := make(chan map[string][]byte)
		go innerSecretValueLoop(values, valueMapCh)
		secretDATA[key] = <-valueMapCh
	}

	secretsDataCh <- secretDATA
}

func innerSecretValueLoop(values *api.KVSecret, valueMapCh chan map[string][]byte) {
	var newMap = make(map[string][]byte)

	for key, value := range values.Data {
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
