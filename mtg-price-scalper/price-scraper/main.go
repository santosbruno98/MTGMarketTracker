package main

import (
	"context"
	"fmt"
    "github.com/santosbruno98/mtg-price-scalper/price-scraper/api/models"
)

type PriceScrapper interface {
    GetCardPrices(ctx context.Context, scryfallID string) (*models.CardPrices, error)
    GetCardPricesBatch(ctx context.Context, scryfallIDs []string) ([]models.CardPrices, error)
}

func main(){
    fmt.Println("Main entry file")
}