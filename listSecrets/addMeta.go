package listSecrets

import "strings"

func AddMeta(path string) string {
	if !strings.Contains(path, "metadata") {
		path = strings.Replace(path, "kv", "kv/metadata", 1)
	}

	return path
}
