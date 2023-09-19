package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"go.uber.org/zap"
)

var createDebitTable string = `CREATE TABLE if not exists debits (
	order integer unique,
	sum integer ,
	time timestamptz,
	user_id integer
	);`

type DebitDB struct {
	DB *sql.DB
}

func NewDebitDB(dsn string) (*DebitDB, error) {
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
	_, err = db.Exec(createDebitTable)
	if err != nil {
		logger.Log.Error("cannot open db", zap.Error(err))
		return nil, err
	}
	return &DebitDB{DB: db}, nil
}

func (ddb *DebitDB) AddDebit(ctx context.Context, d *domain.Debit) error {
	_, err := ddb.DB.ExecContext(ctx, `insert into debits (order,sum,time,user_id)
						values($1,$2,$3,$4)`, d.Order(), d.Sum(), time.Now(), d.UserId())
	if err != nil {
		logger.Log.Error("unable to add debit to db", zap.Error(err))
		return fmt.Errorf("%s %w", "unable to add debit to db", err)
	}
	return nil
}
func (ddb *DebitDB) GetDebitsById(ctx context.Context, id int) ([]*domain.Debit, error) {
	rows, err := ddb.DB.QueryContext(ctx, `select order,sum,time from debits where id = $1;`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		logger.Log.Error("unable to get debits from db", zap.Error(err))
		return nil, fmt.Errorf("%s %w", "unable to get debits from db", err)
	}
	defer rows.Close()
	var (
		o, s int
		t    time.Time
	)
	debits := make([]*domain.Debit, 0, 10)
	for rows.Next() {
		if err := rows.Scan(&o, &s, &t); err != nil {
			logger.Log.Error("unable to scan debit params ", zap.Error(err))
			return nil, fmt.Errorf("%s %w", "unable to scan debit params", err)
		}
		d, _ := domain.NewDebit(o, s, t, id)
		debits = append(debits, d)
	}
	return debits, nil
}
