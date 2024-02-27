package prints

import "fmt"

func PrintAllinPath(path string, secrets map[string]interface{}) {
	fmt.Printf("Found in %s:\n", path)
	for key, val := range secrets {
		fmt.Printf("\t%s:\t%s\n", key, val)
	}
}
