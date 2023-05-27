package order_reader

import (
	"github.com/markaz/models"
)

type LogReader interface {
	RegisterObserver(observer LogObserver)

	Run()
}

type LogObserver interface {
	OnStockUpdateLog(log []models.Product) error
}
