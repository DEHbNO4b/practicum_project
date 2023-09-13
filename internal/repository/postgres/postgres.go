package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"go.uber.org/zap"
)

var createTable string = `CREATE TABLE practicumusers (
	id integer primary key,
	login varchar(100),
	password varchar(200)
	);`

type UserDB struct {
	DB *sql.DB
}

func NewUserDB(dsn string) (*UserDB, error) {
	if dsn == "" {
		err := errors.New("dsn string can not be empty")
		logger.Log.Error("cannot open db", zap.Error(err))
		return nil, err
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Log.Error("cannot open db", zap.Error(err))
		return nil, err
	}
	_, err = db.Exec(createTable)
	if err != nil {
		logger.Log.Error("cannot open db", zap.Error(err))
		return nil, err
	}
	return &UserDB{DB: db}, nil
}
func (pdb *UserDB) Close() {
	if pdb.DB != nil {
		pdb.DB.Close()
	}
}
func (pdb *UserDB) AddUser(ctx context.Context, user *domain.User) error {
	return nil
}
