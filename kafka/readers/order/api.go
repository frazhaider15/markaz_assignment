package order_reader

import "github.com/markaz/dto"

type LogReader interface {
	RegisterObserver(observer LogObserver)
	Run()
}

type LogObserver interface {
	OnOrderLog(log *dto.OrderRequest) error
}
