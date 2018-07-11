package numeronymize

import (
	"fmt"
	"strings"
)

func Numeronymize(orig string) string {
	target := strings.TrimRight(orig, "\n")

	len := len(target)
	if len <= 2 {
		return target
	}

	return fmt.Sprintf("%c%d%c", target[0], len-2, target[len-1])
}
