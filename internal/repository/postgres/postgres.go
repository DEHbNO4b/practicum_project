package postgres

import (
	"database/sql"

	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"go.uber.org/zap"
)

var createTable string = `CREATE TABLE IF NOT EXISTS metrics(
	id varchar(150) UNIQUE,
	type varchar(150),
	delta integer,
	value double precision
	);`

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(dsn string) (*PostgresDB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Log.Error("cannot open db", zap.Error(err))
		return nil, err
	}
	db.Exec(createTable)
	return &PostgresDB{DB: db}, nil
}
func (pdb *PostgresDB) Close() {
	if pdb.DB != nil {
		pdb.DB.Close()
	}
}
