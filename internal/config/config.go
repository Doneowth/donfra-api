package config

import "os"

type Config struct {
	Addr       string
	Passcode   string
	BaseURL    string
	CORSOrigin string
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func Load() Config {
	return Config{
		Addr:       getenv("ADDR", ":8080"),
		Passcode:   getenv("PASSCODE", "19950413"),
		BaseURL:    getenv("BASE_URL", ""),
		CORSOrigin: getenv("CORS_ORIGIN", "http://localhost:3000"),
	}
}
