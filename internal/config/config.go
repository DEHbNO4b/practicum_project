package config

import (
	"flag"
	"os"
	"sync"
)

type Config struct {
	RunAdress           string
	DatabaseURL         string
	AccrualSystemAdress string
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		flag.StringVar(&config.RunAdress, "a", "", "адрес и порт запуска сервера")
		flag.StringVar(&config.DatabaseURL, "d", "", "адрес подключения к базе данных")
		flag.StringVar(&config.AccrualSystemAdress, "r", "", "адрес системы рассчета начислений")
		flag.Parse()
		if addr := os.Getenv("RUN_ADDRESS"); addr != "" {
			config.RunAdress = addr
		}
		if addr := os.Getenv("DATABASE_URI"); addr != "" {
			config.DatabaseURL = addr
		}
		if addr := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); addr != "" {
			config.AccrualSystemAdress = addr
		}
	})
	return &config
}
