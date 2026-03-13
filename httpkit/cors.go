package httpkit

import (
	"net/http"

	internalcors "github.com/Des1red/gohttpkit/internal/cors"
)

func CORSIgnorePaths(paths ...string) {
	internalcors.SetIgnored(paths...)
}

func CORS(runtime Runtime) Middleware {
	return internalcors.New(runtime)
}

func With(runtime Runtime, h http.Handler) http.Handler {
	return Chain(
		h,
		CORS(runtime),
	)
}

func CORSMethods(methods string) {
	internalcors.SetMethods(methods)
}
