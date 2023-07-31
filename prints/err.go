package prints

import (
	"fmt"
	"os"
)

func ErrorPrint(err error) {
	redColor := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, err)
	fmt.Println(redColor)
	os.Exit(2)
}
