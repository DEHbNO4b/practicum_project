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

// fanOut принимает канал данных, порождает 10 горутин
func (m *Manager) FanOut(ctx context.Context, inputCh chan string) []chan AccrualResponse {
	// количество горутин add
	numWorkers := 10
	// каналы, в которые отправляются результаты
	channels := make([]chan AccrualResponse, numWorkers)

	for i := 0; i < numWorkers; i++ {
		// получаем канал из горутины add
		addResultCh := m.agent.GetAccrual()
		// отправляем его в слайс каналов
		channels[i] = addResultCh
	}

	// возвращаем слайс каналов
	return channels
}
