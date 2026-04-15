package kafka

import (
    "context"
    "encoding/json"

    "github.com/santosbruno98/mtg-price-scraper/price-scraper/api"
    "github.com/segmentio/kafka-go"
)
type Producer struct {
	writer 	*kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: brokers,
			Topic: topic,

		}),
	}
}


func (p *Producer) Publish(ctx context.Context, card api.ScryfallCard) error {
	// Publish the api call of an card as an event to the topic
	data, err := json.Marshal(card)
	if err != nil {
		return err
	}
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(card.Name+card.SetCode+card.CollectorNumber),  // partition by card singleton identifier[SolRingBLC002] 
		Value: data,
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}