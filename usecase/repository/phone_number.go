package usecase

import (
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type PhoneNumberRepository interface {
	PurchasePhoneNumber(phoneNumber string) (*twilioApi.ApiV2010IncomingPhoneNumber, error)
	ListAvailableLocalPhoneNumbers(areaCode string) (*[]twilioApi.ApiV2010AvailablePhoneNumberLocal, error)
	ListAvailableTollFreePhoneNumbers() (*[]twilioApi.ApiV2010AvailablePhoneNumberTollFree, error)
}
