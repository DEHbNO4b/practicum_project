package main

import (
	"flag"
	"os"
)

// - адрес и порт запуска сервиса: переменная окружения ОС `RUN_ADDRESS` или флаг `-a`
// - адрес подключения к базе данных: переменная окружения ОС `DATABASE_URI` или флаг `-d`
// - адрес системы расчёта начислений: переменная окружения ОС `ACCRUAL_SYSTEM_ADDRESS` или флаг `-r`

type Config struct {
	Run_adress            string
	Database_url          string
	Accrual_system_adress string
}

func parseFlag() Config {
	cfg := Config{}
	flag.StringVar(&cfg.Run_adress, "a", "", "адрес и порт запуска сервера")
	flag.StringVar(&cfg.Database_url, "d", "", "адрес подключения к базе данных")
	flag.StringVar(&cfg.Accrual_system_adress, "r", "", "адрес системы рассчета начислений")
	flag.Parse()
	if addr := os.Getenv("RUN_ADDRESS"); addr != "" {
		cfg.Run_adress = addr
	}
	if addr := os.Getenv("DATABASE_URL"); addr != "" {
		cfg.Database_url = addr
	}
	if addr := os.Getenv("ACCRUAL_SYSTEM_ADRESS"); addr != "" {
		cfg.Accrual_system_adress = addr
	}
	return cfg
}
