package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq" // Postgres driver - register itself with database/sql
	"github.com/santosbruno98/mtg-price-scraper/price-scraper/api"
)

type priceRepository struct {
	db *sql.DB
}

type cardRepository struct {
	db *sql.DB
}

type CardService struct {
	cardRepo cardRepository
}

func NewCardService(cardRepo cardRepository) (*CardService, error) {
	cardService := &CardService{
		cardRepo: cardRepo}
	return cardService, nil
}

func NewPriceRepository(connStr string) (*priceRepository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &priceRepository{db: db}, nil
}

func (p *priceRepository) Close() error {
	return p.db.Close()
}

func (p *priceRepository) InsertCardPrice(ctx context.Context, cardName string, setCode string, collectorNumber string, eurPrice float32, eurFoilPrice float32) error {
	// executes within an existing transaction
	_, err := p.db.ExecContext(ctx,
		"INSERT INTO card_prices (card_name, set_code, collector_number, eur_price, eur_foil_price) VALUES ($1, $2, $3, $4, $5)",
		cardName, setCode, collectorNumber, eurPrice, eurFoilPrice)
	return err
}

func (p *priceRepository) GetLatestCardPrice(ctx context.Context, setCode string, collectorNumber string) (*api.CardPrices, error) {
	result_row := p.db.QueryRowContext(ctx,
		`SELECT scryfall_id, eur, eur_foil, scraped_at
		FROM card_prices 
		WHERE set_code = $1 AND collector_number = $2 
		ORDER BY scraped_at DESC
		LIMIT 1`,
		setCode, collectorNumber,
	)
	var cardPrice api.CardPrices
	err := result_row.Scan(
		&cardPrice.ScryfallID,
		&cardPrice.EUR,
		&cardPrice.EURFoil,
		&cardPrice.ScrapedAt,
	)
	if err != nil {
		return nil, err
	}

	return &cardPrice, nil
}

func (p *priceRepository) GetPriceHistory(ctx context.Context, setCode string, collectorNumber string, page byte) ([]api.CardPrices, error) {
	rows, err := p.db.QueryContext(ctx,
		`SELECT scryfall_id, eur, eur_foil, scraped_at
		FROM card_prices 
		WHERE set_code = $1 AND collector_number = $2
		ORDER BY scraped_at DESC 
		LIMIT 10 OFFSET $3`,
		setCode,
		collectorNumber,
		(page-1)*10,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cardPrices []api.CardPrices
	for rows.Next() {
		var cardPrice api.CardPrices
		err = rows.Scan(
			&cardPrice.ScryfallID,
			&cardPrice.EUR,
			&cardPrice.EURFoil,
			&cardPrice.ScrapedAt,
		)
		if err != nil {
			return nil, err
		}
		cardPrices = append(cardPrices, cardPrice)
	}

	return cardPrices, nil
}

func (p *cardRepository) GetAllCards(ctx context.Context) ([]api.ScryfallCard, error) {

	rows, err := p.db.QueryContext(ctx,
		`SELECT 
		scryfall_id, 
		name, 
		set_code, 
		collector_number, 
		foil, 
		bought_price, 
		current_price_eur,
		current_price_eur_foil,
		created_at
		FROM cards
		ORDER BY name ASC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cardObjList []api.ScryfallCard
	var cardPriceObj api.CardPrices
	for rows.Next() {
		var cardObj api.ScryfallCard
		err = rows.Scan(
			&cardObj.ScryfallID,
			&cardObj.Name,
			&cardObj.SetCode,
			&cardObj.CollectorNumber,
			&cardObj.Foil,
			&cardObj.Prices.EUR,
			&cardObj.Prices.EURFoil,
			&cardPriceObj.ScrapedAt, // timestmap from the bd ??
		)
		if err != nil {
			return nil, err
		}
		cardObjList = append(cardObjList, cardObj)

	}

	return cardObjList, nil
}

func (c *cardRepository) InsertCard(ctx context.Context, setCode string, collectorNumber string, eurPrice float32, eurFoilPrice float32) error {
	// executes within an existing transaction

	_, err := c.db.ExecContext(ctx,
		// todo: create an incremental ID in both inserts
		"INSERT INTO card (set_code, collector_number, eur_price, eur_foil_price,added_at) VALUES ($1, $2, $3, $4)",
		setCode, collectorNumber, eurPrice, eurFoilPrice, time.Now())
	return err
}
