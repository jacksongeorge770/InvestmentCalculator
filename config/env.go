package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	DBPort     string
	// JWTexpire  int64
	// JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "admin"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBAddress:  getEnv("DB_HOST", ""),
		DBName:     getEnv("DB_NAME", "investcalc"),
		// JWTexpire:  getEnvINT("JWTEXP", 3600*24*7),
		// JWTSecret:  getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// func getEnvINT(key string, fallback int64) int64 {
// 	if value, ok := os.LookupEnv(key); ok {
// 		i, err := strconv.ParseInt(value, 10, 64)
// 		if err != nil {
// 			return fallback
// 		}
// 		return i
// 	}
// 	return fallback
// }
