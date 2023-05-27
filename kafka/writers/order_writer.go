package writers

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/markaz/conf"
	"github.com/markaz/dto"
	kafka "github.com/segmentio/kafka-go"
)

var (
	orderWriter sync.Map
)

// SubmitOrder submits or writes order on kafka
func SubmitOrder(order *dto.OrderRequest) error {
	buf, err := json.Marshal(order)
	if err != nil {
		return err
	}
	err = getOrderWriter().WriteMessages(context.Background(), kafka.Message{Value: buf})
	if err != nil {
		return err
	}
	return nil
}

func getOrderWriter() *kafka.Writer {
	writer, found := orderWriter.Load(TopicOrder)
	if found {
		return writer.(*kafka.Writer)
	}

	newWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      conf.Config.KafkaBrokers,
		Topic:        TopicOrder,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 5 * time.Millisecond,
	})
	orderWriter.Store(TopicOrder, newWriter)
	return newWriter
}
