package prints

import (
	"fmt"
	"strings"
)

func PrintList(list []string, vaultPath string) {
	if len(list) == 0 {
		return
	}

	vaultPath = strings.Replace(vaultPath, "kv/metadata", "kv", 1)
	fmt.Printf("\n"+listFirstString, vaultPath)
	for _, item := range list {
		fmt.Printf("%s%s\n", tabSpace, item)
	}
}
