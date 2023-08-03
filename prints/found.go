package prints

import (
	"fmt"
)

func PrintFound(foundCh chan string) {
	found := <-foundCh
	if found != "" {
		fmt.Println(found)
	}
}
