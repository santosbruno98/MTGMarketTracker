package db

import (
	"context"

	"github.com/santosbruno98/mtg-price-scraper/price-scraper/api"
)

type CardRepository interface {
    GetAllCards(ctx context.Context) ([]api.ScryfallCard, error)
	InsertCard(ctx context.Context, cardName string, cardId string, rarity string) error // cardID = setCode + collectorNumber
}

type PriceRepository interface {
    InsertCardPrice(ctx context.Context, card_name string ,setCode string, collectorNumber string, eurPrice float32, eurFoilPrice float32) error
    GetLatestCardPrice(ctx context.Context, setCode string, collectorNumber string) (*api.CardPrices, error)
    GetPriceHistory(ctx context.Context, setCode string, collectorNumber string, page byte) ([]api.CardPrices, error)
}
