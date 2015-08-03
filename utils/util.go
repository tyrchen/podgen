package utils

import (
	"path"
	"strings"
)

func Urljoin(prefix string, paths ...string) string {
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	return prefix + path.Join(paths...)
}
