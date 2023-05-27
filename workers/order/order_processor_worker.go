package worker

import (
	"github.com/markaz/dto"
	engine "github.com/markaz/kafka/readers/order"
	"github.com/markaz/services"
)

type OrderProcessor struct {
	logReader engine.LogReader
}

func NewOrderProcessor(logReader engine.LogReader) *OrderProcessor {
	t := &OrderProcessor{
		logReader: logReader,
	}
	t.logReader.RegisterObserver(t)
	return t
}

func (t *OrderProcessor) Start() {
	go t.logReader.Run()
}

func (t *OrderProcessor) OnOrderLog(log *dto.OrderRequest) error {
	return services.ProcessOrder(log)
}
