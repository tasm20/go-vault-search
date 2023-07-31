package prints

import "fmt"

func MapsOfFoundSecrets(found map[string]map[string]string) {
	for path := range found {
		if found[path] == nil {
			continue
		}
		fmt.Printf("Found in: %s\n", path)
		for key, value := range found[path] {
			fmt.Printf("%s%s = %s\n", tabSpace, key, value)
		}
		fmt.Println()
	}
}
