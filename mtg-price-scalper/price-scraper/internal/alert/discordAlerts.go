package alert

/*
	sends alerts to user
	to discord
	to email
	to phone number via sms
*/
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordAlertServiceImpl struct {
	client     *http.Client
	webHookURL string
}

type discordPayload struct {
	Content string `json:"content"`
}

func NewDiscordAlertServiceImpl(client *http.Client, url string) *DiscordAlertServiceImpl {
	return &DiscordAlertServiceImpl{
		client:     client,
		webHookURL: url,
	}
}

func (d *DiscordAlertServiceImpl) SendAlert(ctx context.Context, message string) error {
	payload := discordPayload{
		Content: message,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Failed to serialize the body message:%v\n", err)
		return nil
	}

	resp, err := d.client.Post(d.webHookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Printf("Failed to send the message to discord:%v\n", err)
		return nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}
