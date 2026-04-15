package alert

/*
	sends alerts to user
	to discord
	to email
	to phone number via sms
*/
import (
	"context"
	"fmt"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSAlertServiceImpl struct {
	client *twilio.RestClient
	from   string
	to     string
}

const (
	ACCOUNT_SID = "XXXXX"         // TODO: go to twilio and create an account
	AUTH_TOKEN  = "XXXXX"         // TODO: go to twilio and create an auth token
	FROM_NUMBER = "+351960298566" // TODO: go to twilio and create a phone number
)

func NewSMSAlertServiceImpl(toNumber string) *SMSAlertServiceImpl {
	return &SMSAlertServiceImpl{
		client: twilio.NewRestClientWithParams(
			twilio.ClientParams{
				Username: ACCOUNT_SID,
				Password: AUTH_TOKEN,
			}),
		from: FROM_NUMBER,
		to:   toNumber,
	}
}

func (smsService *SMSAlertServiceImpl) SendAlert(ctx context.Context, message string) error {
	// configure the request message
	params := &openapi.CreateMessageParams{}
	params.SetTo(smsService.to)
	params.SetFrom(smsService.from)
	params.SetBody(message)

	// Send the message

	resp, err := smsService.client.Api.CreateMessage(params)
	if err != nil {
		fmt.Printf("Failed to send the message to this number:%v, error:%v\n", smsService.to, err)
		return err
	} else {
		fmt.Printf("Message sent to %s\n", *resp.Sid)
		return nil
	}

}
