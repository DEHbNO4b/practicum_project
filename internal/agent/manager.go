package agent

import (
	"context"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/DEHbNO4b/practicum_project/internal/storage"
	"go.uber.org/zap"
)

type Manager struct {
	ctx     context.Context
	agent   *AccrualAgent
	period  time.Duration
	storage *storage.Storage
}

func NewManager(ctx context.Context, store *storage.Storage) *Manager {
	agent := NewAccrualAgent(ctx)
	period := 10 * time.Second

	return &Manager{
		ctx:     ctx,
		agent:   agent,
		period:  period,
		storage: store,
	}
}
func (m *Manager) Start() {
	ticker := time.NewTicker(m.period)
	defer ticker.Stop()
	for {
		<-ticker.C
		go m.accrualInteraction(m.ctx)
	}
}

func (m *Manager) accrualInteraction(ctx context.Context) {
	numbersCh := make(chan chan string)
	defer close(numbersCh)
	orders := m.getNewOrdersFromDB(ctx)
	inputCh := generator(ctx, orders)
	chanels := m.agent.FanOut(inputCh)
	addResultCh := fanIn(ctx, chanels...)
	m.updateDB(ctx, addResultCh)
}
func (m *Manager) getNewOrdersFromDB(ctx context.Context) []*domain.Order {
	orders, err := m.storage.Order.GetNewOrders(ctx)
	if err != nil {
		logger.Log.Error("unable to GetNewOrders from storage", zap.Error(err))
	}
	return orders
}
func (m *Manager) updateDB(ctx context.Context, inputCh chan AccrualResponse) {
	for resp := range inputCh {
		if resp.err != nil {
			continue
		}
		m.storage.Order.UpdateOrder(ctx, resp.order)
		ord, err := m.storage.Order.GetOrderByNumber(ctx, resp.order.Number())
		if err != nil {
			logger.Log.Error("unable to get order  by numberfrom DB", zap.Error(err))
			continue
		}
		balance, err := m.storage.Balance.GetByID(ctx, ord.UserID())
		if err != nil {
			logger.Log.Error("unable to get balance from DB", zap.Error(err))
			continue
		}
		balance.SetCurrent(balance.Current() + float64(resp.order.Accrual()))
		m.storage.Balance.UpdateBalance(ctx, balance)
	}
}

// generator возвращает канал с данными
func generator(ctx context.Context, input []*domain.Order) chan string {
	// канал, в который будем отправлять данные из слайса
	inputCh := make(chan string)

	// горутина, в которой отправляем в канал  inputCh данные
	go func() {
		// как отправители закрываем канал, когда всё отправим
		defer close(inputCh)

		// перебираем все данные в слайсе
		for _, data := range input {
			select {
			// если doneCh закрыт, сразу выходим из горутины
			case <-ctx.Done():
				return
			// если doneCh не закрыт, кидаем в канал inputCh данные data
			case inputCh <- data.Number():
			}
		}
	}()

	// возвращаем канал для данных
	return inputCh
}
