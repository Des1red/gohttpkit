package cors

import "strings"

var ignored = map[string]struct{}{}

func SetIgnored(paths ...string) {
	for _, p := range paths {
		if p == "" {
			continue
		}
		ignored[p] = struct{}{}
	}
}

func IsIgnored(path string) bool {
	for p := range ignored {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}
