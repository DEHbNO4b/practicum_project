package config

import (
	"flag"
	"os"
	"sync"
)

type Config struct {
	RunAdress           string
	DatabaseUrl         string
	AccrualSystemAdress string
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		flag.StringVar(&config.RunAdress, "a", "", "адрес и порт запуска сервера")
		flag.StringVar(&config.DatabaseUrl, "d", "", "адрес подключения к базе данных")
		flag.StringVar(&config.AccrualSystemAdress, "r", "", "адрес системы рассчета начислений")
		flag.Parse()
		if addr := os.Getenv("RUN_ADDRESS"); addr != "" {
			config.RunAdress = addr
		}
		if addr := os.Getenv("DATABASE_URL"); addr != "" {
			config.DatabaseUrl = addr
		}
		if addr := os.Getenv("ACCRUAL_SYSTEM_ADRESS"); addr != "" {
			config.AccrualSystemAdress = addr
		}
	})
	return &config
}
