package worker

import (
	engine "github.com/markaz/kafka/readers/order_message"
	"github.com/markaz/models"
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

func (t *OrderProcessor) OnStockUpdateLog(products []models.Product) error {
	return services.UpdateProductsStock(products)
}
