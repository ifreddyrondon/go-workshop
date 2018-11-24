package args

import (
	"os"
	"strconv"
)

// GetFirstInt get the first int from args
func GetFirstInt(args []string) (int, bool) {
	if len(args) < 2 {
		return 0, false
	}
	first, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 0, false
	}
	return first, true
}
