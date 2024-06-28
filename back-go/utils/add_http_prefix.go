package utils

import (
	"strings"
)

func AddHTTPPrefixIfNeeded(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "http://" + url
}
