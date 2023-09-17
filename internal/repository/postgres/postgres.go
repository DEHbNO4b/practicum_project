package postgres

import (
	"context"
	"database/sql"
	"errors"

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
	password varchar(1000)
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
func (pdb *UserDB) AddUser(ctx context.Context, u *domain.User) error {
	user := userDomainToStore(u)
	_, err := pdb.DB.Exec(`insert into users (login,password) values($1,$2);`, user.Login, user.Password)
	if err != nil {
		logger.Log.Error("unable to add user", zap.Error(err))
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == `23505` {
				return domain.ErrUniqueViolation
			} else {
				return err
			}
		}
	}
	return nil
}
func (pdb *UserDB) GetUserPassword(ctx context.Context, login string) (string, error) {
	row := pdb.DB.QueryRowContext(ctx, `select password from users  where login = $1;`, login)
	var pas string
	err := row.Scan(&pas)
	if err == sql.ErrNoRows {
		logger.Log.Error("no user in db with that login", zap.Error(err))
		return "", domain.ErrNotFound
	} else if err != nil {
		logger.Log.Error("unable to  get user", zap.Error(err))
		return "", err
	}
	return pas, nil
}
