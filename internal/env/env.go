package env

import (
	"log"
	"os"
	"strings"
)

func Load(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Println("no .env file found")
		return
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if os.Getenv(key) == "" {
			_ = os.Setenv(key, value)
		}
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func Validate() {

	dev := os.Getenv("DEV")
	devOrigin := os.Getenv("DEV_ORIGIN")
	prodOrigin := os.Getenv("PROD_ORIGIN")

	if dev != "true" && dev != "false" {
		log.Fatal("ENV DEV must be 'true' or 'false'")
	}

	if dev == "true" && devOrigin == "" {
		log.Fatal("ENV DEV_ORIGIN is not set")
	}

	if prodOrigin == "" {
		log.Fatal("ENV PROD_ORIGIN is not set")
	}
}
