package main

import (
	"os"

	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/DEHbNO4b/practicum_project/internal/maindb"
)

func main() {
	if err := logger.Initialize("info"); err != nil {
		panic(err)
	}
	if err := run(); err != nil {
		logger.Log.Fatal(err.Error())
		os.Exit(0)
	}
}
func run() error {
	cfg := parseFlag()
	pdb, err := maindb.NewPostgresDB(cfg.Database_url)
	if err != nil {

	}
	return nil
}
