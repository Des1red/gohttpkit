package cors

import (
	"net/http"

	"github.com/Des1red/gohttpkit/internal/config"
)

func New(runtime config.Runtime) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if IsIgnored(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			origin := r.Header.Get("Origin")

			if runtime.Dev && origin == runtime.DevOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			if origin == runtime.ProdOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", Methods())
			w.Header().Set("Vary", "Origin")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
