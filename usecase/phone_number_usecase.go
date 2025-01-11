package usecase

import (
	"fmt"
	"localhost/twilio-go-sample/domain/model"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type PhoneNumberService interface {
	PurchasePhoneNumber(phoneNumber string) (*twilioApi.ApiV2010IncomingPhoneNumber, error)
}

type PhoneNumberUseCase struct {
	twilioClient PhoneNumberService
}

func NewPhoneNumberUseCase(client PhoneNumberService) *PhoneNumberUseCase {
	return &PhoneNumberUseCase{
		twilioClient: client,
	}
}

func (uc *PhoneNumberUseCase) PurchasePhoneNumber(phoneNumber string) (*model.PhoneNumber, error) {
	resp, err := uc.twilioClient.PurchasePhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}

	if resp.PhoneNumber == nil || resp.AccountSid == nil {
		return nil, fmt.Errorf("invalid response from Twilio: missing required fields")
	}

	return &model.PhoneNumber{
		AccountSid:  *resp.AccountSid,
		PhoneNumber: *resp.PhoneNumber,
	}, nil
}
