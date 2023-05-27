package order_reader

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/markaz/logger"
	"github.com/segmentio/kafka-go"
)

const MessageTopic = "orders_message"

type OrderMessageLogReader struct {
	reader   *kafka.Reader
	observer LogObserver
}

func NewOrderMessageLogReader(brokers []string) LogReader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     MessageTopic,
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})
	return &OrderMessageLogReader{reader: reader}
}

func (r *OrderMessageLogReader) RegisterObserver(observer LogObserver) {
	r.observer = observer
}

func (r *OrderMessageLogReader) Run() {
	logger.Logger.Info("reading log from orders_message")

	for {
		kMessage, err := r.reader.FetchMessage(context.Background())
		if err != nil {
			continue
		}

		var base Base
		err = json.Unmarshal(kMessage.Value, &base)
		if err != nil {
			panic(err)
		}
		switch base.Type {

		case LogTypeStockUpdate:
			var log ProductStockUpdateLog
			err := json.Unmarshal(kMessage.Value, &log)
			if err != nil {
				panic(err)
			}
			err = r.observer.OnStockUpdateLog(log.Products)
			if err != nil {
				continue
			}
			r.CommitOffset(kMessage)
		}
	}
}

/*
This method will commit the offset for a given kafka message
1) @params Kafka.Message
*/
func (r *OrderMessageLogReader) CommitOffset(msg kafka.Message) {
	err := r.commit(1, msg)
	if err != nil {
		fmt.Printf("failed to commit offset %v", err)
	}
}

////////////////////////////////////////////////////////////////////////////////////////
///////////////////  HELPER METHODS  ///////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////

/*
This method will try to commit a kafka message 4 times.
1) @params tries, Kafka.Message
@returns err if any
*/
func (r *OrderMessageLogReader) commit(tries int, msg kafka.Message) error {
	if tries == 4 {
		return fmt.Errorf("failed to commit offset after 3 tries on user done log")
	}
	err := r.reader.CommitMessages(context.TODO(), msg)
	if err != nil {
		return r.commit(tries+1, msg)
	}
	return nil
}
