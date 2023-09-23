package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

// 'postgres://user:password@localhost:5432/database'
// -d='postgres://practicum:practicum@localhost:5432/practicum'
var createTable string = `CREATE TABLE if not exists users (
							id serial primary key,
							login varchar(1000) UNIQUE,
							password varchar(1000),
							balance numeric(20,10)
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
func (udb *UserDB) Close() {
	if udb.DB != nil {
		udb.DB.Close()
	}
}
func (udb *UserDB) AddUser(ctx context.Context, u *domain.User) (int64, error) {
	logger.Log.Info("in database: ", zap.String("UserDB", "AddUser"))
	user := userDomainToStore(u)
	var id int64
	err := udb.DB.QueryRowContext(ctx, `insert into users (login,password,balance)
									 values($1,$2,$3) returning id;`, user.Login, user.Password, user.Balance).Scan(&id)
	if err != nil {
		logger.Log.Error("unable to add user", zap.Error(err))
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == `23505` {
				return 0, domain.ErrUniqueViolation
			} else {
				return 0, err
			}
		}
	}

	return id, nil
}
func (udb *UserDB) GetUser(ctx context.Context, login string) (*domain.User, error) {
	logger.Log.Info("in database: ", zap.String("UserDB", "GetUser"))
	row := udb.DB.QueryRowContext(ctx, `select id,password,balance from users  where login = $1;`, login)
	var (
		id       int
		balance  float64
		password string
	)
	err := row.Scan(&id, &password, &balance)
	if err == sql.ErrNoRows {
		logger.Log.Error("no user in db with that login", zap.Error(err))
		return nil, domain.ErrNotFound
	} else if err != nil {
		logger.Log.Error("unable to  get user", zap.Error(err))
		return nil, err
	}
	user, err := domain.NewUser(id, login, password, balance)
	if err != nil {
		logger.Log.Error("unable to create domain.User ", zap.Error(err))
		return nil, fmt.Errorf("%s %w", "unabe to create domain.User", err)
	}
	return user, nil
}
