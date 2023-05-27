package writers

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/markaz/conf"
	"github.com/markaz/logger"
	"github.com/markaz/models"
	kafka "github.com/segmentio/kafka-go"
)

var (
	orderMessageWriter sync.Map
)

func SubmitStockUpdateLog(products []models.Product) error {
	log := &ProductStockUpdateLog{
		Base: Base{
			Type:      LogTypeStockUpdate,
			TimeStamp: time.Now().Unix(),
		},
		Products: products,
	}
	buf, err := json.Marshal(log)
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	err = getOrderMessageWriter().WriteMessages(context.Background(), kafka.Message{Value: buf})
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	return nil
}

func getOrderMessageWriter() *kafka.Writer {
	writer, found := orderMessageWriter.Load(TopicOrderMessage)
	if found {
		return writer.(*kafka.Writer)
	}

	newWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      conf.Config.KafkaBrokers,
		Topic:        TopicOrderMessage,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 5 * time.Millisecond,
	})
	orderMessageWriter.Store(TopicOrderMessage, newWriter)
	return newWriter
}
