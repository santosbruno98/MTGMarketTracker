package api

import (
	"context"
)

type PriceScrapper interface {
	GetCardPriceByName(ctx context.Context, cardName string) (*ScryfallCard, error) // gets a single card by name
	GetCardPricesBySetAndNumber(ctx context.Context, setCode string , collectorNumber string) (*ScryfallCard, error) // gets a single card by set and number
	GetAllCardReprints(ctx context.Context, card_name string) (*ScryfallSearchResponse, error) // gets all cards in a bulk json
	GetAllSetCodes(ctx context.Context) (*ScryfallSet, error) // gets all cards in a bulk json
	GetCardsBySet(ctx context.Context, setCode string) (*int, []ScryfallCard, error) // gets all cards in a bulk json
}
