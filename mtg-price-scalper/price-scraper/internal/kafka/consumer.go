// the consumer topic will receive the price updates from the scraper and publish into a database for analysis
package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/santosbruno98/mtg-price-scraper/price-scraper/api"
	"github.com/santosbruno98/mtg-price-scraper/price-scraper/internal/scraper/db"
	"github.com/segmentio/kafka-go"
)

const (
	KafkaTopic         = "card.price.scraped"
	KafkaBrokerServer1 = "kafka1:9092"
	KafkaBrokerServer2 = "kafka2:9093"
)

type Consumer struct {
	topicReader *kafka.Reader
	priceRepo   db.PriceRepository
}

// NewConsumer
func NewConsumer(brokers []string, topic string, groupID string, priceRepo db.PriceRepository) *Consumer {
	return &Consumer{
		topicReader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			GroupID: groupID,
			Topic:   topic,
		}),
		priceRepo: priceRepo,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	// start consuming the topic
	// for each message received, call the function to insert the data into the database
	for {
		msg, err := c.topicReader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		go c.processMessage(ctx, msg)

	}
}

func (c *Consumer) processMessage(ctx context.Context, msg kafka.Message) {
	/*
			call the function to insert the data into the database
			topic key : CardNameCardSetCodeCollectorNumber
			topic value : PriceEvent
		*/
		var card api.ScryfallCard
		if err := json.Unmarshal(msg.Value, &card); err != nil {
			fmt.Printf("Failed to unmarshall the message: %v\n", err)
			return
		}
		// calls the function to insert the data into the database -> should be a coroutine
		var eurPrice float32
		if card.Prices.EUR != nil {
			val, err := strconv.ParseFloat(*card.Prices.EUR, 32)
			if err != nil {
				fmt.Printf("Failed to parse the EUR price:  %v\n", err)
				return
			}
			eurPrice = float32(val)
		}
		
		var eurPriceFoil float32
		if card.Prices.EURFoil != nil {
			val, err := strconv.ParseFloat(*card.Prices.EURFoil, 32)
			if err != nil {
				fmt.Printf("Failed to parse the EUR foil price:  %v\n", err)
				return 
			}
			eurPriceFoil = float32(val)
		}

		if err := c.priceRepo.InsertCardPrice(
				ctx,
				card.Name,
				card.SetCode, // cardSetCode is a byte from the topic key and needs to be a string
				card.CollectorNumber,// cardCollectorNumber is a byte from the topic key and needs to be a string
				eurPrice,
				eurPriceFoil,
		); err != nil {
			fmt.Printf("Failed to insert card price in the db: %v\n", err)
			return
		}
			
}

// using the consumer to send data to the database

// close the consumer when the program is closed
func (c *Consumer) Close() error {
	return c.topicReader.Close()
}
// We need an interface to connect to the postgres database and insert the data into the database

// we need to create a struct that will contain the data from the scraper and the database
