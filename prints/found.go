package prints

import (
	"fmt"
)

func PrintFound(foundCh chan string) bool {
	found := <-foundCh
	if found == "" {
		return false
	}

	fmt.Println(found)
	return true
}
