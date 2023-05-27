package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	"github.com/markaz/conf"
	"github.com/markaz/logger"
)

var DbConnection *gorm.DB

func OpenDBConnection() (*gorm.DB, error) {
	logger.Logger.Info("Grabbing variables from config file.")
	dbHost := conf.Config.Mysql.Host
	dbPort := conf.Config.Mysql.Port
	dbName := conf.Config.Mysql.DbName
	dbUserName := conf.Config.Mysql.DbUserName
	dbPassword := conf.Config.Mysql.DbPassword

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUserName,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	logger.Logger.Info("Connecting MYSQL Database.")

	connection, err := gorm.Open("mysql", dbUri)
	if err != nil {
		logger.Logger.Error("ERROR: Couldn't establish database connection: %s", zap.Error(err))
		return nil, err
	}

	logger.Logger.Info("MYSQL Database Connected Successfully.")

	return connection, nil
}
