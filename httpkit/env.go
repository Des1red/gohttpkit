package httpkit

import (
	"os"

	internalenv "github.com/Des1red/gohttpkit/internal/env"
)

func BuildRuntime(envPath string) Runtime {

	internalenv.Load(envPath)
	internalenv.Validate()

	return Runtime{
		Dev:        os.Getenv("DEV") == "true",
		DevOrigin:  os.Getenv("DEV_ORIGIN"),
		ProdOrigin: os.Getenv("PROD_ORIGIN"),
		Port:       ":" + internalenv.GetPort(),
	}
}
