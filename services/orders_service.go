package services

import (
	"time"

	"github.com/markaz/db_utils"
	"github.com/markaz/dto"
	"github.com/markaz/kafka/writers"
	"github.com/markaz/models"
	"github.com/rs/xid"
	"github.com/spf13/cast"
)

func CreateOrder(order dto.OrderRequest) error {
	//check user valid
	err := db_utils.CheckActiveUserExists(order.UserId)
	if err != nil {
		return err
	}
	//check products exist
	err = checkProductsExist(order.Products)
	if err != nil {
		return err
	}
	//process payment
	err = processPayment(order.PaymentMethod, order.PaymentMethodData)
	if err != nil {
		return err
	}
	//submit kafka log
	err = writers.SubmitOrder(&order)
	if err != nil {
		refundPayment(order.PaymentMethod, order.PaymentMethodData)
		return err
	}
	return nil
}

func ProcessOrder(order *dto.OrderRequest) error {
	orderModel, orderItems := createOrderModels(order)
	// add order to db
	err := db_utils.AddOrder(orderModel)
	if err != nil {
		return err
	}
	err = db_utils.AddOrderItems(orderItems)
	if err != nil {
		return err
	}
	// send stock update log
	err = writers.SubmitStockUpdateLog(order.Products)
	if err != nil {
		return err
	}
	return nil
}

func GetCustomerOrders(userId string) ([]models.Order, error) {
	//check user valid
	err := db_utils.CheckActiveUserExists(cast.ToInt64(userId))
	if err != nil {
		return nil, err
	}
	return db_utils.GetOrdersByUserId(cast.ToInt64(userId))
}

func UpdateProductsStock(products []models.Product) error {
	return db_utils.UpdateProductsStock(products)
}

// /////////////////////////////////////////////////////////////////
// ///////////////// Private Methods ///////////////////////////////
// /////////////////////////////////////////////////////////////////
func checkProductsExist(products []models.Product) error {
	//TODO : Check products stock in db
	return nil
}

func processPayment(paymentMethod string, paymentMethodDetails map[string]interface{}) error {
	//TODO : Process payment of the order with respect to payment method
	return nil
}

func refundPayment(paymentMethod string, paymentMethodDetails map[string]interface{}) error {
	//TODO : refund payment in case of writing log to kafka fails because order can't be processed without that now
	return nil
}

func createOrderModels(order *dto.OrderRequest) (models.Order, []models.OrderItem) {
	orderModel := models.Order{
		OrderNumber: xid.New().String(),
		UserId:      order.UserId,
		Status:      models.OrderStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	orderItems := make([]models.OrderItem, 0)
	for _, product := range order.Products {
		orderItems = append(orderItems, models.OrderItem{
			OrderNumber: orderModel.OrderNumber,
			UserId:      order.UserId,
			ProductId:   product.Id,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}
	return orderModel, orderItems
}
