package writers

import "github.com/markaz/models"

const (
	TopicOrder        = "orders"
	TopicOrderMessage = "orders_message"
)

const (
	//order  logs
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
