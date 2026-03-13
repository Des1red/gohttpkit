package cors

import "strings"

var allowedMethods = "GET, OPTIONS"

func SetMethods(methods string) {
	methods = strings.TrimSpace(methods)

	if methods == "" {
		allowedMethods = "*"
		return
	}

	parts := strings.Split(methods, ",")
	for i := range parts {
		parts[i] = strings.ToUpper(strings.TrimSpace(parts[i]))
	}

	allowedMethods = strings.Join(parts, ", ")
}

func Methods() string {
	return allowedMethods
}
