package prints

import (
	"fmt"
	"strings"
)

func MapsOfFoundSecrets(found map[string]map[string]string) {
	for path := range found {
		if found[path] == nil {
			continue
		}
		fmt.Printf("Found in: %s\n", path)
		for key, value := range found[path] {
			if strings.Contains(value, "\n") {
				doubleTabSpace := strings.Repeat(tabSpace, 2)
				value = strings.Replace(value, "\n", "\n"+doubleTabSpace, -1)
			}
			fmt.Printf("%s%s: %s\n", tabSpace, key, value)
		}
		fmt.Println()
	}
}
