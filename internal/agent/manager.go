package agent

import (
	"context"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/storage"
)

type Manager struct {
	agent   *AccrualAgent
	period  time.Duration
	storage *storage.Storage
}

func NewManager(ctx context.Context, store *storage.Storage) *Manager {
	agent := NewAccrualAgent(ctx)
	period := 10 * time.Second

	return &Manager{
		agent:   agent,
		period:  period,
		storage: store,
	}
}
func (m *Manager) Start(ctx context.Context) {
	ticker := time.NewTicker(m.period)
	defer ticker.Stop()

	for {
		<-ticker.C
		orders := GetNewOrdersFromDb(ctx)
		for _, o := range orders {
			order, err := m.agent.GetAccrual(o.Number())
		}

	}
}

func GetNewOrdersFromDb(ctx context.Context) []*domain.Order {
	return nil
}
