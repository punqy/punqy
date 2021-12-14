package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	logger "github.com/sirupsen/logrus"
	"log"
	"os"
)

func Load() Config {
	var cfg Config
	cfg.AppEnv = AppEnv(os.Getenv("ENVIRONMENT"))
	if cfg.AppEnv == "" {
		cfg.AppEnv = EnvDev
	}
	envPath := fmt.Sprintf(".env.%s.local", cfg.AppEnv)
	logger.Warnf("Loading conf %s", envPath)
	if err := godotenv.Load(envPath); err != nil {
		logger.Warnf("File not found %s", envPath)
		envPath = fmt.Sprintf(".env.local")
		logger.Warnf("Loading conf %s", envPath)
		if err := godotenv.Load(envPath); err != nil {
		}
		logger.Warnf("File not found %s", envPath)
	}
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
