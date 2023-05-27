package order_reader

import "github.com/markaz/models"

const (
	LogTypeStockUpdate = LogType("PRODUCT_STOCK_UPDATE")
)

type LogType string

type Base struct {
	Type      LogType
	TimeStamp int64
}

type ProductStockUpdateLog struct {
	Base
	Products []models.Product
}
