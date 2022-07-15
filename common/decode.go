package common

import "strings"

func DefaultPathParser(path string) string {
	lastIdx := len(path) - 1
	if path[lastIdx] == '/' {
		path = path[:lastIdx]
	}
	// accept only "/api/resource"
	t := strings.Split(path, "/")
	if len(t) == 3 {
		return t[2]
	}
	return ""
}
