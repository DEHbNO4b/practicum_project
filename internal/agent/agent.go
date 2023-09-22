package agent

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/config"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type AccrualAgent struct {
	ctx    context.Context
	client http.Client
	url    string
}

func NewAccrualAgent(ctx context.Context) *AccrualAgent {
	cfg := config.Get()
	cl := http.Client{Timeout: 1 * time.Second}
	a := &AccrualAgent{ctx: ctx, client: cl, url: cfg.AccrualSystemAdress}
	return a
}
func (a *AccrualAgent) GetAccrual(inputCh chan string) chan AccrualResponse {
	respCh := make(chan AccrualResponse)

	go func() {
		for number := range inputCh {
			defer close(respCh)
			req, err := http.NewRequest(http.MethodGet, a.url+`/api/orders/`+number, nil)
			if err != nil {
				logger.Log.Error("acrual server request err", zap.Error(err))
				respCh <- AccrualResponse{order: nil, err: err}
			}
			resp, err := a.client.Do(req)
			if err != nil {
				respCh <- AccrualResponse{order: nil, err: err}
				return
			}
			order := Order{}
			switch resp.StatusCode {
			case 200:
				render.DecodeJSON(resp.Body, &order)
				resp.Body.Close()
				dOrder, err := orderAgentToDomain(order)
				respCh <- AccrualResponse{order: dOrder, err: err}
			case 204:
				respCh <- AccrualResponse{order: nil, err: domain.ErrNotRegistered}
			case 429:
				respCh <- AccrualResponse{order: nil, err: domain.ErrTooManyRequests}
			default:
				respCh <- AccrualResponse{order: nil, err: domain.ErrUnexpectedRespStatus}
			}
		}
	}()
	return respCh
}

// fanOut принимает канал данных, порождает 10 горутин
func (a *AccrualAgent) FanOut(inputCh chan string) []chan AccrualResponse {
	// количество горутин add
	numWorkers := 10
	// каналы, в которые отправляются результаты
	channels := make([]chan AccrualResponse, numWorkers)

	for i := 0; i < numWorkers; i++ {
		// получаем канал из горутины add
		addResultCh := a.GetAccrual(inputCh)
		// отправляем его в слайс каналов
		channels[i] = addResultCh
	}
	// возвращаем слайс каналов
	return channels
}

// fanIn объединяет несколько каналов resultChs в один.
func fanIn(ctx context.Context, resultChs ...chan AccrualResponse) chan AccrualResponse {
	// конечный выходной канал в который отправляем данные из всех каналов из слайса, назовём его результирующим
	finalCh := make(chan AccrualResponse)

	// понадобится для ожидания всех горутин
	var wg sync.WaitGroup

	// перебираем все входящие каналы
	for _, ch := range resultChs {
		// в горутину передавать переменную цикла нельзя, поэтому делаем так
		chClosure := ch

		// инкрементируем счётчик горутин, которые нужно подождать
		wg.Add(1)

		go func() {
			// откладываем сообщение о том, что горутина завершилась
			defer wg.Done()

			// получаем данные из канала
			for data := range chClosure {
				select {
				// выходим из горутины, если канал закрылся
				case <-ctx.Done():
					return
				// если не закрылся, отправляем данные в конечный выходной канал
				case finalCh <- data:
				}
			}
		}()
	}
	go func() {
		// ждём завершения всех горутин
		wg.Wait()
		// когда все горутины завершились, закрываем результирующий канал
		close(finalCh)
	}()

	// возвращаем результирующий канал
	return finalCh
}
