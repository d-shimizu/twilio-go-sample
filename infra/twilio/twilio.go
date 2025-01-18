package twilio

import (
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Client struct {
	client *twilio.RestClient
}

func NewClient(accountSid, authToken string) *Client {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &Client{
		client: client,
	}
}

func (c *Client) PurchasePhoneNumber(phoneNumber string) (*twilioApi.ApiV2010IncomingPhoneNumber, error) {
	params := &twilioApi.CreateIncomingPhoneNumberParams{}
	params.SetPhoneNumber(phoneNumber)

	resp, err := c.client.Api.CreateIncomingPhoneNumber(params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListAvailableLocalPhoneNumbers(areaCode string) (*[]twilioApi.ApiV2010AvailablePhoneNumberLocal, error) {
	params := &twilioApi.ListAvailablePhoneNumberLocalParams{}
	//params.SetInRegion("AR")
	//params.SetLimit(20)
	resp, err := c.client.Api.ListAvailablePhoneNumberLocal("US", params)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) ListAvailableTollFreePhoneNumbers() (*[]twilioApi.ApiV2010AvailablePhoneNumberTollFree, error) {
	params := &twilioApi.ListAvailablePhoneNumberTollFreeParams{}
	params.SetLimit(20)
	resp, err := c.client.Api.ListAvailablePhoneNumberTollFree("US", params)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
