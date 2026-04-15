package api

import "time"

// CardPrices DB facing struct for price history
type CardPrices struct {
	ScryfallID string
	EUR        *string
	EURFoil    *string
	ScrapedAt  time.Time
}

type ScryfallPrices struct {
	EUR     *string `json:"eur"`
	EURFoil *string `json:"eur_foil"`
}

type ScryfallCard struct {
	ScryfallID      string         `json:"id"`
	CardMarketId    int            `json:"cardmarket_id"`
	Name            string         `json:"name"`
	SetCode         string         `json:"set"`      // can be a list or a single string
	SetName         string         `json:"set_name"` // can be a list or a single string
	Rarity          string         `json:"rarity"`
	CollectorNumber string         `json:"collector_number"` // can be a list or a single string
	Prices          ScryfallPrices `json:"prices"`
	Foil            bool           `json:"foil"`
}

type ScryfallSearchResponse struct {
	Object     string         `json:"object"`
	TotalCards int            `json:"total_cards"`
	HasMore    bool           `json:"has_more"`
	NextPage   *string        `json:"next_page"`
	Data       []ScryfallCard `json:"data"`
}

type ScryfallSet struct {
	Object     string `json:"object"`
	SetId      string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	ReleasedAt string `json:"released_at"`
	CardCount  int    `json:"card_count"`
	SetType    string `json:"set_type"`
}
