package postgres

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

var createOrderTable string = `CREATE TABLE if not exists orders (
	number varchar(1000),
	status varchar(1000) ,
	accrual numeric(20,10),
	uploaded_at timestamptz,
	user_id integer
	);`

type OrderDB struct {
	DB *sql.DB
}

func NewOrderDB(dsn string) (*OrderDB, error) {
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
		logger.Log.Error("cannot open db", zap.Error(err))
		return nil, err
	}
	return &OrderDB{DB: db}, nil
}
func (odb *OrderDB) Close() {
	if odb.DB != nil {
		odb.DB.Close()
	}
}
func (odb *OrderDB) AddOrder(ctx context.Context, order *domain.Order) error {
	_, err := odb.DB.ExecContext(ctx, `INSERT INTO orders (number,status,accrual,uploaded_at,user_id)
								VALUES ($1,$2,$3,$4,$5);`, order.Number(), order.Status(), order.Accrual(), time.Now(), order.UserID())
	if err != nil {
		logger.Log.Error("unable to insert order to db", zap.Error(err))
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
func (odb *OrderDB) UpdateOrder(ctx context.Context, order *domain.Order) error {
	_, err := odb.DB.ExecContext(ctx, `UPDATE orders SET status=$1,accrual=$2 WHERE number=$3 `,
		order.Status(), order.Accrual(), order.Number())
	if err != nil {
		logger.Log.Error("unable to update order", zap.Error(err))
		return err
	}
	return nil
}
func (odb *OrderDB) GetOrdersByID(ctx context.Context, id int) ([]*domain.Order, error) {
	logger.Log.Info("in get orders by id in postgres")
	rows, err := odb.DB.QueryContext(ctx, `SELECT number,status,accrual,uploaded_at from orders where user_id=$1;`, id)
	if err != nil {
		logger.Log.Error("unable to load order params from db", zap.Error(err))
		return nil, err
	}
	var (
		a    float64
		n, s string
		u    time.Time
	)
	orders := make([]*domain.Order, 0, 20)
	for rows.Next() {
		err := rows.Scan(&n, &s, &a, &u)
		if err != nil {
			logger.Log.Error("unable to scan order parameters from db", zap.Error(err))
			return nil, err
		}
		o, err := domain.NewOrder(n, s, a, u, id)
		if err != nil {
			logger.Log.Error("unable to create new order", zap.Error(err))
			return nil, err
		}
		orders = append(orders, o)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		logger.Log.Info("nil order data for user:", zap.String("id", strconv.Itoa(id)))
		return nil, domain.ErrNilData
	}
	if len(orders) == 0 {
		return nil, domain.ErrNotFound
	}
	return orders, nil
}
func (odb *OrderDB) GetOrderByNumber(ctx context.Context, number string) (*domain.Order, error) {
	// row := odb.DB.QueryRowContext(ctx, `SELECT status,accrual,uploaded_at,user_id from orders where number=$1;`, number)
	row := odb.DB.QueryRowContext(ctx, `SELECT status,accrual,uploaded_at,user_id from orders where number=$1;`, number)
	var (
		id int
		a  float64
		s  string
		u  time.Time
	)
	err := row.Scan(&s, &a, &u, &id)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	} else if err != nil {
		logger.Log.Error("unable to scan order parameters from db", zap.Error(err))
		return nil, err
	}
	o, _ := domain.NewOrder(number, s, a, u, id)
	return o, nil
}
func (odb *OrderDB) GetNewOrders(ctx context.Context) ([]*domain.Order, error) {
	logger.Log.Info("in get orders by id in postgres")
	rows, err := odb.DB.QueryContext(ctx, `SELECT number,accrual,uploaded_at,user_id from orders where status='NEW';`)
	if err != nil {
		logger.Log.Error("unable to load order params from db", zap.Error(err))
		return nil, err
	}
	var (
		a  float64
		id int
		n  string
		u  time.Time
	)
	orders := make([]*domain.Order, 0, 30)
	for rows.Next() {
		err := rows.Scan(&n, &a, &u, &id)
		if err != nil {
			logger.Log.Error("unable to scan order parameters from db", zap.Error(err))
			return nil, err
		}
		o, err := domain.NewOrder(n, "NEW", a, u, id)
		if err != nil {
			logger.Log.Error("unable to create new order", zap.Error(err))
			return nil, err
		}
		orders = append(orders, o)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		logger.Log.Info("nil order data for user:", zap.String("id", strconv.Itoa(id)))
		return nil, domain.ErrNilData
	}
	if len(orders) == 0 {
		return nil, domain.ErrNotFound
	}
	return orders, nil
}
