package prints

import (
	"fmt"
)

func PrintList(list []string, vaultPath string) {
	if len(list) == 0 {
		return
	}

	fmt.Printf("\n"+listFirstString, vaultPath)
	for _, item := range list {
		fmt.Printf("%s%s\n", tabSpace, item)
	}
}
