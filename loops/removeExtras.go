package loops

import "strings"

func RemoveExtrasInPath(str string) string {
	res := strings.Replace(str, "kv/metadata", "kv/", 1)
	res = innerCutDoubleSlash(res)

	return res
}

func innerCutDoubleSlash(str string) string {
	res := strings.Replace(str, "//", "/", -1)
	if strings.Contains(res, "//") {
		return innerCutDoubleSlash(res)
	}
	return res
}
