package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"go.uber.org/zap"
)

var createBalanceTable string = `CREATE TABLE if not exists balance (
	current integer,
	withdrawn integer,
	user_id integer UNIQUE
	);`

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(dsn string) (*BalanceDB, error) {
	if dsn == "" {
		err := errors.New("dsn string can not be empty")
		logger.Log.Error("empty dsn string", zap.Error(err))
		return nil, err
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Log.Error("cannot open db", zap.Error(err))
		return nil, err
	}
	_, err = db.Exec(createOrderTable)
	if err != nil {
		logger.Log.Error("cannot create db", zap.Error(err))
		return nil, err
	}
	return &BalanceDB{DB: db}, nil
}
func (bdb *BalanceDB) Close() {
	if bdb.DB != nil {
		bdb.DB.Close()
	}
}
func (bdb *BalanceDB) AddAccrual(ctx context.Context, id, accrual int) error {
	return nil
}
func (bdb *BalanceDB) WriteOff(ctx context.Context, id, sum int) error {
	return nil
}
