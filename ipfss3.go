package ipfss3

import (
	"strings"
)

func ParsePath(orig string) (protocol, path string) {
	s := strings.Split(orig, "://")
	if len(s) == 1 {
		return "local", orig
	}
	return s[0], s[1]
}
