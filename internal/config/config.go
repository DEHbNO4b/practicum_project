package config

import (
	"flag"
	"os"
	"sync"
)

type Config struct {
	Run_adress            string
	Database_url          string
	Accrual_system_adress string
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		flag.StringVar(&config.Run_adress, "a", "", "адрес и порт запуска сервера")
		flag.StringVar(&config.Database_url, "d", "", "адрес подключения к базе данных")
		flag.StringVar(&config.Accrual_system_adress, "r", "", "адрес системы рассчета начислений")
		flag.Parse()
		if addr := os.Getenv("RUN_ADDRESS"); addr != "" {
			config.Run_adress = addr
		}
		if addr := os.Getenv("DATABASE_URL"); addr != "" {
			config.Database_url = addr
		}
		if addr := os.Getenv("ACCRUAL_SYSTEM_ADRESS"); addr != "" {
			config.Accrual_system_adress = addr
		}
	})
	return &config
}
