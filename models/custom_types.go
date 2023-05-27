package models

import "fmt"

type OrderStatus string

func (d OrderStatus) String() string {
	return string(d)
}

// NewOrderStatusString return error if passed string is not one of COMPLETED, PENDING, CANCELLED
func NewOrderStatusString(s string) (OrderStatus, error) {
	status := OrderStatus(s)
	switch status {
	case OrderStatusPending:
	case OrderStatusCancelled:
	case OrderStatusCompleted:
	default:
		return status, fmt.Errorf("invalid order status: %v", s)
	}
	return status, nil
}
