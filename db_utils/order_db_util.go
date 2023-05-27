package db_utils

import (
	"go.uber.org/zap"

	"github.com/markaz/db"
	"github.com/markaz/logger"
	"github.com/markaz/models"
)

func CheckActiveUserExists(userId int64) error {
	var user models.User
	if err := db.DbConnection.Where("user_id = ? and status = ?", userId, "ACTIVE").First(&user).Error; err != nil {
		logger.Logger.Error("error in getting active user by user id: ", zap.Error(err))
		return err
	}
	return nil
}

func AddOrder(order models.Order) error {
	if err := db.DbConnection.Create(&order).Error; err != nil {
		logger.Logger.Error("error in AddOrder", zap.Error(err))
		return err
	}
	return nil
}

func AddOrderItems(orderItems []models.OrderItem) error {
	for _, item := range orderItems {
		if err := db.DbConnection.Create(&item).Error; err != nil {
			logger.Logger.Error("error in AddOrderItems", zap.Error(err))
			return err
		}
	}
	return nil
}

func GetOrdersByUserId(userId int64) ([]models.Order, error) {
	var orders []models.Order
	if err := db.DbConnection.Where("user_id = ? ", userId).Find(&orders).Error; err != nil {
		logger.Logger.Error("error in getting orders  by user id: ", zap.Error(err))
		return orders, err
	}
	return orders, nil
}

func UpdateProductsStock(products []models.Product) error {
	// TODO : Update orders stock here
	return nil
}
