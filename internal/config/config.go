// sentiric-identity-service/internal/config/config.go
package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	GRPCPort string
	HttpPort string
	CertPath string
	KeyPath  string
	CaPath   string
	LogLevel string
	Env      string

	// Identity servisi için kritik ayarlar
	JWTSecret          string
	UserAuthServiceURL string // User DB/Auth için
}

func Load() (*Config, error) {
	godotenv.Load()

	// Harmonik Mimari Portlar (Control Plane, 114XX bloğu atandı)
	return &Config{
		GRPCPort: GetEnv("IDENTITY_SERVICE_GRPC_PORT", "11411"),
		HttpPort: GetEnv("IDENTITY_SERVICE_HTTP_PORT", "11410"),

		CertPath: GetEnvOrFail("IDENTITY_SERVICE_CERT_PATH"),
		KeyPath:  GetEnvOrFail("IDENTITY_SERVICE_KEY_PATH"),
		CaPath:   GetEnvOrFail("GRPC_TLS_CA_PATH"),
		LogLevel: GetEnv("LOG_LEVEL", "info"),
		Env:      GetEnv("ENV", "production"),

		JWTSecret:          GetEnvOrFail("JWT_SECRET"),
		UserAuthServiceURL: GetEnv("USER_AUTH_SERVICE_URL", "user-service:12011"),
	}, nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetEnvOrFail(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal().Str("variable", key).Msg("Gerekli ortam değişkeni tanımlı değil")
	}
	return value
}
