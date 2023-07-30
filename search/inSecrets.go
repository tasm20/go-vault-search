package search

import (
	"fmt"
)

func InSecrets() {

}

func innerSecretValueLoop(values map[string]interface{}) {
	for key, value := range values {
		fmt.Println(key, value)
	}
}
