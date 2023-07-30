package prints

import "fmt"

const (
	listFirstString = "in %s was found:\n"
	tabSpace        = "    "
	notFoundText    = "NOT FOUND"
)

var notFound = fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, notFoundText)
