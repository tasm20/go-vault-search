package prints

import "fmt"

const notFoundText = "NOT FOUND"

var notFound = fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, notFoundText)

func NotFound() {
	fmt.Println(notFound)
}
