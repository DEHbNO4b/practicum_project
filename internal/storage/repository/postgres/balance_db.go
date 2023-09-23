package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"go.uber.org/zap"
)

var createBalanceTable string = `CREATE TABLE if not exists balance (
	current numeric(20,10),
	withdrawn numeric(20,10),
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
	_, err = db.Exec(createBalanceTable)
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
func (bdb *BalanceDB) GetByID(ctx context.Context, id int) (*domain.Balance, error) {
	logger.Log.Info("in database: ", zap.String("BalanceDB", "GetByID"))
	row := bdb.DB.QueryRowContext(ctx, `select current,withdrawn from balance where user_id =$1`, id)
	var c, w float64
	if err := row.Scan(&c, &w); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		logger.Log.Error("unable to get balance ", zap.Error(err))
		return nil, err
	}
	balance, _ := domain.NewBalance(c, w, id)
	return balance, nil
}
func (bdb *BalanceDB) UpdateBalance(ctx context.Context, balance *domain.Balance) error {
	logger.Log.Info("in database: ", zap.String("BalanceDB", "UpdateBalance"))
	_, err := bdb.DB.ExecContext(ctx, `UPDATE balance SET current=$1,withdrawn=$2 WHERE user_id=$3 `,
		balance.Current(), balance.Withdrown(), balance.UserID())
	if err != nil {
		logger.Log.Error("unable to update balance", zap.Error(err))
		return err
	}
	return nil
}

//	func (bdb *BalanceDB) AddAccrual(ctx context.Context, id, accrual int) error {
//		return nil
//	}
//
//	func (bdb *BalanceDB) WriteOff(ctx context.Context, id, sum int) error {
//		return nil
//	}
func (bdb *BalanceDB) NewBalance(ctx context.Context, b *domain.Balance) error {
	logger.Log.Info("in database: ", zap.String("BalanceDB", "NewBalance"))
	_, err := bdb.DB.ExecContext(ctx, `INSERT INTO balance (current,withdrawn,user_id) VALUES (0,0,$1); `, b.UserID())
	if err != nil {
		logger.Log.Error("unable to create new balance", zap.Error(err))
		return err
	}
	return nil
}
