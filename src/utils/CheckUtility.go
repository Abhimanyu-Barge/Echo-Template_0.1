package utils

import (
	"strings"
)

func IsOpenUrl(URl string) bool {
	if strings.Contains(URl, "/o/") {
		return true
	}
	return false
}
