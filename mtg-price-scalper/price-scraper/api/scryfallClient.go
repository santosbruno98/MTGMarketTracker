// client for the Scryfall API
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)
//TODO: create a way to put the fetched cards into the bd

const baseURL = "https://api.scryfall.com/"

type Client struct {
	http      *http.Client
	userAgent string
	delay     time.Duration // rate limiting
}

func NewClient() *Client {
	return &Client{
		http:      &http.Client{Timeout: 10 * time.Second},
		userAgent: "mtg-price-scalper/1.0",
		delay:     100 * time.Millisecond,
	}
}

// GetCardPriceByName
func (c *Client) GetCardPriceByName(ctx context.Context, cardName string) (*ScryfallCard, error) {
	// query the scryfall api
	url := baseURL + "cards/" + "named?fuzzy=" + url.QueryEscape(cardName)
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet, url, nil) // returns the last added reprint of the named card
	if err != nil {
		return nil, err
	}

	scryfallResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer scryfallResp.Body.Close()

	if scryfallResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", scryfallResp.StatusCode)
	}

	var card ScryfallCard
	if err := json.NewDecoder(scryfallResp.Body).Decode(&card); err != nil {
		return nil, err
	}

	return &card, nil
}

// GetCardPricesBySetAndNumber
func (c *Client) GetCardPricesBySetAndNumber(ctx context.Context, setCode string, collectorNumber string) (*ScryfallCard, error) {
	url := baseURL + "cards/" + url.QueryEscape(setCode) + "/" + url.QueryEscape(collectorNumber)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil) // returns the last added reprint of the named card
	if err != nil {
		return nil, err
	}

	scryfallResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer scryfallResp.Body.Close()

	if scryfallResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", scryfallResp.StatusCode)
	}

	var card ScryfallCard
	if err := json.NewDecoder(scryfallResp.Body).Decode(&card); err != nil {
		return nil, err
	}

	return &card, nil

}

// GetAllCardReprints
func (c *Client) GetAllCardReprints(ctx context.Context, cardName string) (*ScryfallSearchResponse, error) {
	// query the scryfall api
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, 
		baseURL + "cards/search?q=" + url.QueryEscape(cardName) + "&unique=prints", nil)
	if err != nil {
		return nil, err
	}

	scryfallResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer scryfallResp.Body.Close()

	if scryfallResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", scryfallResp.StatusCode)
	}
	// convert the response to a ScryfallCard struct
	var response ScryfallSearchResponse
	if err := json.NewDecoder(scryfallResp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetAllSetCodes(ctx context.Context) (*ScryfallSet, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL + "sets", nil)
	if err != nil {
		return nil, err
	}

	scryfallResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer scryfallResp.Body.Close()

	if scryfallResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", scryfallResp.StatusCode)
	}

	// convert the response to a ScryfallSetCode struct
	var scryfallSet ScryfallSet
	err = json.NewDecoder(scryfallResp.Body).Decode(&scryfallSet)
	if err != nil {
		return nil, err
	}

	return &scryfallSet, nil

}

func (c *Client) GetCardsBySet(ctx context.Context, setCode string) (*int ,[]ScryfallCard, error) {
	nextURL := baseURL + "cards/" + "search?order=set&q=e:" + url.QueryEscape(setCode) + "&unique=prints"
	var allCards []ScryfallCard
	var totalCards *int

	for nextURL != "" {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, nextURL, nil)
		if err != nil {
			return nil, nil, err
		}

		scryfallResp, err := c.http.Do(req)
		if err != nil {
			return nil, nil, err
		}

		
		if scryfallResp.StatusCode != http.StatusOK {
			return nil , nil, fmt.Errorf("unexpected status code: %d", scryfallResp.StatusCode)
		}

		// convert the response to a ScryfallCard struct
		var page ScryfallSearchResponse
		if err := json.NewDecoder(scryfallResp.Body).Decode(&page); err != nil {
			return nil, nil , err
		}

		scryfallResp.Body.Close()

		totalCards = &page.TotalCards

		allCards = append(allCards, page.Data...)
		
		if page.HasMore {
			nextURL = *page.NextPage
		} else {
			nextURL = ""
		}
	}

	return totalCards, allCards, nil
}
