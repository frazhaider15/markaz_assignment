package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"github.com/markaz/conf"
	order_reader "github.com/markaz/kafka/readers/order"
	order_message_reader "github.com/markaz/kafka/readers/order_message"
	"github.com/markaz/logger"
	routes "github.com/markaz/routes"
	order_worker "github.com/markaz/workers/order"
	order_message_worker "github.com/markaz/workers/order_message"

	"github.com/markaz/db"
)

func main() {
	logger.InitLogger()
	conf.SetConfig()

	var err error
	db.DbConnection, err = db.OpenDBConnection()
	if err != nil {
		logger.Logger.Error("Error in creating database client: ", zap.Error(err))
		return
	}
	defer db.DbConnection.Close()

	// starting log kafka readers
	order_worker.NewOrderProcessor(order_reader.NewOrderLogReader(conf.Config.KafkaBrokers)).Start()
	order_message_worker.NewOrderProcessor(order_message_reader.NewOrderMessageLogReader(conf.Config.KafkaBrokers)).Start()

	// starting rest server
	r := routes.SetupRouter()
	if err = r.Run(conf.Config.RestServer.Addr); err != nil {
		logger.Logger.Error("Failed to run server", zap.Error(err))
	}
	logger.Logger.Info("Server Listening for API Calls.")
}
